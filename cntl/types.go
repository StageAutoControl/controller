package cntl

import "github.com/StageAutoControl/controller/cntl/dmx"

// SongSelector is a ID selector for a song
type SongSelector struct {
	ID string `json:"id" yaml:"id"`
}

// SetList is a set of songs in a specific order
type SetList struct {
	ID    string          `json:"id" yaml:"id"`
	Name  string          `json:"name" yaml:"name"`
	Songs []*SongSelector `json:"songs" yaml:"songs"`
}

// BarChange describes the changes of tempo and notes during a song
type BarChange struct {
	At              uint16              `json:"at" yaml:"at"`
	NoteValue       uint8               `json:"noteValue" yaml:"noteValue"`
	NoteCount       uint8               `json:"noteCount" yaml:"noteCount"`
	Speed           uint16              `json:"speed" yaml:"speed"`
	DmxScenes       []*dmx.Scene        `json:"dmxScenes" yaml:"dmxScenes"`
	DmxDeviceParams []*dmx.DeviceParams `json:"dmxDeviceParams" yaml:"dmxDeviceParams"`
}

// ScenePosition describes the position of a DMX scene within a song
type ScenePosition struct {
	ID     string `json:"id" yaml:"id"`
	Start  uint16 `json:"start" yaml:"start"`
	Length uint8  `json:"length" yaml:"length"`
}

// Song is the whole container for everything that needs to be controlled during a song.
type Song struct {
	ID              string              `json:"id" yaml:"id"`
	Name            string              `json:"name" yaml:"name"`
	BarChanges      []*BarChange        `json:"barChanges" yaml:"barChanges"`
	DmxScenes       []*ScenePosition    `json:"dmxScenes" yaml:"dmxScenes"`
	DmxDeviceParams []*dmx.DeviceParams `json:"dmxDeviceParams" yaml:"dmxDeviceParams"`
}
