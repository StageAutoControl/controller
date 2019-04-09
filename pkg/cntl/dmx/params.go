package dmx

import (
	"errors"
	"fmt"

	"github.com/StageAutoControl/controller/pkg/cntl"
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
	if dp.Params != nil && len(dp.Params) > 0 {
		valuesSet++
	}
	if dp.Animation != nil {
		valuesSet++
	}
	if dp.Transition != nil {
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

		g, ok := ds.DMXDeviceGroups[*dp.Group]
		if !ok {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to find DMXDeviceGroup %q", *dp.Group)
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
		d, ok := ds.DMXDevices[*dp.Device]
		if !ok {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to find DMXDevice %q", *dp.Device)
		}

		dd = append(dd, d)
	}

	if len(dd) == 0 {
		return []cntl.DMXCommands{}, ErrDeviceParamsNoDevices
	}

	if dp.Animation != nil {
		a, ok := ds.DMXAnimations[*dp.Animation]
		if !ok {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to find DMXAnimation %q", *dp.Animation)
		}

		return RenderAnimation(ds, dd, a)
	}

	if dp.Transition != nil {
		t, ok := ds.DMXTransitions[*dp.Transition]
		if !ok {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to find DMXTransition %q", *dp.Animation)
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

	if err := resolveColorVar(ds, &p); err != nil {
		return cntl.DMXCommands{}, err
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
	if p.White != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelWhite,
			Value:   *p.White,
		})
	}
	if p.Strobe != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelStrobe,
			Value:   *p.Strobe,
		})
	}
	if p.Mode != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelMode,
			Value:   *p.Mode,
		})
	}
	if p.Dimmer != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelDimmer,
			Value:   *p.Dimmer,
		})
	}
	if p.Tilt != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelTilt,
			Value:   *p.Tilt,
		})
	}
	if p.Pan != nil {
		channels = append(channels, cntl.DMXCommand{
			Channel: ChannelPan,
			Value:   *p.Pan,
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

func resolveColorVar(ds *cntl.DataStore, p *cntl.DMXParams) error {
	if p.ColorVar == nil || *p.ColorVar == "" {
		return nil
	}

	if p.Red != nil || p.Green != nil || p.Blue != nil || p.White != nil {
		return ErrDeviceParamsColorVarMustBeExclusive
	}

	colorVar := getColorVar(ds, *p.ColorVar)
	if colorVar == nil {
		return fmt.Errorf("failed to find color variable with the name %q", *p.ColorVar)
	}

	p.Red = colorVar.Red
	p.Green = colorVar.Green
	p.Blue = colorVar.Blue
	p.White = colorVar.White

	return nil
}

func getColorVar(ds *cntl.DataStore, name string) *cntl.DMXColorVariable {
	for _, c := range ds.DMXColorVariables {
		if c.Name == name {
			return c
		}
	}

	return nil
}
