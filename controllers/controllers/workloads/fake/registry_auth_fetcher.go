// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/korifi/controllers/controllers/workloads"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type RegistryAuthFetcher struct {
	Stub        func(context.Context, string) (remote.Option, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	returns struct {
		result1 remote.Option
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 remote.Option
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RegistryAuthFetcher) Spy(arg1 context.Context, arg2 string) (remote.Option, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("RegistryAuthFetcher", []interface{}{arg1, arg2})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return returns.result1, returns.result2
}

func (fake *RegistryAuthFetcher) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *RegistryAuthFetcher) Calls(stub func(context.Context, string) (remote.Option, error)) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *RegistryAuthFetcher) ArgsForCall(i int) (context.Context, string) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2
}

func (fake *RegistryAuthFetcher) Returns(result1 remote.Option, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 remote.Option
		result2 error
	}{result1, result2}
}

func (fake *RegistryAuthFetcher) ReturnsOnCall(i int, result1 remote.Option, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 remote.Option
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 remote.Option
		result2 error
	}{result1, result2}
}

func (fake *RegistryAuthFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RegistryAuthFetcher) recordInvocation(key string, args []interface{}) {
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

var _ workloads.RegistryAuthFetcher = new(RegistryAuthFetcher).Spy
