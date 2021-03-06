// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	context "context"
	sync "sync"

	lager "code.cloudfoundry.org/lager"
	creds "github.com/concourse/concourse/atc/creds"
	db "github.com/concourse/concourse/atc/db"
	resource "github.com/concourse/concourse/atc/resource"
	worker "github.com/concourse/concourse/atc/worker"
)

type FakeResourceFactory struct {
	NewResourceStub        func(context.Context, lager.Logger, db.ContainerOwner, db.ContainerMetadata, worker.ContainerSpec, creds.VersionedResourceTypes, worker.ImageFetchingDelegate) (resource.Resource, error)
	newResourceMutex       sync.RWMutex
	newResourceArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.ContainerOwner
		arg4 db.ContainerMetadata
		arg5 worker.ContainerSpec
		arg6 creds.VersionedResourceTypes
		arg7 worker.ImageFetchingDelegate
	}
	newResourceReturns struct {
		result1 resource.Resource
		result2 error
	}
	newResourceReturnsOnCall map[int]struct {
		result1 resource.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactory) NewResource(arg1 context.Context, arg2 lager.Logger, arg3 db.ContainerOwner, arg4 db.ContainerMetadata, arg5 worker.ContainerSpec, arg6 creds.VersionedResourceTypes, arg7 worker.ImageFetchingDelegate) (resource.Resource, error) {
	fake.newResourceMutex.Lock()
	ret, specificReturn := fake.newResourceReturnsOnCall[len(fake.newResourceArgsForCall)]
	fake.newResourceArgsForCall = append(fake.newResourceArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.ContainerOwner
		arg4 db.ContainerMetadata
		arg5 worker.ContainerSpec
		arg6 creds.VersionedResourceTypes
		arg7 worker.ImageFetchingDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("NewResource", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.newResourceMutex.Unlock()
	if fake.NewResourceStub != nil {
		return fake.NewResourceStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newResourceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceFactory) NewResourceCallCount() int {
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	return len(fake.newResourceArgsForCall)
}

func (fake *FakeResourceFactory) NewResourceArgsForCall(i int) (context.Context, lager.Logger, db.ContainerOwner, db.ContainerMetadata, worker.ContainerSpec, creds.VersionedResourceTypes, worker.ImageFetchingDelegate) {
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	argsForCall := fake.newResourceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeResourceFactory) NewResourceReturns(result1 resource.Resource, result2 error) {
	fake.NewResourceStub = nil
	fake.newResourceReturns = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) NewResourceReturnsOnCall(i int, result1 resource.Resource, result2 error) {
	fake.NewResourceStub = nil
	if fake.newResourceReturnsOnCall == nil {
		fake.newResourceReturnsOnCall = make(map[int]struct {
			result1 resource.Resource
			result2 error
		})
	}
	fake.newResourceReturnsOnCall[i] = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceFactory) recordInvocation(key string, args []interface{}) {
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

var _ resource.ResourceFactory = new(FakeResourceFactory)
