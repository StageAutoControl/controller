package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/StageAutoControl/controller/cntl"
)

func (r *Repository) readDir(data *fileData, dir string) error {
	fileTargets := makefileTargets(data)

	for fileName, target := range fileTargets {
		file := getFileName(dir, fileName)

		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("expected to find %q but does not exist.", file)
		} else if err != nil {
			return fmt.Errorf("error checking file %q: %v", file, err)
		}

		if err := r.readFile(file, target); err != nil {
			return fmt.Errorf("error reading file %q: %v", file, err)
		}
	}

	return nil
}

func (r *Repository) readFile(file string, target interface{}) error {
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
}
