package dmx

import (
	"errors"

	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

// DeviceTypes usable in this application
const (
	ChannelRed cntl.DMXChannel = 1 + iota
	ChannelGreen
	ChannelBlue
	ChannelWhite
	ChannelStrobe
	ChannelMode
	ChannelDimmer
)

// Error messages thrown during device channel calculation
var (
	errDeviceHasDisabledModeChannel   = errors.New("Device has disabled Model channel")
	errDeviceHasDisabledStrobeChannel = errors.New("Device has disabled Strobel channel")
	errDeviceHasDisabledDimmerChannel = errors.New("Device has disabled Dimmerl channel")
)

func getDeviceChannel(ds *cntl.DataStore, d *cntl.DMXDevice, c cntl.DMXChannel, led uint16) (cntl.DMXChannel, error) {
	dt, ok := ds.DMXDeviceTypes[d.TypeID]
	if !ok {
		return cntl.DMXChannel(0), fmt.Errorf("Given DeviceType %q on device %q is unknown.", d.TypeID, d.ID)
	}
	// can a param affect multiple LEDs?
	// Should I switch the scheme of params to have an
	// slice of LEDs and apply all values to that?

	ledLen := len(dt.LEDs)
	if ledLen > 0 && int(led) >= ledLen {
		return cntl.DMXChannel(0), fmt.Errorf("Given device has insufficient biggest index of LEDs %d to handle the given LED index %d", ledLen-1, led)
	}

	switch c {
	case ChannelRed:
		return d.StartChannel + getLED(dt, led).Red, nil

	case ChannelGreen:
		return d.StartChannel + getLED(dt, led).Green, nil

	case ChannelBlue:
		return d.StartChannel + getLED(dt, led).Blue, nil

	case ChannelWhite:
		return d.StartChannel + getLED(dt, led).White, nil

	case ChannelStrobe:
		if !dt.StrobeEnabled {
			return cntl.DMXChannel(0), errDeviceHasDisabledStrobeChannel
		}
		return d.StartChannel + dt.StrobeChannel, nil

	case ChannelMode:
		if !dt.ModeEnabled {
			return cntl.DMXChannel(0), errDeviceHasDisabledModeChannel
		}
		return d.StartChannel + dt.ModeChannel, nil

	case ChannelDimmer:
		if !dt.DimmerEnabled {
			return cntl.DMXChannel(0), errDeviceHasDisabledDimmerChannel
		}
		return d.StartChannel + dt.DimmerChannel, nil

	default:
		return cntl.DMXChannel(0), fmt.Errorf("Channel %q is unknown", c)
	}
}

func getLED(dt *cntl.DMXDeviceType, led uint16) *cntl.LED {
	for _, l := range dt.LEDs {
		if l.Position == led {
			return &l
		}
	}

	return nil
}
