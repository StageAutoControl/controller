package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"strings"

	"log"

	"github.com/StageAutoControl/controller/cntl"
	"gopkg.in/yaml.v2"
)

type fileData struct {
	SetLists        []*cntl.SetList        `json:"setLists" yaml:"setLists"`
	Songs           []*cntl.Song           `json:"songs" yaml:"songs"`
	DMXScenes       []*cntl.DMXScene       `json:"dmxScenes" yaml:"dmxScenes"`
	DMXPresets      []*cntl.DMXPreset      `json:"dmxPresets" yaml:"dmxPresets"`
	DMXAnimations   []*cntl.DMXAnimation   `json:"dmxAnimations" yaml:"dmxAnimations"`
	DMXDevices      []*cntl.DMXDevice      `json:"dmxDevices" yaml:"dmxDevices"`
	DMXDeviceTypes  []*cntl.DMXDeviceType  `json:"dmxDeviceTypes" yaml:"dmxDeviceTypes"`
	DMXDeviceGroups []*cntl.DMXDeviceGroup `json:"dmxDeviceGroups" yaml:"dmxDeviceGroups"`
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
	return r.readDir(store, r.dataDir)
}

func (r *Repository) readDir(data *cntl.DataStore, dir string) (*cntl.DataStore, error) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range fs {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		path := filepath.Join(dir, f.Name())
		if f.IsDir() {
			data, err = r.readDir(data, path)
			if err != nil {
				return nil, err
			}

			continue
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".yml", ".yaml":
			data, err = r.readYAMLFile(data, path)
			if err != nil {
				return nil, err
			}

			break

		case ".json":
			data, err = r.readJSONFile(data, path)
			if err != nil {
				return nil, err
			}

			break

		default:
			log.Printf("Unable to load file %q. No loader for file extension %q known.", path, ext)
		}
	}

	return data, nil
}

func (r *Repository) readJSONFile(data *cntl.DataStore, path string) (*cntl.DataStore, error) {
	var fileData fileData

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %q: %v", path, err)
	}

	err = json.Unmarshal(b, &fileData)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse content of %q: %v", path, err)
	}

	return r.mergeData(data, &fileData), nil
}

func (r *Repository) readYAMLFile(data *cntl.DataStore, path string) (*cntl.DataStore, error) {
	var fileData fileData

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %q: %v", path, err)
	}

	err = yaml.Unmarshal(b, &fileData)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse content of %q: %v", path, err)
	}

	return r.mergeData(data, &fileData), nil
}

func (r *Repository) mergeData(data *cntl.DataStore, fileData *fileData) *cntl.DataStore {
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

	return data
}
