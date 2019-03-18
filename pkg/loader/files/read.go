package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/StageAutoControl/controller/pkg/cntl"
)

func (d *Database) readDir(dir string) (*fileData, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	data := new(fileData)
	for _, file := range files {
		file := filepath.Join(dir, file.Name())
		fData := new(fileData)
		if err := d.readFile(file, fData); err != nil {
			return nil, fmt.Errorf("error reading file %q: %v", file, err)
		}

		data = mergeFileData(data, fData)
	}

	return data, nil
}

func (d *Database) readFile(file string, target interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading file %q: %v", file, err)
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return fmt.Errorf("unable to parse content of %q: %v", file, err)
	}

	return nil
}

func mergeFileData(data, fd *fileData) *fileData {
	newData := new(fileData)

	newData.SetLists = append(data.SetLists, fd.SetLists...)
	newData.Songs = append(data.Songs, fd.Songs...)
	newData.DMXScenes = append(data.DMXScenes, fd.DMXScenes...)
	newData.DMXPresets = append(data.DMXPresets, fd.DMXPresets...)
	newData.DMXAnimations = append(data.DMXAnimations, fd.DMXAnimations...)
	newData.DMXTransitions = append(data.DMXTransitions, fd.DMXTransitions...)
	newData.DMXDevices = append(data.DMXDevices, fd.DMXDevices...)
	newData.DMXDeviceTypes = append(data.DMXDeviceTypes, fd.DMXDeviceTypes...)
	newData.DMXDeviceGroups = append(data.DMXDeviceGroups, fd.DMXDeviceGroups...)
	newData.DMXColorVariables = append(data.DMXColorVariables, fd.DMXColorVariables...)

	return newData
}

func expandData(data *cntl.DataStore, fileData *fileData) {
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

	for _, t := range fileData.DMXTransitions {
		data.DMXTransitions[t.ID] = t
	}

	for _, t := range fileData.DMXColorVariables {
		data.DMXColorVariables[t.ID] = t
	}
}
