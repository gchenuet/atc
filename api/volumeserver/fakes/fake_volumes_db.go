// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/api/volumeserver"
	"github.com/concourse/atc/db"
)

type FakeVolumesDB struct {
	GetVolumesStub        func() ([]db.SavedVolume, error)
	getVolumesMutex       sync.RWMutex
	getVolumesArgsForCall []struct{}
	getVolumesReturns     struct {
		result1 []db.SavedVolume
		result2 error
	}
}

func (fake *FakeVolumesDB) GetVolumes() ([]db.SavedVolume, error) {
	fake.getVolumesMutex.Lock()
	fake.getVolumesArgsForCall = append(fake.getVolumesArgsForCall, struct{}{})
	fake.getVolumesMutex.Unlock()
	if fake.GetVolumesStub != nil {
		return fake.GetVolumesStub()
	} else {
		return fake.getVolumesReturns.result1, fake.getVolumesReturns.result2
	}
}

func (fake *FakeVolumesDB) GetVolumesCallCount() int {
	fake.getVolumesMutex.RLock()
	defer fake.getVolumesMutex.RUnlock()
	return len(fake.getVolumesArgsForCall)
}

func (fake *FakeVolumesDB) GetVolumesReturns(result1 []db.SavedVolume, result2 error) {
	fake.GetVolumesStub = nil
	fake.getVolumesReturns = struct {
		result1 []db.SavedVolume
		result2 error
	}{result1, result2}
}

var _ volumeserver.VolumesDB = new(FakeVolumesDB)
