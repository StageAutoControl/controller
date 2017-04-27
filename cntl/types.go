package cntl

import "encoding/json"

// SongSelector is a ID selector for a song
type SongSelector struct {
	ID string `json:"id" yaml:"id"`
}

// SetList is a set of songs in a specific order
type SetList struct {
	ID    string         `json:"id" yaml:"id"`
	Name  string         `json:"name" yaml:"name"`
	Songs []SongSelector `json:"songs" yaml:"songs"`
}

// BarChange describes the changes of tempo and notes during a song
type BarChange struct {
	At        uint64 `json:"at" yaml:"at"`
	NoteValue uint8  `json:"noteValue" yaml:"noteValue"`
	NoteCount uint8  `json:"noteCount" yaml:"noteCount"`
	Speed     uint16 `json:"speed" yaml:"speed"`
}

// ScenePosition describes the position of a DMX scene within a song
type ScenePosition struct {
	ID     string `json:"id" yaml:"id"`
	At     uint64 `json:"at" yaml:"at"`
	Repeat uint8  `json:"repeat" yaml:"repeat"`
}

// Song is the whole container for everything that needs to be controlled during a song.
type Song struct {
	ID              string            `json:"id" yaml:"id"`
	Name            string            `json:"name" yaml:"name"`
	Length          uint64            `json:"length" yaml:"length"`
	BarChanges      []BarChange       `json:"barChanges" yaml:"barChanges"`
	DMXScenes       []ScenePosition   `json:"dmxScenes" yaml:"dmxScenes"`
	DMXDeviceParams []DMXDeviceParams `json:"dmxDeviceParams" yaml:"dmxDeviceParams"`
}

// Tag is a string literal tagging a DMX device
type Tag string

// DMXDevice is a DMX device
type DMXDevice struct {
	ID           string      `json:"id" yaml:"id"`
	Name         string      `json:"name" yaml:"name"`
	TypeID       string      `json:"typeId" yaml:"typeId"`
	Universe     DMXUniverse `json:"universe" yaml:"universe"`
	StartChannel DMXChannel  `json:"startChannel" yaml:"startChannel"`
	Tags         []Tag       `json:"tags" yaml:"tags"`
}

// DMXDeviceType is the type of a DMXDevice
type DMXDeviceType struct {
	ID             string     `json:"id" yaml:"id"`
	Name           string     `json:"name" yaml:"name"`
	Key            string     `json:"key" yaml:"key"`
	ChannelCount   uint16     `json:"addressCount" yaml:"addressCount"`
	ChannelsPerLED uint16     `json:"channelsPerLED" yaml:"channelsPerLED"`
	StrobeEnabled  bool       `json:"strobeEnabled" yaml:"strobeEnabled"`
	StrobeChannel  DMXChannel `json:"strobeChannel" yaml:"strobeChannel"`
	DimmerEnabled  bool       `json:"dimmerEnabled" yaml:"dimmerEnabled"`
	DimmerChannel  DMXChannel `json:"dimmerChannel" yaml:"dimmerChannel"`
	ModeEnabled    bool       `json:"presetEnabled" yaml:"presetEnabled"`
	ModeChannel    DMXChannel `json:"presetChannel" yaml:"presetChannel"`
	LEDs           []LED      `json:"leds"`
}

// LED maps a single LEDs DMX channels
type LED struct {
	Position uint16     `json:"position" yaml:"position"`
	Red      DMXChannel `json:"red" yaml:"red"`
	Green    DMXChannel `json:"green" yaml:"green"`
	Blue     DMXChannel `json:"blue" yaml:"blue"`
	White    DMXChannel `json:"white" yaml:"white"`
}

// DMXDeviceSelector is a selector for DMX devices
type DMXDeviceSelector struct {
	ID   string `json:"id" yaml:"id"`
	Tags []Tag  `json:"tags" yaml:"tags"`
}

// DMXDeviceGroupSelector is a selector for DMX device groups
type DMXDeviceGroupSelector struct {
	ID string `json:"id" yaml:"id"`
}

// DMXDeviceGroup is a DMX device group
type DMXDeviceGroup struct {
	ID      string              `json:"id" yaml:"id"`
	Name    string              `json:"name" yaml:"name"`
	Devices []DMXDeviceSelector `json:"devices" yaml:"devices"`
}

// DMXDeviceParams is an object storing DMX parameters including the selection of either groups or devices
type DMXDeviceParams struct {
	Group       *DMXDeviceGroupSelector `json:"group" yaml:"group"`
	Device      *DMXDeviceSelector      `json:"device" yaml:"device"`
	Params      *DMXParams              `json:"params" yaml:"params"`
	AnimationID string                  `json:"animationId" yaml:"animationId"`
}

// DMXScene is a whole light scene
type DMXScene struct {
	ID        string        `json:"id" yaml:"id"`
	Name      string        `json:"name" yaml:"name"`
	NoteValue uint8         `json:"noteValue" yaml:"noteValue"`
	NoteCount uint8         `json:"noteCount" yaml:"noteCount"`
	SubScenes []DMXSubScene `json:"subScenes" yaml:"subScenes"`
}

// DMXSubScene is a sub scene of a light scene
type DMXSubScene struct {
	At           []uint8           `json:"at" yaml:"at"`
	DeviceParams []DMXDeviceParams `json:"deviceParams" yaml:"deviceParams"`
	Preset       string            `json:"preset" yaml:"preset"`
}

// DMXParams is a DMX parameter object
type DMXParams struct {
	LED    uint16    `json:"led" yaml:"led"`
	Red    *DMXValue `json:"red" yaml:"red"`
	Green  *DMXValue `json:"green" yaml:"green"`
	Blue   *DMXValue `json:"blue" yaml:"blue"`
	Strobe *DMXValue `json:"strobe" yaml:"strobe"`
	Preset *DMXValue `json:"preset" yaml:"preset"`
}

// DMXAnimation is an animation of dmx params in relation to time
type DMXAnimation struct {
	ID     string              `json:"id" yaml:"id"`
	Length uint8               `json:"length" yaml:"length"`
	Frames []DMXAnimationFrame `json:"frames" yaml:"frames"`
}

// DMXAnimationFrame is a single frame in an animation
type DMXAnimationFrame struct {
	At     uint8     `json:"at" yaml:"at"`
	Params DMXParams `json:"params" yaml:"params"`
}

// DMXPreset is a DMX Preet for devices or device groups
type DMXPreset struct {
	ID           string            `json:"id" yaml:"id"`
	Name         string            `json:"name" yaml:"name"`
	DeviceParams []DMXDeviceParams `json:"deviceParams" yaml:"deviceParams"`
}

// Command is a container to set settings
type Command struct {
	DMXCommands  DMXCommands  `json:"dmxCommands" yaml:"dmxCommands"`
	MIDICommands MIDICommands `json:"midiCommands" yaml:"midiCommands"`
	BarChange    *BarChange   `json:"barChange" yaml:"barChange"`
}

// DMXCommand tells a DMX controller to set a channel on a universe to a specific value
type DMXCommand struct {
	Universe DMXUniverse `json:"universe" yaml:"universe"`
	Channel  DMXChannel  `json:"channel" yaml:"channel"`
	Value    DMXValue    `json:"value" yaml:"value"`
}

// Equals returns true if given DMXCommand is equal to the called one
func (cmd DMXCommand) Equals(c DMXCommand) bool {
	return cmd.EqualsChannel(c) &&
		cmd.Value == c.Value
}

// EqualsChannel returns true if given DMXCommand is equal to the called one in terms of channel and universe
func (cmd DMXCommand) EqualsChannel(c DMXCommand) bool {
	return cmd.Channel == c.Channel &&
		cmd.Universe == c.Universe
}

// DMXCommands is an array of DMXCommands
type DMXCommands []DMXCommand

// Contains returns whether given DMXCommand is in the called collection
func (cmds DMXCommands) Contains(c DMXCommand) bool {
	for _, cmd := range cmds {
		if cmd.Equals(c) {
			return true
		}
	}

	return false
}

// ContainsChannel returns whether given DMXCommand's channel and universe is in the called collection
func (cmds DMXCommands) ContainsChannel(c DMXCommand) bool {
	for _, cmd := range cmds {
		if cmd.EqualsChannel(c) {
			return true
		}
	}

	return false
}

// Equals returns true when both the called and the given one have the same entries without caring for order
func (cmds DMXCommands) Equals(c DMXCommands) bool {
	if len(cmds) != len(c) {
		return false
	}

	for _, cmd := range cmds {
		if !c.Contains(cmd) {
			return false
		}
	}

	return true
}

// DMXUniverse is the universe a DMXDevice is in
type DMXUniverse uint16

// DMXChannel is the channel a command can talk to (0-511)
type DMXChannel uint16

// DMXValue is the value a DMX channel can represent (0-255)
type DMXValue struct {
	Value uint8
}

// MarshalYAML encodes the value to YAML
func (v *DMXValue) MarshalYAML() (interface{}, error) {
	return v.Value, nil
}

// UnmarshalYAML takes the value from YAML
func (v *DMXValue) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal(&v.Value); err != nil {
		return err
	}

	return nil
}

// MarshalJSON converts the value to a json byte array
func (v *DMXValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

// UnmarshalJSON sets the value from a json byte array
func (v *DMXValue) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &v.Value)
}

// Equals returns whether the two given objects are equal
func (v *DMXValue) Equals(v2 *DMXValue) bool {
	return v != nil && v2 != nil && v.Value == v2.Value
}

// MIDICommand tells a MIDI controller to set a channel to a specific value
type MIDICommand struct {
	Channel uint8 `json:"channel" yaml:"channel"`
	Value   uint8 `json:"value" yaml:"value"`
}

// MIDICommands is an array of MIDICommands
type MIDICommands []MIDICommand
