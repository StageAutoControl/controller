package dmx

import (
	"errors"
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

// checkDeviceParams checks a DeviceParams entity to be valid in terms of devices and values
func checkDeviceParams(dp *cntl.DMXDeviceParams) error {
	devicesSet := 0
	if dp.Device != nil {
		devicesSet++
	}
	if dp.Group != nil {
		devicesSet++
	}
	if devicesSet != 1 {
		return ErrDeviceParamsDevicesInvalid
	}

	valuesSet := 0
	if dp.Params != nil {
		valuesSet++
	}
	if dp.AnimationID != "" {
		valuesSet++
	}
	if dp.TransitionID != "" {
		valuesSet++
	}

	if valuesSet != 1 {
		return ErrDeviceParamsValuesInvalid
	}

	return nil
}

// RenderDeviceParams renders the given DMXDeviceParams to an array of DMXCommands to be sent to a DMX device
func RenderDeviceParams(ds *cntl.DataStore, dp *cntl.DMXDeviceParams) ([]cntl.DMXCommands, error) {
	if err := checkDeviceParams(dp); err != nil {
		return []cntl.DMXCommands{}, err
	}

	var dd []*cntl.DMXDevice
	if dp.Group != nil {
		g, ok := ds.DMXDeviceGroups[dp.Group.ID]
		if !ok {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to find DMXDeviceGroup %q", dp.Group)
		}

		for _, sel := range g.Devices {
			d, err := ResolveDeviceSelector(ds, &sel)
			if err != nil {
				return []cntl.DMXCommands{}, err
			}

			dd = append(dd, d...)
		}
	}
	if dp.Device != nil {
		d, err := ResolveDeviceSelector(ds, dp.Device)
		if err != nil {
			return []cntl.DMXCommands{}, err
		}

		dd = append(dd, d...)
	}

	if len(dd) == 0 {
		return []cntl.DMXCommands{}, ErrDeviceParamsNoDevices
	}

	if dp.AnimationID != "" {
		a, ok := ds.DMXAnimations[dp.AnimationID]
		if !ok {
			return []cntl.DMXCommands{}, fmt.Errorf("unable to find DMXAnimation %q", dp.AnimationID)
		}

		return RenderAnimation(ds, dd, a)
	}

	if dp.TransitionID != "" {
		t, ok := ds.DMXTransitions[dp.TransitionID]
		if !ok {
			return []cntl.DMXCommands{}, fmt.Errorf("unable to find DMXTransition %q", dp.AnimationID)
		}

		return RenderTransition(ds, dd, t)
	}

	if dp.Params != nil {
		cs := make([]cntl.DMXCommands, 1)

		for _, p := range dp.Params {
			c, err := RenderParams(ds, dd, p)
			if err != nil {
				return []cntl.DMXCommands{}, err
			}

			cs[0] = append(cs[0], c...)
		}

		return cs, nil
	}

	return []cntl.DMXCommands{}, errors.New("this code should be unreachable. If you see this message please reset the world spin")
}

// RenderParams renders the given DMXParams to an array of DMXCommands to be sent to a DMX device
func RenderParams(ds *cntl.DataStore, dd []*cntl.DMXDevice, p cntl.DMXParams) (cmds cntl.DMXCommands, err error) {
	var channels cntl.DMXCommands

	if p.White != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelWhite,
			Value:   *p.White,
		})
	}
	if p.Red != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelRed,
			Value:   *p.Red,
		})
	}
	if p.Green != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelGreen,
			Value:   *p.Green,
		})
	}
	if p.Blue != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelBlue,
			Value:   *p.Blue,
		})
	}
	if p.Strobe != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelStrobe,
			Value:   *p.Strobe,
		})
	}
	if p.Preset != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelMode,
			Value:   *p.Preset,
		})
	}

	for _, d := range dd {
		for _, c := range channels {
			ch, err := getDeviceChannel(ds, d, c.Channel, p.LED)
			if err != nil {
				return cntl.DMXCommands{}, err
			}
			cmds = append(cmds, cntl.DMXCommand{
				Universe: d.Universe,
				Channel:  ch,
				Value:    c.Value,
			})
		}
	}

	return
}
