package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/StageAutoControl/controller/cntl"
)

func (r *Repository) writeDir(data *fileData, dir string) error {
	fileTargets := makefileTargets(data)

	for fileName, target := range fileTargets {
		file := getFileName(dir, fileName)

		if stat, err := os.Stat(file); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("error checking file %q: %v", file, err)
		} else if stat.IsDir() {
			return fmt.Errorf("error checking file %q: It's a dir", file)
		}

		if err := r.writeFile(file, target); err != nil {
			return fmt.Errorf("error writing file %q: %v", file, err)
		}
	}

	return nil
}

func (r *Repository) writeFile(file string, content interface{}) error {
	b, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, b, 0755)
}

func flattenData(store *cntl.DataStore, fileData *fileData) {
	for _, sl := range store.SetLists {
		fileData.SetLists = append(fileData.SetLists, sl)
	}

	for _, s := range store.Songs {
		fileData.Songs = append(fileData.Songs, s)
	}

	for _, d := range store.DMXDevices {
		fileData.DMXDevices = append(fileData.DMXDevices, d)
	}

	for _, dg := range store.DMXDeviceGroups {
		fileData.DMXDeviceGroups = append(fileData.DMXDeviceGroups, dg)
	}

	for _, dt := range store.DMXDeviceTypes {
		fileData.DMXDeviceTypes = append(fileData.DMXDeviceTypes, dt)
	}

	for _, p := range store.DMXPresets {
		fileData.DMXPresets = append(fileData.DMXPresets, p)
	}

	for _, sc := range store.DMXScenes {
		fileData.DMXScenes = append(fileData.DMXScenes, sc)
	}

	for _, a := range store.DMXAnimations {
		fileData.DMXAnimations = append(fileData.DMXAnimations, a)
	}

	for _, t := range store.DMXTransitions {
		fileData.DMXTransitions = append(fileData.DMXTransitions, t)
	}
}
