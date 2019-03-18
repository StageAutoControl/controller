package files

import (
	"github.com/StageAutoControl/controller/pkg/cntl"
)

type fileData struct {
	SetLists          []*cntl.SetList          `json:"set_lists"`
	Songs             []*cntl.Song             `json:"songs"`
	DMXScenes         []*cntl.DMXScene         `json:"dmx_scenes"`
	DMXPresets        []*cntl.DMXPreset        `json:"dmx_presets"`
	DMXAnimations     []*cntl.DMXAnimation     `json:"dmx_animations"`
	DMXTransitions    []*cntl.DMXTransition    `json:"dmx_transitions"`
	DMXDevices        []*cntl.DMXDevice        `json:"dmx_devices"`
	DMXDeviceTypes    []*cntl.DMXDeviceType    `json:"dmx_device_types"`
	DMXDeviceGroups   []*cntl.DMXDeviceGroup   `json:"dmx_device_groups"`
	DMXColorVariables []*cntl.DMXColorVariable `json:"dmx_color_variables"`
}

// Database is a file repository
type Database struct {
	dataDir string
}

// New crates a new file repository and returns it.
func New(dataDir string) *Database {
	return &Database{
		dataDir: dataDir,
	}
}

// Load implements cntl.Loader and loads the data from filesystem
func (d *Database) Load() (*cntl.DataStore, error) {
	store := cntl.NewStore()
	data, err := d.readDir(d.dataDir)

	if err != nil {
		return nil, err
	}

	expandData(store, data)

	return store, nil
}
