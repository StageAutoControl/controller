package enhance

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

func init() {
	Enhancers = append(Enhancers, &NameToIDEnhancer{})
}

// NameToIDEnhancer enhances the given data by resolving names to IDs
type NameToIDEnhancer struct{}

// Enhance implements the cntl.Enhancer interface, executes the single entity enhancement methods
func (e *NameToIDEnhancer) Enhance(store *cntl.DataStore) []error {
	errs := make([]error, 0)

	errs = append(errs, e.setLists(store)...)
	errs = append(errs, e.songs(store)...)
	errs = append(errs, e.scenes(store)...)
	errs = append(errs, e.presets(store)...)
	errs = append(errs, e.deviceGroups(store)...)

	return errs
}

func (e *NameToIDEnhancer) setLists(store *cntl.DataStore) []error {
	errs := make([]error, 0)

	for _, s := range store.SetLists {
		for i := range s.Songs {
			if s.Songs[i].Name != "" {
				s.Songs[i].ID = e.findSong(s.Songs[i].Name, store.Songs)
				if s.Songs[i].ID == "" {
					errs = append(errs, fmt.Errorf("cannot find Song %q", s.Songs[i].Name))
				}
			}
		}
	}

	return errs
}

func (e *NameToIDEnhancer) songs(store *cntl.DataStore) []error {
	errs := make([]error, 0)

	for _, s := range store.Songs {
		for i := range s.DMXScenes {
			if s.DMXScenes[i].Name != "" {
				s.DMXScenes[i].ID = e.findScene(s.DMXScenes[i].Name, store.DMXScenes)
				if s.DMXScenes[i].ID == "" {
					errs = append(errs, fmt.Errorf("cannot find scene %q", s.DMXScenes[i].Name))
				}
			}
		}
	}

	return errs
}

func (e *NameToIDEnhancer) scenes(store *cntl.DataStore) []error {
	errs := make([]error, 0)

	for _, s := range store.DMXScenes {
		for i := range s.SubScenes {
			for p := range s.SubScenes[i].DeviceParams {
				errs = append(errs, e.nestedDeviceParams(&s.SubScenes[i].DeviceParams[p], store)...)
			}

			if s.SubScenes[i].Preset != nil && s.SubScenes[i].Preset.Name != "" {
				s.SubScenes[i].Preset.ID = e.findPreset(s.SubScenes[i].Preset.Name, store.DMXPresets)
				if s.SubScenes[i].Preset.ID == "" {
					errs = append(errs, fmt.Errorf("cannot find preset %q", s.SubScenes[i].Preset.Name))
				}
			}
		}
	}

	return errs
}

func (e *NameToIDEnhancer) presets(store *cntl.DataStore) []error {
	errs := make([]error, 0)

	for _, s := range store.DMXPresets {
		for i := range s.DeviceParams {
			errs = append(errs, e.nestedDeviceParams(&s.DeviceParams[i], store)...)
		}
	}

	return errs
}

func (e *NameToIDEnhancer) nestedDeviceParams(params *cntl.DMXDeviceParams, store *cntl.DataStore) []error {
	errs := make([]error, 0)

	if params.Device != nil && params.Device.ID == "" {
		params.Device.ID = e.findDevice(params.Device.Name, store.DMXDevices)
		if params.Device.ID == "" {
			errs = append(errs, fmt.Errorf("cannot find device %q", params.Device.Name))
		}
	}

	if params.Group != nil && params.Group.ID == "" {
		params.Group.ID = e.findDeviceGroup(params.Group.Name, store.DMXDeviceGroups)
		if params.Group.ID == "" {
			errs = append(errs, fmt.Errorf("cannot find device group %q", params.Group.Name))
		}
	}

	if params.Animation != nil && params.Animation.ID == "" {
		params.Animation.ID = e.findAnimation(params.Animation.Name, store.DMXAnimations)
		if params.Animation.ID == "" {
			errs = append(errs, fmt.Errorf("cannot find animation %q", params.Animation.Name))
		}
	}

	if params.Transition != nil && params.Transition.ID == "" {
		params.Transition.ID = e.findTransition(params.Transition.Name, store.DMXTransitions)
		if params.Transition.ID == "" {
			errs = append(errs, fmt.Errorf("cannot find transition %q", params.Transition.Name))
		}
	}

	return errs
}

func (e *NameToIDEnhancer) deviceGroups(store *cntl.DataStore) []error {
	errs := make([]error, 0)

	for _, s := range store.DMXDeviceGroups {
		for i := range s.Devices {
			if s.Devices[i].Name != "" {
				s.Devices[i].ID = e.findDevice(s.Devices[i].Name, store.DMXDevices)
				if s.Devices[i].ID == "" {
					errs = append(errs, fmt.Errorf("cannot find device %q", s.Devices[i].Name))
				}
			}
		}
	}

	return errs
}

func (e *NameToIDEnhancer) findSong(name string, values map[string]*cntl.Song) string {
	for _, v := range values {
		if v.Name == name {
			return v.ID
		}
	}

	return ""
}

func (e *NameToIDEnhancer) findScene(name string, values map[string]*cntl.DMXScene) string {
	for _, v := range values {
		if v.Name == name {
			return v.ID
		}
	}

	return ""
}

func (e *NameToIDEnhancer) findPreset(name string, values map[string]*cntl.DMXPreset) string {
	for _, v := range values {
		if v.Name == name {
			return v.ID
		}
	}

	return ""
}

func (e *NameToIDEnhancer) findDevice(name string, values map[string]*cntl.DMXDevice) string {
	for _, v := range values {
		if v.Name == name {
			return v.ID
		}
	}

	return ""
}

func (e *NameToIDEnhancer) findDeviceGroup(name string, values map[string]*cntl.DMXDeviceGroup) string {
	for _, v := range values {
		if v.Name == name {
			return v.ID
		}
	}

	return ""
}

func (e *NameToIDEnhancer) findAnimation(name string, values map[string]*cntl.DMXAnimation) string {
	for _, v := range values {
		if v.Name == name {
			return v.ID
		}
	}

	return ""
}

func (e *NameToIDEnhancer) findTransition(name string, values map[string]*cntl.DMXTransition) string {
	for _, v := range values {
		if v.Name == name {
			return v.ID
		}
	}

	return ""
}
