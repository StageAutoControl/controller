package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"encoding/json"

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
)

type fileData struct {
	SetLists        []*cntl.SetList
	Songs           []*cntl.Song
	DMXScenes       []*cntl.DMXScene
	DMXPresets      []*cntl.DMXPreset
	DMXAnimations   []*cntl.DMXAnimation
	DMXDevices      []*cntl.DMXDevice
	DMXDeviceTypes  []*cntl.DMXDeviceType
	DMXDeviceGroups []*cntl.DMXDeviceGroup
}

// Repository is a file repository
type Repository struct {
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
	data := &fileData{}

	if err := r.readDir(data, r.dataDir); err != nil {
		return nil, err
	}

	r.mergeData(store, data)

	return store, nil
}

func (r *Repository) readDir(data *fileData, dir string) error {
	fileTargets := map[string]interface{}{
		fileNameSetLists:        &data.SetLists,
		fileNameSongs:           &data.Songs,
		fileNameDmxDevices:      &data.DMXDevices,
		fileNameDmxDeviceTypes:  &data.DMXDeviceTypes,
		fileNameDmxDeviceGroups: &data.DMXDeviceGroups,
		fileNameDmxScenes:       &data.DMXScenes,
		fileNameDmxPresets:      &data.DMXPresets,
		fileNameDmxAnimations:   &data.DMXAnimations,
	}

	for fileName, target := range fileTargets {
		file := getFileName(dir, fileName)

		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("Expected to find %q but does not exist.", file)
		} else if err != nil {
			return fmt.Errorf("Error checking file %q: %v", file, err)
		}

		if err := r.readFile(file, target); err != nil {
			return fmt.Errorf("Error reading file %q: %v", file, err)
		}
	}

	return nil
}
func getFileName(dir, fileName string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.json", fileName))
}

func (r *Repository) readFile(file string, target interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Error reading file %q: %v", file, err)
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return fmt.Errorf("Unable to parse content of %q: %v", file, err)
	}

	return nil
}

func (r *Repository) writeFile(file string, content interface{}) error {
	b, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(file, b, 0755); err != nil {
		return err
	}

	return nil
}

func (r *Repository) mergeData(data *cntl.DataStore, fileData *fileData) {
	for _, sl := range fileData.SetLists {
		data.SetLists[sl.ID] = sl
	}

	for _, s := range fileData.Songs {
		data.Songs[s.ID] = s
	}

	for _, d := range fileData.DMXDevices {
		data.DMXDevices[d.ID] = d
	}

	for _, dg := range fileData.DMXDeviceGroups {
		data.DMXDeviceGroups[dg.ID] = dg
	}

	for _, dt := range fileData.DMXDeviceTypes {
		data.DMXDeviceTypes[dt.ID] = dt
	}

	for _, p := range fileData.DMXPresets {
		data.DMXPresets[p.ID] = p
	}

	for _, sc := range fileData.DMXScenes {
		data.DMXScenes[sc.ID] = sc
	}

	for _, a := range fileData.DMXAnimations {
		data.DMXAnimations[a.ID] = a
	}
}
