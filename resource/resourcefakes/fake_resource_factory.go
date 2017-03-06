// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeResourceFactory struct {
	NewBuildResourceStub        func(logger lager.Logger, id worker.Identifier, metadata worker.Metadata, containerSpec worker.ContainerSpec, resourceTypes dbng.ResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate, inputSources []resource.InputSource, outputPaths map[string]string) (resource.Resource, []resource.InputSource, error)
	newBuildResourceMutex       sync.RWMutex
	newBuildResourceArgsForCall []struct {
		logger                lager.Logger
		id                    worker.Identifier
		metadata              worker.Metadata
		containerSpec         worker.ContainerSpec
		resourceTypes         dbng.ResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
		inputSources          []resource.InputSource
		outputPaths           map[string]string
	}
	newBuildResourceReturns struct {
		result1 resource.Resource
		result2 []resource.InputSource
		result3 error
	}
	NewCheckResourceStub        func(logger lager.Logger, resourceUser dbng.ResourceUser, id worker.Identifier, metadata worker.Metadata, resourceSpec worker.ContainerSpec, resourceTypes dbng.ResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate, resourceConfig atc.ResourceConfig) (resource.Resource, error)
	newCheckResourceMutex       sync.RWMutex
	newCheckResourceArgsForCall []struct {
		logger                lager.Logger
		resourceUser          dbng.ResourceUser
		id                    worker.Identifier
		metadata              worker.Metadata
		resourceSpec          worker.ContainerSpec
		resourceTypes         dbng.ResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
		resourceConfig        atc.ResourceConfig
	}
	newCheckResourceReturns struct {
		result1 resource.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactory) NewBuildResource(logger lager.Logger, id worker.Identifier, metadata worker.Metadata, containerSpec worker.ContainerSpec, resourceTypes dbng.ResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate, inputSources []resource.InputSource, outputPaths map[string]string) (resource.Resource, []resource.InputSource, error) {
	var inputSourcesCopy []resource.InputSource
	if inputSources != nil {
		inputSourcesCopy = make([]resource.InputSource, len(inputSources))
		copy(inputSourcesCopy, inputSources)
	}
	fake.newBuildResourceMutex.Lock()
	fake.newBuildResourceArgsForCall = append(fake.newBuildResourceArgsForCall, struct {
		logger                lager.Logger
		id                    worker.Identifier
		metadata              worker.Metadata
		containerSpec         worker.ContainerSpec
		resourceTypes         dbng.ResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
		inputSources          []resource.InputSource
		outputPaths           map[string]string
	}{logger, id, metadata, containerSpec, resourceTypes, imageFetchingDelegate, inputSourcesCopy, outputPaths})
	fake.recordInvocation("NewBuildResource", []interface{}{logger, id, metadata, containerSpec, resourceTypes, imageFetchingDelegate, inputSourcesCopy, outputPaths})
	fake.newBuildResourceMutex.Unlock()
	if fake.NewBuildResourceStub != nil {
		return fake.NewBuildResourceStub(logger, id, metadata, containerSpec, resourceTypes, imageFetchingDelegate, inputSources, outputPaths)
	}
	return fake.newBuildResourceReturns.result1, fake.newBuildResourceReturns.result2, fake.newBuildResourceReturns.result3
}

func (fake *FakeResourceFactory) NewBuildResourceCallCount() int {
	fake.newBuildResourceMutex.RLock()
	defer fake.newBuildResourceMutex.RUnlock()
	return len(fake.newBuildResourceArgsForCall)
}

func (fake *FakeResourceFactory) NewBuildResourceArgsForCall(i int) (lager.Logger, worker.Identifier, worker.Metadata, worker.ContainerSpec, dbng.ResourceTypes, worker.ImageFetchingDelegate, []resource.InputSource, map[string]string) {
	fake.newBuildResourceMutex.RLock()
	defer fake.newBuildResourceMutex.RUnlock()
	return fake.newBuildResourceArgsForCall[i].logger, fake.newBuildResourceArgsForCall[i].id, fake.newBuildResourceArgsForCall[i].metadata, fake.newBuildResourceArgsForCall[i].containerSpec, fake.newBuildResourceArgsForCall[i].resourceTypes, fake.newBuildResourceArgsForCall[i].imageFetchingDelegate, fake.newBuildResourceArgsForCall[i].inputSources, fake.newBuildResourceArgsForCall[i].outputPaths
}

func (fake *FakeResourceFactory) NewBuildResourceReturns(result1 resource.Resource, result2 []resource.InputSource, result3 error) {
	fake.NewBuildResourceStub = nil
	fake.newBuildResourceReturns = struct {
		result1 resource.Resource
		result2 []resource.InputSource
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceFactory) NewCheckResource(logger lager.Logger, resourceUser dbng.ResourceUser, id worker.Identifier, metadata worker.Metadata, resourceSpec worker.ContainerSpec, resourceTypes dbng.ResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate, resourceConfig atc.ResourceConfig) (resource.Resource, error) {
	fake.newCheckResourceMutex.Lock()
	fake.newCheckResourceArgsForCall = append(fake.newCheckResourceArgsForCall, struct {
		logger                lager.Logger
		resourceUser          dbng.ResourceUser
		id                    worker.Identifier
		metadata              worker.Metadata
		resourceSpec          worker.ContainerSpec
		resourceTypes         dbng.ResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
		resourceConfig        atc.ResourceConfig
	}{logger, resourceUser, id, metadata, resourceSpec, resourceTypes, imageFetchingDelegate, resourceConfig})
	fake.recordInvocation("NewCheckResource", []interface{}{logger, resourceUser, id, metadata, resourceSpec, resourceTypes, imageFetchingDelegate, resourceConfig})
	fake.newCheckResourceMutex.Unlock()
	if fake.NewCheckResourceStub != nil {
		return fake.NewCheckResourceStub(logger, resourceUser, id, metadata, resourceSpec, resourceTypes, imageFetchingDelegate, resourceConfig)
	}
	return fake.newCheckResourceReturns.result1, fake.newCheckResourceReturns.result2
}

func (fake *FakeResourceFactory) NewCheckResourceCallCount() int {
	fake.newCheckResourceMutex.RLock()
	defer fake.newCheckResourceMutex.RUnlock()
	return len(fake.newCheckResourceArgsForCall)
}

func (fake *FakeResourceFactory) NewCheckResourceArgsForCall(i int) (lager.Logger, dbng.ResourceUser, worker.Identifier, worker.Metadata, worker.ContainerSpec, dbng.ResourceTypes, worker.ImageFetchingDelegate, atc.ResourceConfig) {
	fake.newCheckResourceMutex.RLock()
	defer fake.newCheckResourceMutex.RUnlock()
	return fake.newCheckResourceArgsForCall[i].logger, fake.newCheckResourceArgsForCall[i].resourceUser, fake.newCheckResourceArgsForCall[i].id, fake.newCheckResourceArgsForCall[i].metadata, fake.newCheckResourceArgsForCall[i].resourceSpec, fake.newCheckResourceArgsForCall[i].resourceTypes, fake.newCheckResourceArgsForCall[i].imageFetchingDelegate, fake.newCheckResourceArgsForCall[i].resourceConfig
}

func (fake *FakeResourceFactory) NewCheckResourceReturns(result1 resource.Resource, result2 error) {
	fake.NewCheckResourceStub = nil
	fake.newCheckResourceReturns = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newBuildResourceMutex.RLock()
	defer fake.newBuildResourceMutex.RUnlock()
	fake.newCheckResourceMutex.RLock()
	defer fake.newCheckResourceMutex.RUnlock()
	return fake.invocations
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
