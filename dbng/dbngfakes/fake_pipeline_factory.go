// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakePipelineFactory struct {
	GetPipelineByIDStub        func(teamID int, pipelineID int) dbng.Pipeline
	getPipelineByIDMutex       sync.RWMutex
	getPipelineByIDArgsForCall []struct {
		teamID     int
		pipelineID int
	}
	getPipelineByIDReturns struct {
		result1 dbng.Pipeline
	}
	getPipelineByIDReturnsOnCall map[int]struct {
		result1 dbng.Pipeline
	}
	PublicPipelinesStub        func() ([]dbng.Pipeline, error)
	publicPipelinesMutex       sync.RWMutex
	publicPipelinesArgsForCall []struct{}
	publicPipelinesReturns     struct {
		result1 []dbng.Pipeline
		result2 error
	}
	publicPipelinesReturnsOnCall map[int]struct {
		result1 []dbng.Pipeline
		result2 error
	}
	AllPipelinesStub        func() ([]dbng.Pipeline, error)
	allPipelinesMutex       sync.RWMutex
	allPipelinesArgsForCall []struct{}
	allPipelinesReturns     struct {
		result1 []dbng.Pipeline
		result2 error
	}
	allPipelinesReturnsOnCall map[int]struct {
		result1 []dbng.Pipeline
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePipelineFactory) GetPipelineByID(teamID int, pipelineID int) dbng.Pipeline {
	fake.getPipelineByIDMutex.Lock()
	ret, specificReturn := fake.getPipelineByIDReturnsOnCall[len(fake.getPipelineByIDArgsForCall)]
	fake.getPipelineByIDArgsForCall = append(fake.getPipelineByIDArgsForCall, struct {
		teamID     int
		pipelineID int
	}{teamID, pipelineID})
	fake.recordInvocation("GetPipelineByID", []interface{}{teamID, pipelineID})
	fake.getPipelineByIDMutex.Unlock()
	if fake.GetPipelineByIDStub != nil {
		return fake.GetPipelineByIDStub(teamID, pipelineID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getPipelineByIDReturns.result1
}

func (fake *FakePipelineFactory) GetPipelineByIDCallCount() int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return len(fake.getPipelineByIDArgsForCall)
}

func (fake *FakePipelineFactory) GetPipelineByIDArgsForCall(i int) (int, int) {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return fake.getPipelineByIDArgsForCall[i].teamID, fake.getPipelineByIDArgsForCall[i].pipelineID
}

func (fake *FakePipelineFactory) GetPipelineByIDReturns(result1 dbng.Pipeline) {
	fake.GetPipelineByIDStub = nil
	fake.getPipelineByIDReturns = struct {
		result1 dbng.Pipeline
	}{result1}
}

func (fake *FakePipelineFactory) GetPipelineByIDReturnsOnCall(i int, result1 dbng.Pipeline) {
	fake.GetPipelineByIDStub = nil
	if fake.getPipelineByIDReturnsOnCall == nil {
		fake.getPipelineByIDReturnsOnCall = make(map[int]struct {
			result1 dbng.Pipeline
		})
	}
	fake.getPipelineByIDReturnsOnCall[i] = struct {
		result1 dbng.Pipeline
	}{result1}
}

func (fake *FakePipelineFactory) PublicPipelines() ([]dbng.Pipeline, error) {
	fake.publicPipelinesMutex.Lock()
	ret, specificReturn := fake.publicPipelinesReturnsOnCall[len(fake.publicPipelinesArgsForCall)]
	fake.publicPipelinesArgsForCall = append(fake.publicPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("PublicPipelines", []interface{}{})
	fake.publicPipelinesMutex.Unlock()
	if fake.PublicPipelinesStub != nil {
		return fake.PublicPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.publicPipelinesReturns.result1, fake.publicPipelinesReturns.result2
}

func (fake *FakePipelineFactory) PublicPipelinesCallCount() int {
	fake.publicPipelinesMutex.RLock()
	defer fake.publicPipelinesMutex.RUnlock()
	return len(fake.publicPipelinesArgsForCall)
}

func (fake *FakePipelineFactory) PublicPipelinesReturns(result1 []dbng.Pipeline, result2 error) {
	fake.PublicPipelinesStub = nil
	fake.publicPipelinesReturns = struct {
		result1 []dbng.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) PublicPipelinesReturnsOnCall(i int, result1 []dbng.Pipeline, result2 error) {
	fake.PublicPipelinesStub = nil
	if fake.publicPipelinesReturnsOnCall == nil {
		fake.publicPipelinesReturnsOnCall = make(map[int]struct {
			result1 []dbng.Pipeline
			result2 error
		})
	}
	fake.publicPipelinesReturnsOnCall[i] = struct {
		result1 []dbng.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) AllPipelines() ([]dbng.Pipeline, error) {
	fake.allPipelinesMutex.Lock()
	ret, specificReturn := fake.allPipelinesReturnsOnCall[len(fake.allPipelinesArgsForCall)]
	fake.allPipelinesArgsForCall = append(fake.allPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("AllPipelines", []interface{}{})
	fake.allPipelinesMutex.Unlock()
	if fake.AllPipelinesStub != nil {
		return fake.AllPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allPipelinesReturns.result1, fake.allPipelinesReturns.result2
}

func (fake *FakePipelineFactory) AllPipelinesCallCount() int {
	fake.allPipelinesMutex.RLock()
	defer fake.allPipelinesMutex.RUnlock()
	return len(fake.allPipelinesArgsForCall)
}

func (fake *FakePipelineFactory) AllPipelinesReturns(result1 []dbng.Pipeline, result2 error) {
	fake.AllPipelinesStub = nil
	fake.allPipelinesReturns = struct {
		result1 []dbng.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) AllPipelinesReturnsOnCall(i int, result1 []dbng.Pipeline, result2 error) {
	fake.AllPipelinesStub = nil
	if fake.allPipelinesReturnsOnCall == nil {
		fake.allPipelinesReturnsOnCall = make(map[int]struct {
			result1 []dbng.Pipeline
			result2 error
		})
	}
	fake.allPipelinesReturnsOnCall[i] = struct {
		result1 []dbng.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	fake.publicPipelinesMutex.RLock()
	defer fake.publicPipelinesMutex.RUnlock()
	fake.allPipelinesMutex.RLock()
	defer fake.allPipelinesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePipelineFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.PipelineFactory = new(FakePipelineFactory)