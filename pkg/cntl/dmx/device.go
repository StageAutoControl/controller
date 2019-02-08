package dmx

import (
	"fmt"
	"github.com/StageAutoControl/controller/pkg/cntl"
)

func getDeviceChannel(ds *cntl.DataStore, d *cntl.DMXDevice, c cntl.DMXChannel, led uint16) (cntl.DMXChannel, error) {
	dt, ok := ds.DMXDeviceTypes[d.TypeID]
	if !ok {
		return cntl.DMXChannel(0), fmt.Errorf("given DeviceType %q on device %q is unknown", d.TypeID, d.ID)
	}
	// can a param affect multiple LEDs?
	// Should I switch the scheme of params to have an
	// slice of LEDs and apply all values to that?

	ledLen := len(dt.LEDs)
	if ledLen > 0 && int(led) >= ledLen {
		return cntl.DMXChannel(0), fmt.Errorf("given device has insufficient biggest index of LEDs %d to handle the given LED index %d", ledLen-1, led)
	}

	var channel cntl.DMXChannel

	switch c {
	case ChannelRed:
		channel = getLED(dt, led).Red

	case ChannelGreen:
		channel = getLED(dt, led).Green

	case ChannelBlue:
		channel = getLED(dt, led).Blue

	case ChannelWhite:
		channel = getLED(dt, led).White

	case ChannelStrobe:
		if !dt.StrobeEnabled {
			return cntl.DMXChannel(0), ErrDeviceHasDisabledStrobeChannel
		}
		channel = dt.StrobeChannel

	case ChannelMode:
		if !dt.ModeEnabled {
			return cntl.DMXChannel(0), ErrDeviceHasDisabledModeChannel
		}
		channel = dt.ModeChannel

	case ChannelDimmer:
		if !dt.DimmerEnabled {
			return cntl.DMXChannel(0), ErrDeviceHasDisabledDimmerChannel
		}
		channel = dt.DimmerChannel

	case ChannelTilt:
		if !dt.Moving {
			return cntl.DMXChannel(0), ErrDeviceIsNotMoving
		}
		channel = dt.TiltChannel

	case ChannelPan:
		if !dt.Moving {
			return cntl.DMXChannel(0), ErrDeviceIsNotMoving
		}
		channel = dt.PanChannel

	default:
		return cntl.DMXChannel(0), fmt.Errorf("channel %q is unknown", c)
	}

	return d.StartChannel + channel, nil
}

func getLED(dt *cntl.DMXDeviceType, led uint16) *cntl.LED {
	for _, l := range dt.LEDs {
		if l.Position == led {
			return &l
		}
	}

	return nil
}

// ResolveDeviceSelector returns all DMXDevices that match the given selector
func ResolveDeviceSelector(ds *cntl.DataStore, sel *cntl.DMXDeviceSelector) ([]*cntl.DMXDevice, error) {
	if sel.ID != "" && len(sel.Tags) > 0 {
		return []*cntl.DMXDevice{}, ErrDeviceSelectorCannotHaveTagsAndID
	}

	if sel.ID != "" {
		d, ok := ds.DMXDevices[sel.ID]
		if !ok {
			return []*cntl.DMXDevice{}, fmt.Errorf("unable to find device by id %q", sel.ID)
		}

		return []*cntl.DMXDevice{d}, nil
	}

	if len(sel.Tags) > 0 {
		return ResolveDevicesByTags(ds, sel.Tags), nil
	}

	return []*cntl.DMXDevice{}, ErrDeviceSelectorMustHaveTagsOrID
}

// ResolveDevicesByTags returns all DMXDevices that match *all* of the given tags
func ResolveDevicesByTags(ds *cntl.DataStore, tags []cntl.Tag) (dd []*cntl.DMXDevice) {
	var matches [][]*cntl.DMXDevice

	for _, t := range tags {
		matches = append(matches, ResolveDevicesByTag(ds, t))
	}

	if len(matches) == 0 {
		return []*cntl.DMXDevice{}
	}
	if len(matches) == 1 {
		return matches[0]
	}

	for _, d := range matches[0] {
		var count int
		for i, ds := range matches {
			if i == 0 {
				continue
			}

			if has(ds, d) {
				count++
			}
		}

		if count == len(matches)-1 {
			dd = append(dd, d)
		}

	}

	return
}

// has returns weather the given slice contains the given device
func has(ds []*cntl.DMXDevice, d *cntl.DMXDevice) bool {
	for _, dd := range ds {
		if dd.ID == d.ID {
			return true
		}
	}

	return false
}

// ResolveDevicesByTag returns all DMXDevices that match the given tag
func ResolveDevicesByTag(ds *cntl.DataStore, tag cntl.Tag) (dd []*cntl.DMXDevice) {
	for _, d := range ds.DMXDevices {
		for _, t := range d.Tags {
			if t == tag {
				dd = append(dd, d)
				break
			}
		}
	}

	return
}
