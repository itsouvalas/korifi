package repositories_test

import (
	"context"
	"fmt"
	"time"

	"code.cloudfoundry.org/korifi/api/apierrors"
	"code.cloudfoundry.org/korifi/api/repositories"
	workloads "code.cloudfoundry.org/korifi/controllers/apis/workloads/v1alpha1"
	"code.cloudfoundry.org/korifi/tests/matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	hncv1alpha2 "sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

var _ = Describe("SpaceRepository", func() {
	var (
		ctx       context.Context
		orgRepo   *repositories.OrgRepo
		spaceRepo *repositories.SpaceRepo
	)

	const (
		spaceLevel = "2"
	)

	BeforeEach(func() {
		ctx = context.Background()
		orgRepo = repositories.NewOrgRepo(rootNamespace, k8sClient, userClientFactory, nsPerms, time.Millisecond*2000)
		spaceRepo = repositories.NewSpaceRepo(orgRepo, k8sClient, userClientFactory, nsPerms, time.Millisecond*2000)
	})

	Describe("Create", func() {
		var (
			createErr                   error
			orgGUID                     string
			spaceName                   string
			space                       repositories.SpaceRecord
			doSpaceControllerSimulation bool
			createRoleBindings          bool
		)

		waitForCFSpace := func(anchorNamespace string, spaceName string, done chan bool) (*workloads.CFSpace, error) {
			for {
				select {
				case <-done:
					return nil, fmt.Errorf("waitForCFSpace was 'signalled' to stop polling")
				default:
				}

				var spaceList workloads.CFSpaceList
				err := k8sClient.List(ctx, &spaceList, client.InNamespace(anchorNamespace))
				if err != nil {
					return nil, fmt.Errorf("waitForCFSpace failed")
				}

				var matches []workloads.CFSpace
				for _, space := range spaceList.Items {
					if space.Spec.DisplayName == spaceName {
						matches = append(matches, space)
					}
				}
				if len(matches) > 1 {
					return nil, fmt.Errorf("waitForCFSpace found multiple anchors")
				}

				if len(matches) == 1 {
					return &matches[0], nil
				}

				time.Sleep(time.Millisecond * 100)
			}
		}

		simulateSpaceController := func(anchorNamespace string, spaceName string, createRoleBindings bool, depthLevel string, done chan bool) {
			defer GinkgoRecover()

			space, err := waitForCFSpace(anchorNamespace, spaceName, done)
			if err != nil {
				return
			}

			createNamespace(ctx, anchorNamespace, space.Name, map[string]string{
				rootNamespace + hncv1alpha2.LabelTreeDepthSuffix: depthLevel,
			})

			if createRoleBindings {
				createRoleBinding(ctx, userName, adminRole.Name, space.Name)
			}

			meta.SetStatusCondition(&(space.Status.Conditions), metav1.Condition{
				Type:    "Ready",
				Status:  metav1.ConditionTrue,
				Reason:  "blah",
				Message: "blah",
			})
			Expect(
				k8sClient.Status().Update(ctx, space),
			).To(Succeed())
		}

		BeforeEach(func() {
			doSpaceControllerSimulation = true
			spaceName = prefixedGUID("space-name")
			org := createOrgWithCleanup(ctx, prefixedGUID("org"))
			orgGUID = org.Name
			createRoleBindings = true
		})

		JustBeforeEach(func() {
			if doSpaceControllerSimulation {
				done := make(chan bool, 1)
				defer func(done chan bool) { done <- true }(done)

				go simulateSpaceController(orgGUID, spaceName, createRoleBindings, spaceLevel, done)
			}

			space, createErr = spaceRepo.CreateSpace(ctx, authInfo, repositories.CreateSpaceMessage{
				Name:                     spaceName,
				OrganizationGUID:         orgGUID,
				ImageRegistryCredentials: "imageRegistryCredentials",
			})
		})

		When("the user doesn't have the admin role", func() {
			BeforeEach(func() {
				createRoleBinding(ctx, userName, orgUserRole.Name, orgGUID)
			})

			It("fails when creating a space", func() {
				Expect(createErr).To(matchers.WrapErrorAssignableToTypeOf(apierrors.ForbiddenError{}))
			})
		})

		When("the user has the admin role", func() {
			BeforeEach(func() {
				createRoleBinding(ctx, userName, adminRole.Name, orgGUID)
			})

			It("creates a CFSpace resource in the org namespace", func() {
				Expect(createErr).NotTo(HaveOccurred())

				spaceCR := new(workloads.CFSpace)
				Expect(k8sClient.Get(ctx, client.ObjectKey{Namespace: orgGUID, Name: space.GUID}, spaceCR)).To(Succeed())

				Expect(space.Name).To(Equal(spaceName))
				Expect(space.GUID).To(HavePrefix("cf-space-"))
				Expect(space.CreatedAt).To(BeTemporally("~", time.Now(), 2*time.Second))
				Expect(space.UpdatedAt).To(BeTemporally("~", time.Now(), 2*time.Second))
			})

			When("the org does not exist", func() {
				BeforeEach(func() {
					orgGUID = "does-not-exist"
				})

				It("returns an error", func() {
					Expect(createErr).To(matchers.WrapErrorAssignableToTypeOf(apierrors.NotFoundError{}))
				})
			})

			When("the client fails to create the space", func() {
				BeforeEach(func() {
					spaceName = "this-string-has-illegal-characters-ц"
				})

				It("fails", func() {
					Expect(createErr).To(HaveOccurred())
				})
			})

			When("the space isn't ready in the timeout", func() {
				BeforeEach(func() {
					doSpaceControllerSimulation = false
				})

				It("returns an error", func() {
					Expect(createErr).To(MatchError(ContainSubstring("cf space did not get Condition `Ready`: 'True'")))
				})
			})

			When("CFSpaceController fails to propagate the role bindings in the timeout", func() {
				BeforeEach(func() {
					createRoleBindings = false
				})

				It("returns an error", func() {
					Expect(createErr).To(MatchError(ContainSubstring("failed establishing permissions in new namespace")))
				})
			})
		})
	})

	Describe("List", func() {
		var cfOrg1, cfOrg2, cfOrg3 *workloads.CFOrg
		var space11, space12, space21, space22, space31, space32 *workloads.CFSpace

		BeforeEach(func() {
			ctx = context.Background()

			cfOrg1 = createOrgWithCleanup(ctx, prefixedGUID("org1"))
			createRoleBinding(ctx, userName, orgUserRole.Name, cfOrg1.Name)
			cfOrg2 = createOrgWithCleanup(ctx, prefixedGUID("org2"))
			createRoleBinding(ctx, userName, orgUserRole.Name, cfOrg2.Name)
			cfOrg3 = createOrgWithCleanup(ctx, prefixedGUID("org3"))
			createRoleBinding(ctx, userName, orgUserRole.Name, cfOrg3.Name)

			space11 = createSpaceWithCleanup(ctx, cfOrg1.Name, "space1")
			createRoleBinding(ctx, userName, spaceDeveloperRole.Name, space11.Name)
			space12 = createSpaceWithCleanup(ctx, cfOrg1.Name, "space2")
			createRoleBinding(ctx, userName, spaceDeveloperRole.Name, space12.Name)

			space21 = createSpaceWithCleanup(ctx, cfOrg2.Name, "space1")
			createRoleBinding(ctx, userName, spaceDeveloperRole.Name, space21.Name)
			space22 = createSpaceWithCleanup(ctx, cfOrg2.Name, "space3")
			createRoleBinding(ctx, userName, spaceDeveloperRole.Name, space22.Name)

			space31 = createSpaceWithCleanup(ctx, cfOrg3.Name, "space1")
			createRoleBinding(ctx, userName, spaceDeveloperRole.Name, space31.Name)
			space32 = createSpaceWithCleanup(ctx, cfOrg3.Name, "space4")
			createRoleBinding(ctx, userName, spaceDeveloperRole.Name, space32.Name)
		})

		It("returns the 6 spaces", func() {
			spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{})
			Expect(err).NotTo(HaveOccurred())

			Expect(spaces).To(ConsistOf(
				repositories.SpaceRecord{
					Name:             "space1",
					CreatedAt:        space11.CreationTimestamp.Time,
					UpdatedAt:        space11.CreationTimestamp.Time,
					GUID:             space11.Name,
					OrganizationGUID: cfOrg1.Name,
				},
				repositories.SpaceRecord{
					Name:             "space2",
					CreatedAt:        space12.CreationTimestamp.Time,
					UpdatedAt:        space12.CreationTimestamp.Time,
					GUID:             space12.Name,
					OrganizationGUID: cfOrg1.Name,
				},
				repositories.SpaceRecord{
					Name:             "space1",
					CreatedAt:        space21.CreationTimestamp.Time,
					UpdatedAt:        space21.CreationTimestamp.Time,
					GUID:             space21.Name,
					OrganizationGUID: cfOrg2.Name,
				},
				repositories.SpaceRecord{
					Name:             "space3",
					CreatedAt:        space22.CreationTimestamp.Time,
					UpdatedAt:        space22.CreationTimestamp.Time,
					GUID:             space22.Name,
					OrganizationGUID: cfOrg2.Name,
				},
				repositories.SpaceRecord{
					Name:             "space1",
					CreatedAt:        space31.CreationTimestamp.Time,
					UpdatedAt:        space31.CreationTimestamp.Time,
					GUID:             space31.Name,
					OrganizationGUID: cfOrg3.Name,
				},
				repositories.SpaceRecord{
					Name:             "space4",
					CreatedAt:        space32.CreationTimestamp.Time,
					UpdatedAt:        space32.CreationTimestamp.Time,
					GUID:             space32.Name,
					OrganizationGUID: cfOrg3.Name,
				},
			))
		})

		When("the space anchor is not ready", func() {
			BeforeEach(func() {
				meta.SetStatusCondition(&(space11.Status.Conditions), metav1.Condition{
					Type:    "Ready",
					Status:  metav1.ConditionFalse,
					Reason:  "cus",
					Message: "cus",
				})
				Expect(k8sClient.Status().Update(ctx, space11)).To(Succeed())
			})

			It("does not list it", func() {
				spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{})
				Expect(err).NotTo(HaveOccurred())

				Expect(spaces).NotTo(ContainElement(
					repositories.SpaceRecord{
						Name:             "space1",
						CreatedAt:        space11.CreationTimestamp.Time,
						UpdatedAt:        space11.CreationTimestamp.Time,
						GUID:             space11.Name,
						OrganizationGUID: cfOrg1.Name,
					},
				))
			})
		})

		When("filtering by org guids", func() {
			It("only returns the spaces belonging to the specified org guids", func() {
				spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{
					OrganizationGUIDs: []string{cfOrg1.Name, cfOrg3.Name, "does-not-exist"},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(spaces).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg1.Name),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg3.Name),
					}),
					MatchFields(IgnoreExtras, Fields{"Name": Equal("space2")}),
					MatchFields(IgnoreExtras, Fields{"Name": Equal("space4")}),
				))
			})
		})

		When("filtering by space names", func() {
			It("only returns the spaces matching the specified names", func() {
				spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{
					Names: []string{"space1", "space3", "does-not-exist"},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(spaces).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg1.Name),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg2.Name),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg3.Name),
					}),
					MatchFields(IgnoreExtras, Fields{"Name": Equal("space3")}),
				))
			})
		})

		When("filtering by space guids", func() {
			It("only returns the spaces matching the specified guids", func() {
				spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{
					GUIDs: []string{space11.Name, space21.Name, "does-not-exist"},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(spaces).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg1.Name),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg2.Name),
					}),
				))
			})
		})

		When("filtering by org guids, space names and space guids", func() {
			It("only returns the spaces matching the specified names", func() {
				spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{
					OrganizationGUIDs: []string{cfOrg1.Name, cfOrg2.Name},
					Names:             []string{"space1", "space2", "space4"},
					GUIDs:             []string{space11.Name, space21.Name},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(spaces).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg1.Name),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Name":             Equal("space1"),
						"OrganizationGUID": Equal(cfOrg2.Name),
					}),
				))
			})
		})

		When("filtering by space names that don't exist", func() {
			It("only returns the spaces matching the specified names", func() {
				spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{
					Names: []string{"does-not-exist", "still-does-not-exist"},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(spaces).To(BeEmpty())
			})
		})

		When("filtering by org guids that don't exist", func() {
			It("only returns the spaces matching the specified names", func() {
				spaces, err := spaceRepo.ListSpaces(ctx, authInfo, repositories.ListSpacesMessage{
					OrganizationGUIDs: []string{"does-not-exist", "still-does-not-exist"},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(spaces).To(BeEmpty())
			})
		})
	})

	Describe("Get", func() {
		var (
			cfOrg   *workloads.CFOrg
			cfSpace *workloads.CFSpace
		)

		BeforeEach(func() {
			cfOrg = createOrgWithCleanup(ctx, "the-org")
			createRoleBinding(ctx, userName, orgUserRole.Name, cfOrg.Name)
			cfSpace = createSpaceWithCleanup(ctx, cfOrg.Name, "the-space")
			createRoleBinding(ctx, userName, spaceDeveloperRole.Name, cfSpace.Name)
		})

		It("gets the space resource", func() {
			spaceRecord, err := spaceRepo.GetSpace(ctx, authInfo, cfSpace.Name)
			Expect(err).NotTo(HaveOccurred())
			Expect(spaceRecord.Name).To(Equal("the-space"))
			Expect(spaceRecord.OrganizationGUID).To(Equal(cfOrg.Name))
		})

		When("the space doesn't exist", func() {
			It("errors", func() {
				_, err := spaceRepo.GetSpace(ctx, authInfo, "non-existent-space")
				Expect(err).To(MatchError(ContainSubstring("not found")))
			})
		})
	})

	Describe("Delete", func() {
		var cfOrg *workloads.CFOrg
		var cfSpace *workloads.CFSpace

		BeforeEach(func() {
			cfOrg = createOrgWithCleanup(ctx, prefixedGUID("org"))
			cfSpace = createSpaceWithCleanup(ctx, cfOrg.Name, "the-space")
		})

		When("the user has permission to delete spaces", func() {
			BeforeEach(func() {
				beforeCtx := context.Background()
				createRoleBinding(beforeCtx, userName, adminRole.Name, cfSpace.Namespace)
			})

			It("deletes the space resource", func() {
				err := spaceRepo.DeleteSpace(ctx, authInfo, repositories.DeleteSpaceMessage{
					GUID:             cfSpace.Name,
					OrganizationGUID: cfOrg.Name,
				})
				Expect(err).NotTo(HaveOccurred())

				anchor := &hncv1alpha2.SubnamespaceAnchor{}
				err = k8sClient.Get(ctx, client.ObjectKey{Namespace: cfOrg.Name, Name: cfSpace.Name}, anchor)
				Expect(err).To(MatchError(ContainSubstring("not found")))
			})

			When("the space doesn't exist", func() {
				It("errors", func() {
					err := spaceRepo.DeleteSpace(ctx, authInfo, repositories.DeleteSpaceMessage{
						GUID:             "non-existent-space",
						OrganizationGUID: cfOrg.Name,
					})
					Expect(err).To(MatchError(ContainSubstring("not found")))
				})
			})
		})

		When("the user does not have permission to delete spaces", func() {
			It("errors with forbidden", func() {
				err := spaceRepo.DeleteSpace(ctx, authInfo, repositories.DeleteSpaceMessage{
					GUID:             cfSpace.Name,
					OrganizationGUID: cfOrg.Name,
				})
				Expect(err).To(matchers.WrapErrorAssignableToTypeOf(apierrors.ForbiddenError{}))
			})

			When("the space doesn't exist", func() {
				It("errors with forbidden", func() {
					err := spaceRepo.DeleteSpace(ctx, authInfo, repositories.DeleteSpaceMessage{
						GUID:             "non-existent-space",
						OrganizationGUID: cfOrg.Name,
					})
					Expect(err).To(matchers.WrapErrorAssignableToTypeOf(apierrors.ForbiddenError{}))
				})
			})
		})
	})
})
