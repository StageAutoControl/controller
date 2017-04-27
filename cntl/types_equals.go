package cntl

// Equals returns whether the two given objects are equal
func (v1 SongSelector) Equals(v2 SongSelector) bool {
	return v1.ID == v2.ID
}

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
func (v1 ScenePosition) Equals(v2 ScenePosition) bool {
	return v1.At == v2.At &&
		v1.ID == v2.ID &&
		v1.Repeat == v2.Repeat
}

// Equals returns whether the two given objects are equal
func (v1 *Song) Equals(v2 *Song) bool {
	return v1.ID == v2.ID &&
		v1.Name == v2.Name &&
		v1.Length == v2.Length &&
		barChangeList(v1.BarChanges).Equals(barChangeList(v2.BarChanges)) &&
		scenePositionList(v1.DMXScenes).Equals(scenePositionList(v2.DMXScenes)) &&
		dmxDeviceParamsList(v1.DMXDeviceParams).Equals(dmxDeviceParamsList(v2.DMXDeviceParams))
}

// Equals returns whether the two given objects are equal
func (v1 Tag) Equals(v2 Tag) bool {
	return v1 == v2
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
		v1.Key == v2.Key &&
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
	return v1.Position == v2.Position &&
		v1.Red == v2.Red &&
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
func (v1 DMXDeviceGroupSelector) Equals(v2 DMXDeviceGroupSelector) bool {
	return v1.ID == v2.ID
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

	return (v1.Group == nil && v2.Group == nil || (*v1.Group).Equals(*v2.Group)) &&
		(v1.Device == nil && v2.Device == nil || (*v1.Device).Equals(*v2.Device)) &&
		(v1.Params == nil && v2.Params == nil || (*v1.Params).Equals(*v2.Params)) &&
		v1.AnimationID == v2.AnimationID
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
		(v1.Preset == nil && v2.Preset == nil || v1.Preset != nil && v2.Preset != nil && v1.Preset.Equals(v2.Preset)) &&
		(v1.Strobe == nil && v2.Strobe == nil || v1.Strobe != nil && v2.Strobe != nil && v1.Strobe.Equals(v2.Strobe)) &&
		(v1.Red == nil && v2.Red == nil || v1.Red != nil && v2.Red != nil && v1.Red.Equals(v2.Red)) &&
		(v1.Green == nil && v2.Green == nil || v1.Green != nil && v2.Green != nil && v1.Green.Equals(v2.Green)) &&
		(v1.Blue == nil && v2.Blue == nil || v1.Blue != nil && v2.Blue != nil && v1.Blue.Equals(v2.Blue))
}

// Equals returns whether the two given objects are equal
func (v1 DMXAnimation) Equals(v2 DMXAnimation) bool {
	return v2.ID == v2.ID &&
		v1.Length == v2.Length &&
		dmxAnimationFrameList(v1.Frames).Equals(dmxAnimationFrameList(v2.Frames))
}

// Equals returns whether the two given objects are equal
func (v1 DMXAnimationFrame) Equals(v2 DMXAnimationFrame) bool {
	return v1.At == v2.At &&
		v1.Params.Equals(v2.Params)
}

// Equals returns whether the two given objects are equal
func (v1 DMXPreset) Equals(v2 DMXPreset) bool {
	return v2.ID == v2.ID &&
		v1.Name == v2.Name &&
		dmxDeviceParamsList(v1.DeviceParams).Equals(dmxDeviceParamsList(v2.DeviceParams))
}
