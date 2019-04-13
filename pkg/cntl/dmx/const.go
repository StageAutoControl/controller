package dmx

import "github.com/StageAutoControl/controller/pkg/cntl"

// DeviceTypes usable in this application
const (
	ChannelRed cntl.DMXChannel = 1 + iota
	ChannelGreen
	ChannelBlue
	ChannelWhite
	ChannelStrobe
	ChannelMode
	ChannelDimmer
	ChannelTilt
	ChannelTiltFine
	ChannelPan
	ChannelPanFine
	ChannelPanTiltSpeed
)
