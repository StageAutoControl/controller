package disk

import "github.com/StageAutoControl/controller/pkg/cntl"

// Loader loads the DataStore from the storage
type Loader struct {
	storage *Storage
}

// NewLoader returns a new Loader instance
func NewLoader(storage *Storage) *Loader {
	return &Loader{
		storage: storage,
	}
}

// Load the data from the storage and return a populated data store
func (l *Loader) Load() (*cntl.DataStore, error) {
	data := cntl.NewStore()

	setList := &cntl.SetList{}
	for _, id := range l.storage.List(setList) {
		err := l.storage.Read(id, setList)
		if err != nil {
			return nil, err
		}

		data.SetLists[id] = setList
	}

	song := &cntl.Song{}
	for _, id := range l.storage.List(song) {
		err := l.storage.Read(id, song)
		if err != nil {
			return nil, err
		}

		data.Songs[id] = song
	}

	dmxDevice := &cntl.DMXDevice{}
	for _, id := range l.storage.List(dmxDevice) {
		err := l.storage.Read(id, dmxDevice)
		if err != nil {
			return nil, err
		}

		data.DMXDevices[id] = dmxDevice
	}

	dmxDeviceGroup := &cntl.DMXDeviceGroup{}
	for _, id := range l.storage.List(dmxDeviceGroup) {
		err := l.storage.Read(id, dmxDeviceGroup)
		if err != nil {
			return nil, err
		}

		data.DMXDeviceGroups[id] = dmxDeviceGroup
	}

	dmxDeviceType := &cntl.DMXDeviceType{}
	for _, id := range l.storage.List(dmxDeviceType) {
		err := l.storage.Read(id, dmxDeviceType)
		if err != nil {
			return nil, err
		}

		data.DMXDeviceTypes[id] = dmxDeviceType
	}

	dmxPreset := &cntl.DMXPreset{}
	for _, id := range l.storage.List(dmxPreset) {
		err := l.storage.Read(id, dmxPreset)
		if err != nil {
			return nil, err
		}

		data.DMXPresets[id] = dmxPreset
	}

	dmxScene := &cntl.DMXScene{}
	for _, id := range l.storage.List(dmxScene) {
		err := l.storage.Read(id, dmxScene)
		if err != nil {
			return nil, err
		}

		data.DMXScenes[id] = dmxScene
	}

	dmxAnimation := &cntl.DMXAnimation{}
	for _, id := range l.storage.List(dmxAnimation) {
		err := l.storage.Read(id, dmxAnimation)
		if err != nil {
			return nil, err
		}

		data.DMXAnimations[id] = dmxAnimation
	}

	dmxTransition := &cntl.DMXTransition{}
	for _, id := range l.storage.List(dmxTransition) {
		err := l.storage.Read(id, dmxTransition)
		if err != nil {
			return nil, err
		}

		data.DMXTransitions[id] = dmxTransition
	}

	return data, nil
}
