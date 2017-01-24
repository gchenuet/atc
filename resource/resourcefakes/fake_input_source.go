// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeInputSource struct {
	NameStub        func() worker.ArtifactName
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 worker.ArtifactName
	}
	SourceStub        func() worker.ArtifactSource
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct{}
	sourceReturns     struct {
		result1 worker.ArtifactSource
	}
	MountPathStub        func() string
	mountPathMutex       sync.RWMutex
	mountPathArgsForCall []struct{}
	mountPathReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInputSource) Name() worker.ArtifactName {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeInputSource) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeInputSource) NameReturns(result1 worker.ArtifactName) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 worker.ArtifactName
	}{result1}
}

func (fake *FakeInputSource) Source() worker.ArtifactSource {
	fake.sourceMutex.Lock()
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct{}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	} else {
		return fake.sourceReturns.result1
	}
}

func (fake *FakeInputSource) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeInputSource) SourceReturns(result1 worker.ArtifactSource) {
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 worker.ArtifactSource
	}{result1}
}

func (fake *FakeInputSource) MountPath() string {
	fake.mountPathMutex.Lock()
	fake.mountPathArgsForCall = append(fake.mountPathArgsForCall, struct{}{})
	fake.recordInvocation("MountPath", []interface{}{})
	fake.mountPathMutex.Unlock()
	if fake.MountPathStub != nil {
		return fake.MountPathStub()
	} else {
		return fake.mountPathReturns.result1
	}
}

func (fake *FakeInputSource) MountPathCallCount() int {
	fake.mountPathMutex.RLock()
	defer fake.mountPathMutex.RUnlock()
	return len(fake.mountPathArgsForCall)
}

func (fake *FakeInputSource) MountPathReturns(result1 string) {
	fake.MountPathStub = nil
	fake.mountPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInputSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.mountPathMutex.RLock()
	defer fake.mountPathMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeInputSource) recordInvocation(key string, args []interface{}) {
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

var _ resource.InputSource = new(FakeInputSource)
