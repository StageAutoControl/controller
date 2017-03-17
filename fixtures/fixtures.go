package fixtures

import (
	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/cntl/dmx"
)

var data = &cntl.DataStore{
	SetLists: map[string]*cntl.SetList{
		"f5b4be8a-0b18-11e7-b837-4bac99d86956": {
			ID:   "f5b4be8a-0b18-11e7-b837-4bac99d86956",
			Name: "Regular gig",
			Songs: []*cntl.SongSelector{
				{ID: "3c1065c8-0b14-11e7-96eb-5b134621c411"},
			},
		},
	},
	Songs: map[string]*cntl.Song{
		"3c1065c8-0b14-11e7-96eb-5b134621c411": {
			ID:   "3c1065c8-0b14-11e7-96eb-5b134621c411",
			Name: "Test song",
			BarChanges: []*cntl.BarChange{
				{At: 1, NoteValue: 4, NoteCount: 4, Speed: 160},
			},
			DmxScenes: []*cntl.ScenePosition{
				{ID: "492cef2e-0b14-11e7-be89-c3fa25f9cabb", Start: 1, Length: 4},
				{ID: "a44f8dee-0b14-11e7-b5b9-bf1015384192", Start: 5, Length: 4},
			},
		},
	},
	DmxPresets: map[string]*dmx.Preset{
		"5d3a415a-0b15-11e7-90b9-03c2b960e034": {
			ID:   "5d3a415a-0b15-11e7-90b9-03c2b960e034",
			Name: "strobe",
			DeviceParams: []*dmx.DeviceParams{
				{
					Group:  &dmx.DeviceGroupSelector{ID: "cb58bc10-0b16-11e7-b45a-7bee591b0adb"},
					Params: &dmx.Params{Strobe: 255},
				},
			},
		},
		"4e3c2e84-0b15-11e7-a076-4b5bbb4c19bf": {
			ID:   "4e3c2e84-0b15-11e7-a076-4b5bbb4c19bf",
			Name: "Left red",
			DeviceParams: []*dmx.DeviceParams{
				{
					Group:  &dmx.DeviceGroupSelector{ID: "475b71a0-0b16-11e7-9406-e3f678e8b788"},
					Params: &dmx.Params{Red: 200},
				},
			},
		},
	},
	DmxScenes: map[string]*dmx.Scene{
		"492cef2e-0b14-11e7-be89-c3fa25f9cabb": {
			ID:        "492cef2e-0b14-11e7-be89-c3fa25f9cabb",
			Name:      "Test-Scene",
			NoteCount: 4,
			NoteValue: 4,
			SubScenes: []*dmx.SubScene{
				{At: []uint8{1}, Preset: "4e3c2e84-0b15-11e7-a076-4b5bbb4c19bf"},
				{At: []uint8{2}, Preset: "5d3a415a-0b15-11e7-90b9-03c2b960e034"},
			},
		},
		"a44f8dee-0b14-11e7-b5b9-bf1015384192": {
			ID:        "a44f8dee-0b14-11e7-b5b9-bf1015384192",
			Name:      "Second test scene",
			NoteCount: 4,
			NoteValue: 4,
			SubScenes: []*dmx.SubScene{
				{At: []uint8{1}, Preset: "5d3a415a-0b15-11e7-90b9-03c2b960e034"},
				{At: []uint8{2}, Preset: "4e3c2e84-0b15-11e7-a076-4b5bbb4c19bf"},
			},
		},
	},
	DmxDeviceGroups: map[string]*dmx.DeviceGroup{
		"475b71a0-0b16-11e7-9406-e3f678e8b788": {
			ID:   "475b71a0-0b16-11e7-9406-e3f678e8b788",
			Name: "All PARs on the left side",
			Devices: []*dmx.DeviceSelector{
				{
					Tags: []dmx.Tag{"par", "left"},
				},
			},
		},
		"29f7adf8-0b17-11e7-bd45-9f82a70b477b": {
			ID:   "29f7adf8-0b17-11e7-bd45-9f82a70b477b",
			Name: "All PARs on the right side",
			Devices: []*dmx.DeviceSelector{
				{
					Tags: []dmx.Tag{"par", "right"},
				},
			},
		},
		"cb58bc10-0b16-11e7-b45a-7bee591b0adb": {
			ID:   "cb58bc10-0b16-11e7-b45a-7bee591b0adb",
			Name: "LED Bar infront the drums",
			Devices: []*dmx.DeviceSelector{
				{
					ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
				},
			},
		},
	},
	DmxDevices: map[string]*dmx.Device{
		"35cae00a-0b17-11e7-8bca-bbf30c56f20e": {
			ID:             "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
			Name:           "LED-Bar below drums front",
			Type:           "led-bar",
			Universe:       1,
			StartAddress:   222,
			AddressRange:   52,
			ChannelsPerLED: 3,
			Tags:           []dmx.Tag{"bar", "drums"},
		},
		"s429fc37c-0b17-11e7-8b94-c3b6519355d3": {
			ID:           "s429fc37c-0b17-11e7-8b94-c3b6519355d3",
			Name:         "PAR inner left, stand 1 position 1",
			Type:         "par",
			Universe:     2,
			StartAddress: 10,
			AddressRange: 4,
			Tags:         []dmx.Tag{"par", "left", "inner", "stand", "drums-left"},
		},
		"4a545466-0b17-11e7-9c61-d3c0693099ab": {
			ID:           "4a545466-0b17-11e7-9c61-d3c0693099ab",
			Name:         "PAR inner left, stand 1 position 2",
			Type:         "par",
			Universe:     2,
			StartAddress: 14,
			AddressRange: 4,
			Tags:         []dmx.Tag{"par", "left", "inner", "stand", "drums-left"},
		},
		"5e0335e0-0b17-11e7-ad6c-63a7138d926c": {
			ID:           "5e0335e0-0b17-11e7-ad6c-63a7138d926c",
			Name:         "PAR inner right, stand 2 position 1",
			Type:         "par",
			Universe:     2,
			StartAddress: 26,
			AddressRange: 4,
			Tags:         []dmx.Tag{"par", "right", "inner", "stand", "drums-right"},
		},
		"620101f4-0b17-11e7-85cc-539952d9aef2": {
			ID:           "620101f4-0b17-11e7-85cc-539952d9aef2",
			Name:         "PAR inner right, stand 2 position 2",
			Type:         "par",
			Universe:     2,
			StartAddress: 30,
			AddressRange: 4,
			Tags:         []dmx.Tag{"par", "right", "inner", "stand", "drums-right"},
		},
		"6f7bca8a-0b17-11e7-b604-a356da737e54": {
			ID:           "6f7bca8a-0b17-11e7-b604-a356da737e54",
			Name:         "Strobe Vocs",
			Type:         "strobe",
			Universe:     1,
			StartAddress: 202,
			AddressRange: 4,
			Tags:         []dmx.Tag{"strobe-back", "vocs"},
		},
	},
}

// DataStore returns the go object representation of a working set of fixtures
func DataStore() *cntl.DataStore {
	return data
}
