package files

import (
	"fmt"
	"path/filepath"

	"sync"

	"github.com/StageAutoControl/controller/cntl"
)

const (
	fileNameSetLists        = "set_lists"
	fileNameSongs           = "songs"
	fileNameDmxDevices      = "dmx_devices"
	fileNameDmxDeviceTypes  = "dmx_device_types"
	fileNameDmxDeviceGroups = "dmx_device_groups"
	fileNameDmxScenes       = "dmx_scenes"
	fileNameDmxPresets      = "dmx_presets"
	fileNameDmxAnimations   = "dmx_animations"
	fileNameDmxTransitions  = "dmx_transitions"
)

type fileData struct {
	SetLists        []*cntl.SetList
	Songs           []*cntl.Song
	DMXScenes       []*cntl.DMXScene
	DMXPresets      []*cntl.DMXPreset
	DMXAnimations   []*cntl.DMXAnimation
	DMXTransitions  []*cntl.DMXTransition
	DMXDevices      []*cntl.DMXDevice
	DMXDeviceTypes  []*cntl.DMXDeviceType
	DMXDeviceGroups []*cntl.DMXDeviceGroup
}

// Repository is a file repository
type Repository struct {
	sync.Mutex
	dataDir string
}

// New crates a new file repository and returns it.
func New(dataDir string) *Repository {
	return &Repository{
		dataDir: dataDir,
	}
}

// Load implements cntl.Loader and loads the data from filesystem
func (r *Repository) Load() (*cntl.DataStore, error) {
	store := cntl.NewStore()
	data := new(fileData)

	if err := r.readDir(data, r.dataDir); err != nil {
		return nil, err
	}

	expandData(store, data)

	return store, nil
}

// Save stores the given DataStore to the filesystem
func (r *Repository) Save(store *cntl.DataStore) error {
	data := &fileData{}
	flattenData(store, data)

	return r.writeDir(data, r.dataDir)
}

func makefileTargets(data *fileData) map[string]interface{} {
	fileTargets := map[string]interface{}{
		fileNameSetLists:        &data.SetLists,
		fileNameSongs:           &data.Songs,
		fileNameDmxDevices:      &data.DMXDevices,
		fileNameDmxDeviceTypes:  &data.DMXDeviceTypes,
		fileNameDmxDeviceGroups: &data.DMXDeviceGroups,
		fileNameDmxScenes:       &data.DMXScenes,
		fileNameDmxPresets:      &data.DMXPresets,
		fileNameDmxAnimations:   &data.DMXAnimations,
		fileNameDmxTransitions:  &data.DMXTransitions,
	}
	return fileTargets
}

func getFileName(dir, fileName string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.json", fileName))
}
