package cntl

// SetList is a set of songs in a specific order
type SetList struct {
	ID    string   `json:"id" yaml:"id"`
	Name  string   `json:"name" yaml:"name"`
	Songs []string `json:"songs" yaml:"songs"`
}

// BarParams are a reusable informational struct on how fast and in what scheme something should be played
type BarParams struct {
	NoteValue uint8  `json:"noteValue" yaml:"noteValue"`
	NoteCount uint8  `json:"noteCount" yaml:"noteCount"`
	Speed     uint16 `json:"speed" yaml:"speed"`
}

// BarChange describes the changes of tempo and notes during a song
type BarChange struct {
	BarParams
	At uint64 `json:"at" yaml:"at"`
}

// DMXScenePosition describes the position of a DMX scene within a song
type DMXScenePosition struct {
	ID     string `json:"id" yaml:"id"`
	At     uint64 `json:"at" yaml:"at"`
	Repeat uint8  `json:"repeat" yaml:"repeat"`
}

// Song is the whole container for everything that needs to be controlled during a song.
type Song struct {
	ID           string             `json:"id" yaml:"id"`
	Name         string             `json:"name" yaml:"name"`
	BarChanges   []BarChange        `json:"barChanges" yaml:"barChanges"`
	DMXScenes    []DMXScenePosition `json:"dmxScenes" yaml:"dmxScenes"`
	MIDICommands []MIDICommand      `json:"midiCommands" yaml:"midiCommands"`
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
	ChannelCount   uint16     `json:"channelCount" yaml:"channelCount"`
	ChannelsPerLED uint16     `json:"channelsPerLED" yaml:"channelsPerLED"`
	StrobeEnabled  bool       `json:"strobeEnabled" yaml:"strobeEnabled"`
	StrobeChannel  DMXChannel `json:"strobeChannel" yaml:"strobeChannel"`
	DimmerEnabled  bool       `json:"dimmerEnabled" yaml:"dimmerEnabled"`
	DimmerChannel  DMXChannel `json:"dimmerChannel" yaml:"dimmerChannel"`
	ModeEnabled    bool       `json:"modeEnabled" yaml:"modeEnabled"`
	ModeChannel    DMXChannel `json:"modeChannel" yaml:"modeChannel"`
	Moving         bool       `json:"moving" yaml:"moving"`
	TiltChannel    DMXChannel `json:"tiltChannel" yaml:"tiltChannel"`
	PanChannel     DMXChannel `json:"panChannel" yaml:"panChannel"`
	LEDs           []LED      `json:"leds"`
}

// LED maps a single LEDs DMX channels
type LED struct {
	Red   DMXChannel `json:"red" yaml:"red"`
	Green DMXChannel `json:"green" yaml:"green"`
	Blue  DMXChannel `json:"blue" yaml:"blue"`
	White DMXChannel `json:"white" yaml:"white"`
}

// DMXDeviceSelector is a selector for DMX devices
type DMXDeviceSelector struct {
	ID   string `json:"id" yaml:"id"`
	Tags []Tag  `json:"tags" yaml:"tags"`
}

// DMXDeviceGroup is a DMX device group
type DMXDeviceGroup struct {
	ID      string              `json:"id" yaml:"id"`
	Name    string              `json:"name" yaml:"name"`
	Devices []DMXDeviceSelector `json:"devices" yaml:"devices"`
}

// DMXDeviceParams is an object storing DMX parameters including the selection of either groups or devices
type DMXDeviceParams struct {
	Group      *string     `json:"group" yaml:"group"`
	Device     *string     `json:"device" yaml:"device"`
	Params     []DMXParams `json:"params" yaml:"params"`
	Animation  *string     `json:"animation" yaml:"animation"`
	Transition *string     `json:"transition" yaml:"transition"`
}

// DMXScene is a whole light scene
type DMXScene struct {
	ID        string        `json:"id" yaml:"id"`
	Name      string        `json:"name" yaml:"name"`
	NoteValue uint8         `json:"noteValue" yaml:"noteValue"`
	NoteCount uint16        `json:"noteCount" yaml:"noteCount"`
	SubScenes []DMXSubScene `json:"subScenes" yaml:"subScenes"`
}

// DMXSubScene is a sub scene of a light scene
type DMXSubScene struct {
	At           []uint64          `json:"at" yaml:"at"`
	DeviceParams []DMXDeviceParams `json:"deviceParams" yaml:"deviceParams"`
	Preset       *string           `json:"preset" yaml:"preset"`
}

// DMXColorVariable is a global variable for a DMX color
type DMXColorVariable struct {
	ID    string    `json:"id" yaml:"id"`
	Name  string    `json:"name" yaml:"name"`
	Red   *DMXValue `json:"red" yaml:"red"`
	Green *DMXValue `json:"green" yaml:"green"`
	Blue  *DMXValue `json:"blue" yaml:"blue"`
	White *DMXValue `json:"white" yaml:"white"`
}

// DMXParams is a DMX parameter object
type DMXParams struct {
	LED      uint16    `json:"led" yaml:"led"`
	ColorVar *string   `json:"$color" yaml:"$color"`
	Red      *DMXValue `json:"red" yaml:"red"`
	Green    *DMXValue `json:"green" yaml:"green"`
	Blue     *DMXValue `json:"blue" yaml:"blue"`
	White    *DMXValue `json:"white" yaml:"white"`
	Pan      *DMXValue `json:"pan" yaml:"pan"`
	Tilt     *DMXValue `json:"tilt" yaml:"tilt"`
	Strobe   *DMXValue `json:"strobe" yaml:"strobe"`
	Mode     *DMXValue `json:"mode" yaml:"mode"`
	Dimmer   *DMXValue `json:"dimmer" yaml:"dimmer"`
}

// DMXAnimation is an animation of dmx params in relation to time
type DMXAnimation struct {
	ID     string              `json:"id" yaml:"id"`
	Name   string              `json:"name" yaml:"name"`
	Frames []DMXAnimationFrame `json:"frames" yaml:"frames"`
}

// DMXTransition is a transition from a given state to another one using an ease function
type DMXTransition struct {
	ID     string                `json:"id" yaml:"id"`
	Name   string                `json:"name" yaml:"name"`
	Ease   EaseFunc              `json:"ease" yaml:"ease"`
	Length uint8                 `json:"length" yaml:"length"`
	Params []DMXTransitionParams `json:"params" yaml:"params"`
}

// DMXTransitionParams hold the params for a transition
type DMXTransitionParams struct {
	From DMXParams `json:"from" yaml:"from"`
	To   DMXParams `json:"to" yaml:"to"`
}

// EaseFunc names a function that is used to ease a transition
type EaseFunc string

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
	FrameState
	DMXCommands  DMXCommands  `json:"dmxCommands" yaml:"dmxCommands"`
	MIDICommands MIDICommands `json:"midiCommands" yaml:"midiCommands"`
	BarChange    *BarChange   `json:"barChange" yaml:"barChange"`
}

// FrameState stores information about which bar and note the command is in
type FrameState struct {
	Frame uint64 `json:"frame" yaml:"frame"`
	Bar   uint16 `json:"bar" yaml:"bar"`
	Note  uint8  `json:"note" yaml:"note"`
}

// DMXCommand tells a DMX controller to set a channel on a universe to a specific value
type DMXCommand struct {
	Universe DMXUniverse `json:"universe" yaml:"universe"`
	Channel  DMXChannel  `json:"channel" yaml:"channel"`
	Value    DMXValue    `json:"value" yaml:"value"`
}

// DMXCommands is an array of DMXCommands
type DMXCommands []DMXCommand

// DMXUniverse is the universe a DMXDevice is in
type DMXUniverse uint16

// DMXChannel is the channel a command can talk to (0-511)
type DMXChannel uint16

// DMXValue is the value a DMX channel can represent (0-255)
type DMXValue struct {
	Value uint8
}

// MIDICommand tells a MIDI controller to set a channel to a specific value
type MIDICommand struct {
	At     uint64 `json:"at" yaml:"at"`
	Status uint8  `json:"status" yaml:"status"`
	Data1  uint8  `json:"data1" yaml:"data1"`
	Data2  uint8  `json:"data2" yaml:"data2"`
}

// MIDICommands is an array of MIDICommands
type MIDICommands []MIDICommand
