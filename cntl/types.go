package cntl

type Tag string

type Device struct {
	ID           string
	Name         string
	Type         string
	Universe     uint16
	StartAddress uint16
	AddressRange uint16
	Tags         []Tag
}

type DeviceSelector struct {
	ID   string
	Tags []Tag
}

type GroupSelector struct {
	ID string
}

type DeviceGroup struct {
	ID      string
	Name    string
	Devices []*DeviceSelector
}

type DeviceParams struct {
	Group  *GroupSelector
	Device *DeviceSelector
	Params *Params
}

type Scene struct {
	ID   string
	Name string
}

type SubScene struct {
	At           uint16
	DeviceParams []*DeviceParams
}

type Params struct {
	Red    uint8
	Green  uint8
	Blue   uint8
	Strobe uint8
}
