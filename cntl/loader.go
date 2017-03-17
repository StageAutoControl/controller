package cntl

import "github.com/StageAutoControl/controller/cntl/dmx"

// A Loader is responsible for loading the applications data. This could either be a remote or a local store.
type Loader interface {
	Load() (*DataStore, error)
}

// A DataStore holds the applications data state
type DataStore struct {
	SetLists        map[string]*SetList
	Songs           map[string]*Song
	DmxScenes       map[string]*dmx.Scene
	DmxPresets      map[string]*dmx.Preset
	DmxAnimations   map[string]*dmx.Animation
	DmxDevices      map[string]*dmx.Device
	DmxDeviceGroups map[string]*dmx.DeviceGroup
}

// NewStore creates a new DataStore instance
func NewStore() *DataStore {
	return &DataStore{
		SetLists:        make(map[string]*SetList),
		Songs:           make(map[string]*Song),
		DmxScenes:       make(map[string]*dmx.Scene),
		DmxPresets:      make(map[string]*dmx.Preset),
		DmxAnimations:   make(map[string]*dmx.Animation),
		DmxDevices:      make(map[string]*dmx.Device),
		DmxDeviceGroups: make(map[string]*dmx.DeviceGroup),
	}
}
