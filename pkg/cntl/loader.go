package cntl

// A DataStore holds the controllers data state during playback, or more specifically during the rendering of a song into DMX frames
type DataStore struct {
	SetLists          map[string]*SetList
	Songs             map[string]*Song
	DMXScenes         map[string]*DMXScene
	DMXPresets        map[string]*DMXPreset
	DMXAnimations     map[string]*DMXAnimation
	DMXTransitions    map[string]*DMXTransition
	DMXDevices        map[string]*DMXDevice
	DMXDeviceTypes    map[string]*DMXDeviceType
	DMXDeviceGroups   map[string]*DMXDeviceGroup
	DmxColorVariables map[string]*DMXColorVariable
}

// NewStore creates a new DataStore instance
func NewStore() *DataStore {
	return &DataStore{
		SetLists:          make(map[string]*SetList),
		Songs:             make(map[string]*Song),
		DMXScenes:         make(map[string]*DMXScene),
		DMXPresets:        make(map[string]*DMXPreset),
		DMXAnimations:     make(map[string]*DMXAnimation),
		DMXTransitions:    make(map[string]*DMXTransition),
		DMXDevices:        make(map[string]*DMXDevice),
		DMXDeviceTypes:    make(map[string]*DMXDeviceType),
		DMXDeviceGroups:   make(map[string]*DMXDeviceGroup),
		DmxColorVariables: make(map[string]*DMXColorVariable),
	}
}
