package dmx

// Tag is a string literal tagging a DMX device
type Tag string

// Device is a DMX device
type Device struct {
	ID             string `json:"id" yaml:"id"`
	Name           string `json:"name" yaml:"name"`
	Type           string `json:"type" yaml:"type"`
	Universe       uint16 `json:"universe" yaml:"universe"`
	StartAddress   uint16 `json:"startAddress" yaml:"startAddress"`
	AddressRange   uint16 `json:"addressRange" yaml:"addressRange"`
	ChannelsPerLED uint8  `json:"channelsPerLED" yaml:"channelsPerLED"`
	Tags           []Tag  `json:"tags" yaml:"tags"`
}

// DeviceSelector is a selector for DMX devices
type DeviceSelector struct {
	ID   string `json:"id" yaml:"id"`
	Tags []Tag  `json:"tags" yaml:"tags"`
}

// DeviceGroupSelector is a selector for DMX device groups
type DeviceGroupSelector struct {
	ID string `json:"id" yaml:"id"`
}

// DeviceGroup is a DMX device group
type DeviceGroup struct {
	ID      string            `json:"id" yaml:"id"`
	Name    string            `json:"name" yaml:"name"`
	Devices []*DeviceSelector `json:"devices" yaml:"devices"`
}

// DeviceParams is an object storing DMX parameters including the selection of either groups or devices
type DeviceParams struct {
	Group  *DeviceGroupSelector `json:"group" yaml:"group"`
	Device *DeviceSelector      `json:"device" yaml:"device"`
	Params *Params              `json:"params" yaml:"params"`
}

// Scene is a whole light scene
type Scene struct {
	ID        string      `json:"id" yaml:"id"`
	Name      string      `json:"name" yaml:"name"`
	NoteValue uint8       `json:"noteValue" yaml:"noteValue"`
	NoteCount uint8       `json:"noteCount" yaml:"noteCount"`
	SubScenes []*SubScene `json:"subScenes" yaml:"subScenes"`
}

// SubScene is a sub scene of a light scene
type SubScene struct {
	At           []uint8         `json:"at" yaml:"at"`
	DeviceParams []*DeviceParams `json:"deviceParams" yaml:"deviceParams"`
	Animation    string          `json:"animation" yaml:"animations"`
	Preset       string          `json:"preset" yaml:"preset"`
}

// Params is a DMX parameter object
type Params struct {
	LED    uint64 `json:"led" yaml:"led"`
	Red    uint8  `json:"red" yaml:"red"`
	Green  uint8  `json:"green" yaml:"green"`
	Blue   uint8  `json:"blue" yaml:"blue"`
	Strobe uint8  `json:"strobe" yaml:"strobe"`
}

// Animation is an animation of dmx params in relation to time
type Animation struct {
	ID     string            `json:"id" yaml:"id"`
	Length uint8             `json:"length" yaml:"length"`
	Frames []*AnimationFrame `json:"frames" yaml:"frames"`
}

// AnimationFrame is a single frame in an animation
type AnimationFrame struct {
	At     uint8  `json:"at" yaml:"at"`
	Params Params `json:"params" yaml:"params"`
}

// Preset is a DMX Preet for devices or device groups
type Preset struct {
	ID           string          `json:"id" yaml:"id"`
	Name         string          `json:"name" yaml:"name"`
	DeviceParams []*DeviceParams `json:"deviceParams" yaml:"deviceParams"`
}
