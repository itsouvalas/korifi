// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"code.cloudfoundry.org/cf-k8s-controllers/api/authorization"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
)

type SpaceRepository struct {
	CreateSpaceStub        func(context.Context, authorization.Info, repositories.CreateSpaceMessage) (repositories.SpaceRecord, error)
	createSpaceMutex       sync.RWMutex
	createSpaceArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateSpaceMessage
	}
	createSpaceReturns struct {
		result1 repositories.SpaceRecord
		result2 error
	}
	createSpaceReturnsOnCall map[int]struct {
		result1 repositories.SpaceRecord
		result2 error
	}
	ListSpacesStub        func(context.Context, authorization.Info, []string, []string) ([]repositories.SpaceRecord, error)
	listSpacesMutex       sync.RWMutex
	listSpacesArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 []string
		arg4 []string
	}
	listSpacesReturns struct {
		result1 []repositories.SpaceRecord
		result2 error
	}
	listSpacesReturnsOnCall map[int]struct {
		result1 []repositories.SpaceRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SpaceRepository) CreateSpace(arg1 context.Context, arg2 authorization.Info, arg3 repositories.CreateSpaceMessage) (repositories.SpaceRecord, error) {
	fake.createSpaceMutex.Lock()
	ret, specificReturn := fake.createSpaceReturnsOnCall[len(fake.createSpaceArgsForCall)]
	fake.createSpaceArgsForCall = append(fake.createSpaceArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateSpaceMessage
	}{arg1, arg2, arg3})
	stub := fake.CreateSpaceStub
	fakeReturns := fake.createSpaceReturns
	fake.recordInvocation("CreateSpace", []interface{}{arg1, arg2, arg3})
	fake.createSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SpaceRepository) CreateSpaceCallCount() int {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return len(fake.createSpaceArgsForCall)
}

func (fake *SpaceRepository) CreateSpaceCalls(stub func(context.Context, authorization.Info, repositories.CreateSpaceMessage) (repositories.SpaceRecord, error)) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = stub
}

func (fake *SpaceRepository) CreateSpaceArgsForCall(i int) (context.Context, authorization.Info, repositories.CreateSpaceMessage) {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	argsForCall := fake.createSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SpaceRepository) CreateSpaceReturns(result1 repositories.SpaceRecord, result2 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	fake.createSpaceReturns = struct {
		result1 repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *SpaceRepository) CreateSpaceReturnsOnCall(i int, result1 repositories.SpaceRecord, result2 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	if fake.createSpaceReturnsOnCall == nil {
		fake.createSpaceReturnsOnCall = make(map[int]struct {
			result1 repositories.SpaceRecord
			result2 error
		})
	}
	fake.createSpaceReturnsOnCall[i] = struct {
		result1 repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *SpaceRepository) ListSpaces(arg1 context.Context, arg2 authorization.Info, arg3 []string, arg4 []string) ([]repositories.SpaceRecord, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	var arg4Copy []string
	if arg4 != nil {
		arg4Copy = make([]string, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.listSpacesMutex.Lock()
	ret, specificReturn := fake.listSpacesReturnsOnCall[len(fake.listSpacesArgsForCall)]
	fake.listSpacesArgsForCall = append(fake.listSpacesArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 []string
		arg4 []string
	}{arg1, arg2, arg3Copy, arg4Copy})
	stub := fake.ListSpacesStub
	fakeReturns := fake.listSpacesReturns
	fake.recordInvocation("ListSpaces", []interface{}{arg1, arg2, arg3Copy, arg4Copy})
	fake.listSpacesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SpaceRepository) ListSpacesCallCount() int {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	return len(fake.listSpacesArgsForCall)
}

func (fake *SpaceRepository) ListSpacesCalls(stub func(context.Context, authorization.Info, []string, []string) ([]repositories.SpaceRecord, error)) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = stub
}

func (fake *SpaceRepository) ListSpacesArgsForCall(i int) (context.Context, authorization.Info, []string, []string) {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	argsForCall := fake.listSpacesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *SpaceRepository) ListSpacesReturns(result1 []repositories.SpaceRecord, result2 error) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = nil
	fake.listSpacesReturns = struct {
		result1 []repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *SpaceRepository) ListSpacesReturnsOnCall(i int, result1 []repositories.SpaceRecord, result2 error) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = nil
	if fake.listSpacesReturnsOnCall == nil {
		fake.listSpacesReturnsOnCall = make(map[int]struct {
			result1 []repositories.SpaceRecord
			result2 error
		})
	}
	fake.listSpacesReturnsOnCall[i] = struct {
		result1 []repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *SpaceRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SpaceRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ apis.SpaceRepository = new(SpaceRepository)