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

	for _, id := range l.storage.List(&cntl.SetList{}) {
		setList := &cntl.SetList{}
		err := l.storage.Read(id, setList)
		if err != nil {
			return nil, err
		}

		data.SetLists[id] = setList
	}

	for _, id := range l.storage.List(&cntl.Song{}) {
		song := &cntl.Song{}
		err := l.storage.Read(id, song)
		if err != nil {
			return nil, err
		}

		data.Songs[id] = song
	}

	for _, id := range l.storage.List(&cntl.DMXDevice{}) {
		dmxDevice := &cntl.DMXDevice{}
		err := l.storage.Read(id, dmxDevice)
		if err != nil {
			return nil, err
		}

		data.DMXDevices[id] = dmxDevice
	}

	for _, id := range l.storage.List(&cntl.DMXDeviceGroup{}) {
		dmxDeviceGroup := &cntl.DMXDeviceGroup{}
		err := l.storage.Read(id, dmxDeviceGroup)
		if err != nil {
			return nil, err
		}

		data.DMXDeviceGroups[id] = dmxDeviceGroup
	}

	for _, id := range l.storage.List(&cntl.DMXDeviceType{}) {
		dmxDeviceType := &cntl.DMXDeviceType{}
		err := l.storage.Read(id, dmxDeviceType)
		if err != nil {
			return nil, err
		}

		data.DMXDeviceTypes[id] = dmxDeviceType
	}

	for _, id := range l.storage.List(&cntl.DMXPreset{}) {
		dmxPreset := &cntl.DMXPreset{}
		err := l.storage.Read(id, dmxPreset)
		if err != nil {
			return nil, err
		}

		data.DMXPresets[id] = dmxPreset
	}

	for _, id := range l.storage.List(&cntl.DMXScene{}) {
		dmxScene := &cntl.DMXScene{}
		err := l.storage.Read(id, dmxScene)
		if err != nil {
			return nil, err
		}

		data.DMXScenes[id] = dmxScene
	}

	for _, id := range l.storage.List(&cntl.DMXAnimation{}) {
		dmxAnimation := &cntl.DMXAnimation{}
		err := l.storage.Read(id, dmxAnimation)
		if err != nil {
			return nil, err
		}

		data.DMXAnimations[id] = dmxAnimation
	}

	for _, id := range l.storage.List(&cntl.DMXTransition{}) {
		dmxTransition := &cntl.DMXTransition{}
		err := l.storage.Read(id, dmxTransition)
		if err != nil {
			return nil, err
		}

		data.DMXTransitions[id] = dmxTransition
	}

	for _, id := range l.storage.List(&cntl.DMXColorVariable{}) {
		dmxColorVariable := &cntl.DMXColorVariable{}
		err := l.storage.Read(id, dmxColorVariable)
		if err != nil {
			return nil, err
		}

		data.DMXColorVariables[id] = dmxColorVariable
	}

	return data, nil
}
