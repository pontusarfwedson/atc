// This file was generated by counterfeiter
package resourcefakes

import (
	"os"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeResource struct {
	GetStub        func(worker.Volume, resource.IOConfig, atc.Source, atc.Params, atc.Version, <-chan os.Signal, chan<- struct{}) (resource.VersionedSource, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 worker.Volume
		arg2 resource.IOConfig
		arg3 atc.Source
		arg4 atc.Params
		arg5 atc.Version
		arg6 <-chan os.Signal
		arg7 chan<- struct{}
	}
	getReturns struct {
		result1 resource.VersionedSource
		result2 error
	}
	PutStub        func(resource.IOConfig, atc.Source, atc.Params, worker.ArtifactSource, <-chan os.Signal, chan<- struct{}) (resource.VersionedSource, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 resource.IOConfig
		arg2 atc.Source
		arg3 atc.Params
		arg4 worker.ArtifactSource
		arg5 <-chan os.Signal
		arg6 chan<- struct{}
	}
	putReturns struct {
		result1 resource.VersionedSource
		result2 error
	}
	CheckStub        func(atc.Source, atc.Version) ([]atc.Version, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 atc.Source
		arg2 atc.Version
	}
	checkReturns struct {
		result1 []atc.Version
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResource) Get(arg1 worker.Volume, arg2 resource.IOConfig, arg3 atc.Source, arg4 atc.Params, arg5 atc.Version, arg6 <-chan os.Signal, arg7 chan<- struct{}) (resource.VersionedSource, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 worker.Volume
		arg2 resource.IOConfig
		arg3 atc.Source
		arg4 atc.Params
		arg5 atc.Version
		arg6 <-chan os.Signal
		arg7 chan<- struct{}
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeResource) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeResource) GetArgsForCall(i int) (worker.Volume, resource.IOConfig, atc.Source, atc.Params, atc.Version, <-chan os.Signal, chan<- struct{}) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3, fake.getArgsForCall[i].arg4, fake.getArgsForCall[i].arg5, fake.getArgsForCall[i].arg6, fake.getArgsForCall[i].arg7
}

func (fake *FakeResource) GetReturns(result1 resource.VersionedSource, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Put(arg1 resource.IOConfig, arg2 atc.Source, arg3 atc.Params, arg4 worker.ArtifactSource, arg5 <-chan os.Signal, arg6 chan<- struct{}) (resource.VersionedSource, error) {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 resource.IOConfig
		arg2 atc.Source
		arg3 atc.Params
		arg4 worker.ArtifactSource
		arg5 <-chan os.Signal
		arg6 chan<- struct{}
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("Put", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3, arg4, arg5, arg6)
	} else {
		return fake.putReturns.result1, fake.putReturns.result2
	}
}

func (fake *FakeResource) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeResource) PutArgsForCall(i int) (resource.IOConfig, atc.Source, atc.Params, worker.ArtifactSource, <-chan os.Signal, chan<- struct{}) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1, fake.putArgsForCall[i].arg2, fake.putArgsForCall[i].arg3, fake.putArgsForCall[i].arg4, fake.putArgsForCall[i].arg5, fake.putArgsForCall[i].arg6
}

func (fake *FakeResource) PutReturns(result1 resource.VersionedSource, result2 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Check(arg1 atc.Source, arg2 atc.Version) ([]atc.Version, error) {
	fake.checkMutex.Lock()
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 atc.Source
		arg2 atc.Version
	}{arg1, arg2})
	fake.recordInvocation("Check", []interface{}{arg1, arg2})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(arg1, arg2)
	} else {
		return fake.checkReturns.result1, fake.checkReturns.result2
	}
}

func (fake *FakeResource) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeResource) CheckArgsForCall(i int) (atc.Source, atc.Version) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return fake.checkArgsForCall[i].arg1, fake.checkArgsForCall[i].arg2
}

func (fake *FakeResource) CheckReturns(result1 []atc.Version, result2 error) {
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 []atc.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResource) recordInvocation(key string, args []interface{}) {
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

var _ resource.Resource = new(FakeResource)
