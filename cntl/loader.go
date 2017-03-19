package cntl

// A Loader is responsible for loading the applications data. This could either be a remote or a local store.
type Loader interface {
	Load() (*DataStore, error)
}

// A DataStore holds the applications data state
type DataStore struct {
	SetLists        map[string]*SetList
	Songs           map[string]*Song
	DMXScenes       map[string]*DMXScene
	DMXPresets      map[string]*DMXPreset
	DMXAnimations   map[string]*DMXAnimation
	DMXDevices      map[string]*DMXDevice
	DMXDeviceTypes  map[string]*DMXDeviceType
	DMXDeviceGroups map[string]*DMXDeviceGroup
}

// NewStore creates a new DataStore instance
func NewStore() *DataStore {
	return &DataStore{
		SetLists:        make(map[string]*SetList),
		Songs:           make(map[string]*Song),
		DMXScenes:       make(map[string]*DMXScene),
		DMXPresets:      make(map[string]*DMXPreset),
		DMXAnimations:   make(map[string]*DMXAnimation),
		DMXDevices:      make(map[string]*DMXDevice),
		DMXDeviceTypes:  make(map[string]*DMXDeviceType),
		DMXDeviceGroups: make(map[string]*DMXDeviceGroup),
	}
}
