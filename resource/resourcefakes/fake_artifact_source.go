// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeArtifactSource struct {
	StreamToStub        func(resource.ArtifactDestination) error
	streamToMutex       sync.RWMutex
	streamToArgsForCall []struct {
		arg1 resource.ArtifactDestination
	}
	streamToReturns struct {
		result1 error
	}
	streamToReturnsOnCall map[int]struct {
		result1 error
	}
	VolumeOnStub        func(worker.Worker) (worker.Volume, bool, error)
	volumeOnMutex       sync.RWMutex
	volumeOnArgsForCall []struct {
		arg1 worker.Worker
	}
	volumeOnReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	volumeOnReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArtifactSource) StreamTo(arg1 resource.ArtifactDestination) error {
	fake.streamToMutex.Lock()
	ret, specificReturn := fake.streamToReturnsOnCall[len(fake.streamToArgsForCall)]
	fake.streamToArgsForCall = append(fake.streamToArgsForCall, struct {
		arg1 resource.ArtifactDestination
	}{arg1})
	fake.recordInvocation("StreamTo", []interface{}{arg1})
	fake.streamToMutex.Unlock()
	if fake.StreamToStub != nil {
		return fake.StreamToStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.streamToReturns.result1
}

func (fake *FakeArtifactSource) StreamToCallCount() int {
	fake.streamToMutex.RLock()
	defer fake.streamToMutex.RUnlock()
	return len(fake.streamToArgsForCall)
}

func (fake *FakeArtifactSource) StreamToArgsForCall(i int) resource.ArtifactDestination {
	fake.streamToMutex.RLock()
	defer fake.streamToMutex.RUnlock()
	return fake.streamToArgsForCall[i].arg1
}

func (fake *FakeArtifactSource) StreamToReturns(result1 error) {
	fake.StreamToStub = nil
	fake.streamToReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactSource) StreamToReturnsOnCall(i int, result1 error) {
	fake.StreamToStub = nil
	if fake.streamToReturnsOnCall == nil {
		fake.streamToReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamToReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactSource) VolumeOn(arg1 worker.Worker) (worker.Volume, bool, error) {
	fake.volumeOnMutex.Lock()
	ret, specificReturn := fake.volumeOnReturnsOnCall[len(fake.volumeOnArgsForCall)]
	fake.volumeOnArgsForCall = append(fake.volumeOnArgsForCall, struct {
		arg1 worker.Worker
	}{arg1})
	fake.recordInvocation("VolumeOn", []interface{}{arg1})
	fake.volumeOnMutex.Unlock()
	if fake.VolumeOnStub != nil {
		return fake.VolumeOnStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.volumeOnReturns.result1, fake.volumeOnReturns.result2, fake.volumeOnReturns.result3
}

func (fake *FakeArtifactSource) VolumeOnCallCount() int {
	fake.volumeOnMutex.RLock()
	defer fake.volumeOnMutex.RUnlock()
	return len(fake.volumeOnArgsForCall)
}

func (fake *FakeArtifactSource) VolumeOnArgsForCall(i int) worker.Worker {
	fake.volumeOnMutex.RLock()
	defer fake.volumeOnMutex.RUnlock()
	return fake.volumeOnArgsForCall[i].arg1
}

func (fake *FakeArtifactSource) VolumeOnReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.VolumeOnStub = nil
	fake.volumeOnReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArtifactSource) VolumeOnReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.VolumeOnStub = nil
	if fake.volumeOnReturnsOnCall == nil {
		fake.volumeOnReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.volumeOnReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArtifactSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.streamToMutex.RLock()
	defer fake.streamToMutex.RUnlock()
	fake.volumeOnMutex.RLock()
	defer fake.volumeOnMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeArtifactSource) recordInvocation(key string, args []interface{}) {
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

var _ resource.ArtifactSource = new(FakeArtifactSource)