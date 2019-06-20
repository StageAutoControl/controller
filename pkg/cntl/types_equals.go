package cntl

// Equals returns whether the two given objects are equal
func (v1 *SetList) Equals(v2 *SetList) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		songSelectorList(v1.Songs).Equals(songSelectorList(v2.Songs))
}

// Equals returns whether the two given objects are equal
func (v1 BarChange) Equals(v2 BarChange) bool {
	return v1.At == v2.At &&
		v1.NoteCount == v2.NoteCount &&
		v1.NoteValue == v2.NoteValue &&
		v1.Speed == v2.Speed
}

// Equals returns whether the two given objects are equal
func (v1 DMXScenePosition) Equals(v2 DMXScenePosition) bool {
	return v1.At == v2.At &&
		v1.ID == v2.ID &&
		v1.Repeat == v2.Repeat
}

// Equals returns whether the two given objects are equal
func (v1 *Song) Equals(v2 *Song) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		barChangeList(v1.BarChanges).Equals(barChangeList(v2.BarChanges)) &&
		scenePositionList(v1.DMXScenes).Equals(scenePositionList(v2.DMXScenes))
}

// Equals returns whether the two given objects are equal
func (t Tag) Equals(v2 Tag) bool {
	return t == v2
}

// Equals returns whether the two given objects are equal
func (v1 *DMXDevice) Equals(v2 *DMXDevice) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		v1.TypeID == v2.TypeID &&
		v1.Universe == v2.Universe &&
		v1.StartChannel == v2.StartChannel &&
		tagList(v1.Tags).Equals(tagList(v2.Tags))
}

// Equals returns whether the two given objects are equal
func (v1 *DMXDeviceType) Equals(v2 *DMXDeviceType) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		v1.ChannelCount == v2.ChannelCount &&
		v1.ChannelsPerLED == v2.ChannelsPerLED &&
		v1.StrobeEnabled == v2.StrobeEnabled &&
		v1.StrobeChannel == v2.StrobeChannel &&
		v1.DimmerEnabled == v2.DimmerEnabled &&
		v1.DimmerChannel == v2.DimmerChannel &&
		v1.ModeEnabled == v2.ModeEnabled &&
		v1.ModeChannel == v2.ModeChannel &&
		ledList(v1.LEDs).Equals(ledList(v2.LEDs))

}

// Equals returns whether the two given objects are equal
func (v1 LED) Equals(v2 LED) bool {
	return v1.Red == v2.Red &&
		v1.Green == v2.Green &&
		v1.Blue == v2.Blue &&
		v1.White == v2.White
}

// Equals returns whether the two given objects are equal
func (v1 DMXDeviceSelector) Equals(v2 DMXDeviceSelector) bool {
	return v1.ID == v2.ID &&
		tagList(v1.Tags).Equals(tagList(v2.Tags))
}

// Equals returns whether the two given objects are equal
func (v1 *DMXDeviceGroup) Equals(v2 *DMXDeviceGroup) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		dmxDeviceSelectorList(v1.Devices).Equals(dmxDeviceSelectorList(v2.Devices))
}

// Equals returns whether the two given objects are equal
func (v1 DMXDeviceParams) Equals(v2 DMXDeviceParams) bool {
	if v1.Group == nil && v2.Group != nil || v1.Group != nil && v2.Group == nil {
		return false
	}

	if v1.Device == nil && v2.Device != nil || v1.Device != nil && v2.Device == nil {
		return false
	}

	if v1.Params == nil && v2.Params != nil || v1.Params != nil && v2.Params == nil {
		return false
	}

	return dmxParamsList(v1.Params).Equals(dmxParamsList(v2.Params))
}

// Equals returns whether the two given objects are equal
func (v1 DMXScene) Equals(v2 DMXScene) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		v1.NoteValue == v2.NoteValue &&
		v1.NoteCount == v2.NoteCount &&
		dmxSubSceneList(v1.SubScenes).Equals(dmxSubSceneList(v2.SubScenes))
}

// Equals returns whether the two given objects are equal
func (v1 DMXSubScene) Equals(v2 DMXSubScene) bool {
	return atList(v1.At).Equals(atList(v2.At)) &&
		v1.Preset == v2.Preset &&
		dmxDeviceParamsList(v1.DeviceParams).Equals(dmxDeviceParamsList(v2.DeviceParams))
}

// Equals returns whether the two given objects are equal
func (v1 DMXParams) Equals(v2 DMXParams) bool {
	return v1.LED == v2.LED &&
		(v1.Mode == nil && v2.Mode == nil || v1.Mode != nil && v2.Mode != nil && v1.Mode.Equals(v2.Mode)) &&
		(v1.Strobe == nil && v2.Strobe == nil || v1.Strobe != nil && v2.Strobe != nil && v1.Strobe.Equals(v2.Strobe)) &&
		(v1.White == nil && v2.White == nil || v1.White != nil && v2.White != nil && v1.White.Equals(v2.White)) &&
		(v1.Red == nil && v2.Red == nil || v1.Red != nil && v2.Red != nil && v1.Red.Equals(v2.Red)) &&
		(v1.Green == nil && v2.Green == nil || v1.Green != nil && v2.Green != nil && v1.Green.Equals(v2.Green)) &&
		(v1.Blue == nil && v2.Blue == nil || v1.Blue != nil && v2.Blue != nil && v1.Blue.Equals(v2.Blue))
}

// Equals returns whether the two given objects are equal
func (v1 DMXAnimation) Equals(v2 DMXAnimation) bool {
	return v1.ID == v2.ID &&
		dmxAnimationFrameList(v1.Frames).Equals(dmxAnimationFrameList(v2.Frames))
}

// Equals returns whether the two given objects are equal
func (v1 DMXAnimationFrame) Equals(v2 DMXAnimationFrame) bool {
	return v1.At == v2.At &&
		v1.Params.Equals(v2.Params)
}

// Equals returns whether the two given objects are equal
func (v1 DMXPreset) Equals(v2 DMXPreset) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		dmxDeviceParamsList(v1.DeviceParams).Equals(dmxDeviceParamsList(v2.DeviceParams))
}

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
