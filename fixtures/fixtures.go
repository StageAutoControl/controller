package fixtures

import (
	"github.com/StageAutoControl/controller/cntl"
)

// fixture values
var (
	Value0   = &cntl.DMXValue{0}
	Value31  = &cntl.DMXValue{31}
	Value63  = &cntl.DMXValue{63}
	Value127 = &cntl.DMXValue{127}
	Value200 = &cntl.DMXValue{200}
	Value255 = &cntl.DMXValue{255}
)

var data = &cntl.DataStore{
	SetLists: map[string]*cntl.SetList{
		"f5b4be8a-0b18-11e7-b837-4bac99d86956": {
			ID:   "f5b4be8a-0b18-11e7-b837-4bac99d86956",
			Name: "Regular gig",
			Songs: []cntl.SongSelector{
				{ID: "3c1065c8-0b14-11e7-96eb-5b134621c411"},
			},
		},
	},
	Songs: map[string]*cntl.Song{
		"3c1065c8-0b14-11e7-96eb-5b134621c411": {
			ID:     "3c1065c8-0b14-11e7-96eb-5b134621c411",
			Name:   "Test song",
			Length: 3200,
			BarChanges: []cntl.BarChange{
				{At: 0, NoteCount: 4, NoteValue: 4, Speed: 160},
				{At: 512, NoteCount: 3, NoteValue: 4},
				{At: 1184, NoteCount: 7, NoteValue: 8},
				{At: 1632, NoteCount: 4, NoteValue: 4},
			},
			DMXScenes: []cntl.ScenePosition{
				{At: 0, ID: "492cef2e-0b14-11e7-be89-c3fa25f9cabb", Repeat: 4},
				{At: 512, ID: "a44f8dee-0b14-11e7-b5b9-bf1015384192", Repeat: 3},
				{At: 1408, ID: "99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6", Repeat: 2},
				{At: 1920, ID: "b82f4750-0e7a-11e7-9522-0f9d6d69958a"},
			},
		},
	},
	DMXPresets: map[string]*cntl.DMXPreset{
		"0de258e0-0e7b-11e7-afd4-ebf6036983dc": {
			ID:   "0de258e0-0e7b-11e7-afd4-ebf6036983dc",
			Name: "Test-Preset 1",
			DeviceParams: []cntl.DMXDeviceParams{
				{
					Device: &cntl.DMXDeviceSelector{
						ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
					},
					Params: &cntl.DMXParams{
						Red: Value255,
					},
				},
			}},
		"11adf93e-0e7b-11e7-998c-5bd2bd0df396": {
			ID:   "11adf93e-0e7b-11e7-998c-5bd2bd0df396",
			Name: "Test-Preset 2",
			DeviceParams: []cntl.DMXDeviceParams{
				{
					Device: &cntl.DMXDeviceSelector{
						ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
					},
					Params: &cntl.DMXParams{
						Blue: Value255,
					},
				},
			}},
		"652e716a-0e7b-11e7-b92a-8f2ff28ba235": {
			ID:   "652e716a-0e7b-11e7-b92a-8f2ff28ba235",
			Name: "Test-Preset 3",
			DeviceParams: []cntl.DMXDeviceParams{
				{
					Device: &cntl.DMXDeviceSelector{
						ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
					},
					Params: &cntl.DMXParams{
						Green: Value255,
					},
				},
			}},

		"5d3a415a-0b15-11e7-90b9-03c2b960e034": {
			ID:   "5d3a415a-0b15-11e7-90b9-03c2b960e034",
			Name: "Test-Preset 4",
			DeviceParams: []cntl.DMXDeviceParams{
				{Group: &cntl.DMXDeviceGroupSelector{
					ID: "cb58bc10-0b16-11e7-b45a-7bee591b0adb",
				},
					Params: &cntl.DMXParams{
						Strobe: Value255,
					},
				},
			},
		},
		"4e3c2e84-0b15-11e7-a076-4b5bbb4c19bf": {
			ID:   "4e3c2e84-0b15-11e7-a076-4b5bbb4c19bf",
			Name: "Test-Preset 5",
			DeviceParams: []cntl.DMXDeviceParams{
				{
					Group: &cntl.DMXDeviceGroupSelector{ID: "475b71a0-0b16-11e7-9406-e3f678e8b788"},
					Params: &cntl.DMXParams{
						Red: Value200,
					},
				},
			},
		},
	},
	DMXScenes: map[string]*cntl.DMXScene{
		"492cef2e-0b14-11e7-be89-c3fa25f9cabb": {
			ID:        "492cef2e-0b14-11e7-be89-c3fa25f9cabb",
			Name:      "Test-Scene 1",
			NoteCount: 4,
			NoteValue: 4,
			SubScenes: []cntl.DMXSubScene{
				{
					At:     []uint64{0, 1, 2, 3},
					Preset: "0de258e0-0e7b-11e7-afd4-ebf6036983dc",
				},
			},
		},
		"a44f8dee-0b14-11e7-b5b9-bf1015384192": {
			ID:        "a44f8dee-0b14-11e7-b5b9-bf1015384192",
			Name:      "Test-Scene 2",
			NoteCount: 2,
			NoteValue: 4,
			SubScenes: []cntl.DMXSubScene{
				{
					At:     []uint64{0, 1},
					Preset: "11adf93e-0e7b-11e7-998c-5bd2bd0df396",
				},
			},
		},
		"99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6": {
			ID:        "99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6",
			Name:      "Test-Scene 3",
			NoteCount: 8,
			NoteValue: 4,
			SubScenes: []cntl.DMXSubScene{
				{
					At: []uint64{0, 1, 2, 3, 4, 5, 6, 7},
					DeviceParams: []cntl.DMXDeviceParams{
						{
							Device: &cntl.DMXDeviceSelector{
								ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
							},
							AnimationID: "a51f7b2a-0e7b-11e7-bfc8-57da167865d7",
						},
					},
				},
			},
		},
		"b82f4750-0e7a-11e7-9522-0f9d6d69958a": {
			ID:        "b82f4750-0e7a-11e7-9522-0f9d6d69958a",
			Name:      "Test-Scene 4",
			NoteCount: 4,
			NoteValue: 4,
			SubScenes: []cntl.DMXSubScene{
				{
					At:     []uint64{0, 1, 2, 3},
					Preset: "0de258e0-0e7b-11e7-afd4-ebf6036983dc",
				},
			},
		},
		"5adec126-6618-42ca-b930-4d18f4524328": {
			ID:        "5adec126-6618-42ca-b930-4d18f4524328",
			Name:      "Test-Scene 5",
			NoteCount: 32,
			NoteValue: 32,
			SubScenes: []cntl.DMXSubScene{
				{
					At: []uint64{0},
					DeviceParams: []cntl.DMXDeviceParams{
						{
							Device: &cntl.DMXDeviceSelector{
								ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
							},
							TransitionID: "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21",
						},
					},
				},
			},
		},
	},
	DMXDeviceGroups: map[string]*cntl.DMXDeviceGroup{
		"475b71a0-0b16-11e7-9406-e3f678e8b788": {
			ID:   "475b71a0-0b16-11e7-9406-e3f678e8b788",
			Name: "All PARs on the left side",
			Devices: []cntl.DMXDeviceSelector{
				{
					Tags: []cntl.Tag{"par", "left"},
				},
			},
		},
		"29f7adf8-0b17-11e7-bd45-9f82a70b477b": {
			ID:   "29f7adf8-0b17-11e7-bd45-9f82a70b477b",
			Name: "All PARs on the right side",
			Devices: []cntl.DMXDeviceSelector{
				{
					Tags: []cntl.Tag{"par", "right"},
				},
			},
		},
		"cb58bc10-0b16-11e7-b45a-7bee591b0adb": {
			ID:   "cb58bc10-0b16-11e7-b45a-7bee591b0adb",
			Name: "LED Bar infront the drums",
			Devices: []cntl.DMXDeviceSelector{
				{
					ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
				},
			},
		},
	},
	DMXDeviceTypes: map[string]*cntl.DMXDeviceType{
		"1555d67e-1187-11e7-8135-9b41038b5b75": {
			ID:             "1555d67e-1187-11e7-8135-9b41038b5b75",
			Name:           "LED-Bar 67 Channel",
			Key:            "LEDBar67",
			ChannelCount:   67,
			ChannelsPerLED: 4,
			ModeEnabled:    true,
			ModeChannel:    0,
			DimmerEnabled:  true,
			DimmerChannel:  1,
			StrobeEnabled:  true,
			StrobeChannel:  2,
			LEDs: []cntl.LED{
				{Position: 0, Red: 0, Green: 1, Blue: 2, White: 3},
				{Position: 1, Red: 4, Green: 5, Blue: 6, White: 7},
				{Position: 2, Red: 8, Green: 9, Blue: 10, White: 11},
				{Position: 3, Red: 12, Green: 13, Blue: 14, White: 15},
				{Position: 4, Red: 16, Green: 17, Blue: 18, White: 19},
				{Position: 5, Red: 20, Green: 21, Blue: 22, White: 23},
				{Position: 6, Red: 24, Green: 25, Blue: 26, White: 27},
				{Position: 7, Red: 28, Green: 29, Blue: 30, White: 31},
				{Position: 8, Red: 32, Green: 33, Blue: 34, White: 35},
				{Position: 9, Red: 36, Green: 37, Blue: 38, White: 39},
				{Position: 10, Red: 40, Green: 41, Blue: 42, White: 43},
				{Position: 11, Red: 44, Green: 45, Blue: 46, White: 47},
				{Position: 12, Red: 48, Green: 49, Blue: 50, White: 51},
				{Position: 13, Red: 52, Green: 53, Blue: 54, White: 55},
				{Position: 14, Red: 56, Green: 57, Blue: 58, White: 59},
				{Position: 15, Red: 60, Green: 61, Blue: 62, White: 63},
			},
		},
		"628fc3ea-1188-11e7-8824-5f72d80c17b6": {
			ID:             "628fc3ea-1188-11e7-8824-5f72d80c17b6",
			Name:           "PAR 5 channel",
			Key:            "PAR5",
			ChannelCount:   5,
			ChannelsPerLED: 3,
			ModeEnabled:    false,
			ModeChannel:    0,
			DimmerEnabled:  true,
			DimmerChannel:  3,
			StrobeEnabled:  true,
			StrobeChannel:  4,
			LEDs: []cntl.LED{
				{Position: 0, Red: 0, Green: 1, Blue: 2},
			},
		},
		"5ccc43ee-118c-11e7-8d53-974b41748b71": {
			ID:            "5ccc43ee-118c-11e7-8d53-974b41748b71",
			Name:          "Strobe",
			Key:           "Strobe",
			StrobeEnabled: true,
			StrobeChannel: 0,
		},
	},
	DMXDevices: map[string]*cntl.DMXDevice{
		"35cae00a-0b17-11e7-8bca-bbf30c56f20e": {
			ID:           "35cae00a-0b17-11e7-8bca-bbf30c56f20e",
			Name:         "LED-Bar below drums front",
			TypeID:       "1555d67e-1187-11e7-8135-9b41038b5b75",
			Universe:     1,
			StartChannel: 222,
			Tags:         []cntl.Tag{"bar", "drums"},
		},
		"s429fc37c-0b17-11e7-8b94-c3b6519355d3": {
			ID:           "s429fc37c-0b17-11e7-8b94-c3b6519355d3",
			Name:         "PAR inner left, stand 1 position 1",
			TypeID:       "628fc3ea-1188-11e7-8824-5f72d80c17b6",
			Universe:     2,
			StartChannel: 10,
			Tags:         []cntl.Tag{"par", "left", "inner", "stand", "drums-left"},
		},
		"4a545466-0b17-11e7-9c61-d3c0693099ab": {
			ID:           "4a545466-0b17-11e7-9c61-d3c0693099ab",
			Name:         "PAR inner left, stand 1 position 2",
			TypeID:       "628fc3ea-1188-11e7-8824-5f72d80c17b6",
			Universe:     2,
			StartChannel: 14,
			Tags:         []cntl.Tag{"par", "left", "inner", "stand", "drums-left"},
		},
		"5e0335e0-0b17-11e7-ad6c-63a7138d926c": {
			ID:           "5e0335e0-0b17-11e7-ad6c-63a7138d926c",
			Name:         "PAR inner right, stand 2 position 1",
			TypeID:       "628fc3ea-1188-11e7-8824-5f72d80c17b6",
			Universe:     2,
			StartChannel: 26,
			Tags:         []cntl.Tag{"par", "right", "inner", "stand", "drums-right"},
		},
		"620101f4-0b17-11e7-85cc-539952d9aef2": {
			ID:           "620101f4-0b17-11e7-85cc-539952d9aef2",
			Name:         "PAR inner right, stand 2 position 2",
			TypeID:       "628fc3ea-1188-11e7-8824-5f72d80c17b6",
			Universe:     2,
			StartChannel: 30,
			Tags:         []cntl.Tag{"par", "right", "inner", "stand", "drums-right"},
		},
		"6f7bca8a-0b17-11e7-b604-a356da737e54": {
			ID:           "6f7bca8a-0b17-11e7-b604-a356da737e54",
			Name:         "Strobe Vocs",
			TypeID:       "5ccc43ee-118c-11e7-8d53-974b41748b71",
			Universe:     1,
			StartChannel: 202,
			Tags:         []cntl.Tag{"strobe-back", "vocs"},
		},
	},
	DMXAnimations: map[string]*cntl.DMXAnimation{
		"a51f7b2a-0e7b-11e7-bfc8-57da167865d7": {
			ID:     "a51f7b2a-0e7b-11e7-bfc8-57da167865d7",
			Length: 4,
			Frames: []cntl.DMXAnimationFrame{
				{At: 0, Params: cntl.DMXParams{LED: 1, Blue: Value31}},
				{At: 1, Params: cntl.DMXParams{LED: 1, Blue: Value63}},
				{At: 2, Params: cntl.DMXParams{LED: 1, Blue: Value127}},
				{At: 3, Params: cntl.DMXParams{LED: 1, Blue: Value255}},
			},
		},
	},
	DMXTransitions: map[string]*cntl.DMXTransition{
		"a1a02b6c-12dd-4d7b-bc3e-24cc823adf21": {
			ID:   "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21",
			Name: "Blue Bar pulsing on",
			Ease: cntl.EaseInOutQuad,
			Length: 8,
			Params: []cntl.DMXTransitionParams{
				{
					From: cntl.DMXParams{
						LED:  1,
						Blue: Value0,
					},
					To: cntl.DMXParams{
						LED:  1,
						Blue: Value255,
					},
				},
				{
					From: cntl.DMXParams{
						LED:  2,
						Blue: Value0,
					},
					To: cntl.DMXParams{
						LED:  2,
						Blue: Value255,
					},
				},
			},
		},
	},
}

// DataStore returns the go object representation of a working set of fixtures
func DataStore() *cntl.DataStore {
	return data
}
