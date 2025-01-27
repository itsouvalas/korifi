// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/korifi/api/actions"
	"code.cloudfoundry.org/korifi/api/authorization"
	"code.cloudfoundry.org/korifi/api/repositories"
)

type PodRepository struct {
	GetRuntimeLogsForAppStub        func(context.Context, authorization.Info, repositories.RuntimeLogsMessage) ([]repositories.LogRecord, error)
	getRuntimeLogsForAppMutex       sync.RWMutex
	getRuntimeLogsForAppArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.RuntimeLogsMessage
	}
	getRuntimeLogsForAppReturns struct {
		result1 []repositories.LogRecord
		result2 error
	}
	getRuntimeLogsForAppReturnsOnCall map[int]struct {
		result1 []repositories.LogRecord
		result2 error
	}
	ListPodStatsStub        func(context.Context, authorization.Info, repositories.ListPodStatsMessage) ([]repositories.PodStatsRecord, error)
	listPodStatsMutex       sync.RWMutex
	listPodStatsArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.ListPodStatsMessage
	}
	listPodStatsReturns struct {
		result1 []repositories.PodStatsRecord
		result2 error
	}
	listPodStatsReturnsOnCall map[int]struct {
		result1 []repositories.PodStatsRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PodRepository) GetRuntimeLogsForApp(arg1 context.Context, arg2 authorization.Info, arg3 repositories.RuntimeLogsMessage) ([]repositories.LogRecord, error) {
	fake.getRuntimeLogsForAppMutex.Lock()
	ret, specificReturn := fake.getRuntimeLogsForAppReturnsOnCall[len(fake.getRuntimeLogsForAppArgsForCall)]
	fake.getRuntimeLogsForAppArgsForCall = append(fake.getRuntimeLogsForAppArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.RuntimeLogsMessage
	}{arg1, arg2, arg3})
	stub := fake.GetRuntimeLogsForAppStub
	fakeReturns := fake.getRuntimeLogsForAppReturns
	fake.recordInvocation("GetRuntimeLogsForApp", []interface{}{arg1, arg2, arg3})
	fake.getRuntimeLogsForAppMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PodRepository) GetRuntimeLogsForAppCallCount() int {
	fake.getRuntimeLogsForAppMutex.RLock()
	defer fake.getRuntimeLogsForAppMutex.RUnlock()
	return len(fake.getRuntimeLogsForAppArgsForCall)
}

func (fake *PodRepository) GetRuntimeLogsForAppCalls(stub func(context.Context, authorization.Info, repositories.RuntimeLogsMessage) ([]repositories.LogRecord, error)) {
	fake.getRuntimeLogsForAppMutex.Lock()
	defer fake.getRuntimeLogsForAppMutex.Unlock()
	fake.GetRuntimeLogsForAppStub = stub
}

func (fake *PodRepository) GetRuntimeLogsForAppArgsForCall(i int) (context.Context, authorization.Info, repositories.RuntimeLogsMessage) {
	fake.getRuntimeLogsForAppMutex.RLock()
	defer fake.getRuntimeLogsForAppMutex.RUnlock()
	argsForCall := fake.getRuntimeLogsForAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PodRepository) GetRuntimeLogsForAppReturns(result1 []repositories.LogRecord, result2 error) {
	fake.getRuntimeLogsForAppMutex.Lock()
	defer fake.getRuntimeLogsForAppMutex.Unlock()
	fake.GetRuntimeLogsForAppStub = nil
	fake.getRuntimeLogsForAppReturns = struct {
		result1 []repositories.LogRecord
		result2 error
	}{result1, result2}
}

func (fake *PodRepository) GetRuntimeLogsForAppReturnsOnCall(i int, result1 []repositories.LogRecord, result2 error) {
	fake.getRuntimeLogsForAppMutex.Lock()
	defer fake.getRuntimeLogsForAppMutex.Unlock()
	fake.GetRuntimeLogsForAppStub = nil
	if fake.getRuntimeLogsForAppReturnsOnCall == nil {
		fake.getRuntimeLogsForAppReturnsOnCall = make(map[int]struct {
			result1 []repositories.LogRecord
			result2 error
		})
	}
	fake.getRuntimeLogsForAppReturnsOnCall[i] = struct {
		result1 []repositories.LogRecord
		result2 error
	}{result1, result2}
}

func (fake *PodRepository) ListPodStats(arg1 context.Context, arg2 authorization.Info, arg3 repositories.ListPodStatsMessage) ([]repositories.PodStatsRecord, error) {
	fake.listPodStatsMutex.Lock()
	ret, specificReturn := fake.listPodStatsReturnsOnCall[len(fake.listPodStatsArgsForCall)]
	fake.listPodStatsArgsForCall = append(fake.listPodStatsArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.ListPodStatsMessage
	}{arg1, arg2, arg3})
	stub := fake.ListPodStatsStub
	fakeReturns := fake.listPodStatsReturns
	fake.recordInvocation("ListPodStats", []interface{}{arg1, arg2, arg3})
	fake.listPodStatsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PodRepository) ListPodStatsCallCount() int {
	fake.listPodStatsMutex.RLock()
	defer fake.listPodStatsMutex.RUnlock()
	return len(fake.listPodStatsArgsForCall)
}

func (fake *PodRepository) ListPodStatsCalls(stub func(context.Context, authorization.Info, repositories.ListPodStatsMessage) ([]repositories.PodStatsRecord, error)) {
	fake.listPodStatsMutex.Lock()
	defer fake.listPodStatsMutex.Unlock()
	fake.ListPodStatsStub = stub
}

func (fake *PodRepository) ListPodStatsArgsForCall(i int) (context.Context, authorization.Info, repositories.ListPodStatsMessage) {
	fake.listPodStatsMutex.RLock()
	defer fake.listPodStatsMutex.RUnlock()
	argsForCall := fake.listPodStatsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PodRepository) ListPodStatsReturns(result1 []repositories.PodStatsRecord, result2 error) {
	fake.listPodStatsMutex.Lock()
	defer fake.listPodStatsMutex.Unlock()
	fake.ListPodStatsStub = nil
	fake.listPodStatsReturns = struct {
		result1 []repositories.PodStatsRecord
		result2 error
	}{result1, result2}
}

func (fake *PodRepository) ListPodStatsReturnsOnCall(i int, result1 []repositories.PodStatsRecord, result2 error) {
	fake.listPodStatsMutex.Lock()
	defer fake.listPodStatsMutex.Unlock()
	fake.ListPodStatsStub = nil
	if fake.listPodStatsReturnsOnCall == nil {
		fake.listPodStatsReturnsOnCall = make(map[int]struct {
			result1 []repositories.PodStatsRecord
			result2 error
		})
	}
	fake.listPodStatsReturnsOnCall[i] = struct {
		result1 []repositories.PodStatsRecord
		result2 error
	}{result1, result2}
}

func (fake *PodRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getRuntimeLogsForAppMutex.RLock()
	defer fake.getRuntimeLogsForAppMutex.RUnlock()
	fake.listPodStatsMutex.RLock()
	defer fake.listPodStatsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PodRepository) recordInvocation(key string, args []interface{}) {
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

var _ actions.PodRepository = new(PodRepository)
