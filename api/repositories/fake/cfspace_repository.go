// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
)

type CFSpaceRepository struct {
	CreateSpaceStub        func(context.Context, repositories.SpaceCreateMessage) (repositories.SpaceRecord, error)
	createSpaceMutex       sync.RWMutex
	createSpaceArgsForCall []struct {
		arg1 context.Context
		arg2 repositories.SpaceCreateMessage
	}
	createSpaceReturns struct {
		result1 repositories.SpaceRecord
		result2 error
	}
	createSpaceReturnsOnCall map[int]struct {
		result1 repositories.SpaceRecord
		result2 error
	}
	FetchSpacesStub        func(context.Context, []string, []string) ([]repositories.SpaceRecord, error)
	fetchSpacesMutex       sync.RWMutex
	fetchSpacesArgsForCall []struct {
		arg1 context.Context
		arg2 []string
		arg3 []string
	}
	fetchSpacesReturns struct {
		result1 []repositories.SpaceRecord
		result2 error
	}
	fetchSpacesReturnsOnCall map[int]struct {
		result1 []repositories.SpaceRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFSpaceRepository) CreateSpace(arg1 context.Context, arg2 repositories.SpaceCreateMessage) (repositories.SpaceRecord, error) {
	fake.createSpaceMutex.Lock()
	ret, specificReturn := fake.createSpaceReturnsOnCall[len(fake.createSpaceArgsForCall)]
	fake.createSpaceArgsForCall = append(fake.createSpaceArgsForCall, struct {
		arg1 context.Context
		arg2 repositories.SpaceCreateMessage
	}{arg1, arg2})
	stub := fake.CreateSpaceStub
	fakeReturns := fake.createSpaceReturns
	fake.recordInvocation("CreateSpace", []interface{}{arg1, arg2})
	fake.createSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFSpaceRepository) CreateSpaceCallCount() int {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return len(fake.createSpaceArgsForCall)
}

func (fake *CFSpaceRepository) CreateSpaceCalls(stub func(context.Context, repositories.SpaceCreateMessage) (repositories.SpaceRecord, error)) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = stub
}

func (fake *CFSpaceRepository) CreateSpaceArgsForCall(i int) (context.Context, repositories.SpaceCreateMessage) {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	argsForCall := fake.createSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CFSpaceRepository) CreateSpaceReturns(result1 repositories.SpaceRecord, result2 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	fake.createSpaceReturns = struct {
		result1 repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFSpaceRepository) CreateSpaceReturnsOnCall(i int, result1 repositories.SpaceRecord, result2 error) {
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

func (fake *CFSpaceRepository) FetchSpaces(arg1 context.Context, arg2 []string, arg3 []string) ([]repositories.SpaceRecord, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.fetchSpacesMutex.Lock()
	ret, specificReturn := fake.fetchSpacesReturnsOnCall[len(fake.fetchSpacesArgsForCall)]
	fake.fetchSpacesArgsForCall = append(fake.fetchSpacesArgsForCall, struct {
		arg1 context.Context
		arg2 []string
		arg3 []string
	}{arg1, arg2Copy, arg3Copy})
	stub := fake.FetchSpacesStub
	fakeReturns := fake.fetchSpacesReturns
	fake.recordInvocation("FetchSpaces", []interface{}{arg1, arg2Copy, arg3Copy})
	fake.fetchSpacesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFSpaceRepository) FetchSpacesCallCount() int {
	fake.fetchSpacesMutex.RLock()
	defer fake.fetchSpacesMutex.RUnlock()
	return len(fake.fetchSpacesArgsForCall)
}

func (fake *CFSpaceRepository) FetchSpacesCalls(stub func(context.Context, []string, []string) ([]repositories.SpaceRecord, error)) {
	fake.fetchSpacesMutex.Lock()
	defer fake.fetchSpacesMutex.Unlock()
	fake.FetchSpacesStub = stub
}

func (fake *CFSpaceRepository) FetchSpacesArgsForCall(i int) (context.Context, []string, []string) {
	fake.fetchSpacesMutex.RLock()
	defer fake.fetchSpacesMutex.RUnlock()
	argsForCall := fake.fetchSpacesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFSpaceRepository) FetchSpacesReturns(result1 []repositories.SpaceRecord, result2 error) {
	fake.fetchSpacesMutex.Lock()
	defer fake.fetchSpacesMutex.Unlock()
	fake.FetchSpacesStub = nil
	fake.fetchSpacesReturns = struct {
		result1 []repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFSpaceRepository) FetchSpacesReturnsOnCall(i int, result1 []repositories.SpaceRecord, result2 error) {
	fake.fetchSpacesMutex.Lock()
	defer fake.fetchSpacesMutex.Unlock()
	fake.FetchSpacesStub = nil
	if fake.fetchSpacesReturnsOnCall == nil {
		fake.fetchSpacesReturnsOnCall = make(map[int]struct {
			result1 []repositories.SpaceRecord
			result2 error
		})
	}
	fake.fetchSpacesReturnsOnCall[i] = struct {
		result1 []repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFSpaceRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	fake.fetchSpacesMutex.RLock()
	defer fake.fetchSpacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFSpaceRepository) recordInvocation(key string, args []interface{}) {
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

var _ repositories.CFSpaceRepository = new(CFSpaceRepository)
