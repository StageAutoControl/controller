package dmx

import (
	"errors"
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

// Render Errors
var (
	ErrDeviceParamsDevicesInvalid        = errors.New("DMXDeviceParams can only have a group or a Device")
	ErrDeviceParamsValuesInvalid         = errors.New("DMXDeviceParams must not have more the one of [Animation, Transition, Params]")
	ErrDeviceParamsNoDevices             = errors.New("DMXDeviceParams matches no device")
	ErrDeviceSelectorMustHaveTagsOrID    = errors.New("DMXDeviceSelector must have either tags or an ID")
	ErrDeviceSelectorCannotHaveTagsAndID = errors.New("DMXDeviceSelector cannot have tags and an ID")
)

// StreamlineScenes returns a map of frame -> scene that is easier to handle then a plain array
func StreamlineScenes(ds *cntl.DataStore, s *cntl.Song) (map[uint64][]*cntl.DMXScene, error) {
	scs := make(map[uint64][]*cntl.DMXScene)
	for _, sp := range s.DMXScenes {
		sc, ok := ds.DMXScenes[sp.ID]
		if !ok {
			return map[uint64][]*cntl.DMXScene{}, fmt.Errorf("cannot find DMXScene %q", sp.ID)
		}

		l := CalcSceneLength(sc)
		at := uint64(sp.At)

		for i := uint64(0); i <= uint64(sp.Repeat); i++ {
			pos := at + i*l
			if _, ok := scs[pos]; !ok {
				scs[pos] = make([]*cntl.DMXScene, 0)
			}

			scs[pos] = append(scs[pos], sc)
		}
	}

	return scs, nil
}

// CalcSceneLength calculates the length of a given scene in render frames
func CalcSceneLength(sc *cntl.DMXScene) uint64 {
	return uint64(sc.NoteCount * (cntl.RenderFrames / sc.NoteValue))
}

// RenderScene renders the given dmx scene to dmx commands and returns them.
// The first array dimension contains the render frames, the second dimension contains all
// dmx commands for a render frame.
func RenderScene(ds *cntl.DataStore, sc *cntl.DMXScene) ([]cntl.DMXCommands, error) {
	sceneLength := uint8(CalcSceneLength(sc))
	cmds := make([]cntl.DMXCommands, sceneLength)

	for i, ss := range sc.SubScenes {
		var scs []cntl.DMXCommands

		if len(ss.DeviceParams) > 0 && ss.Preset != "" {
			return []cntl.DMXCommands{}, fmt.Errorf("SubScene %d of scene %q cannot have both params and a preset", i, sc.ID)
		}

		if ss.Preset != "" {
			p, ok := ds.DMXPresets[ss.Preset]
			if !ok {
				return []cntl.DMXCommands{}, fmt.Errorf("cannot find DMXPreset %q", ss.Preset)
			}

			pcs, err := RenderPreset(ds, p)
			if err != nil {
				return []cntl.DMXCommands{}, err
			}

			scs = Merge(scs, pcs)
		}

		for _, dp := range ss.DeviceParams {
			dcs, err := RenderDeviceParams(ds, &dp)
			if err != nil {
				return []cntl.DMXCommands{}, fmt.Errorf("failed to render scene %q: %v", sc.ID, err)
			}

			scs = Merge(scs, dcs)
		}

		for _, at := range ss.At {
			pos := uint64(sceneLength/sc.NoteCount) * at
			cmds = MergeAtOffset(cmds, scs, int(pos))
		}
	}

	return cmds, nil
}

// RenderPreset renders a preset and returns an array of commands for every frame
func RenderPreset(ds *cntl.DataStore, p *cntl.DMXPreset) ([]cntl.DMXCommands, error) {
	var cmds []cntl.DMXCommands
	for _, dp := range p.DeviceParams {
		dpcs, err := RenderDeviceParams(ds, &dp)
		if err != nil {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to handle preset %q: %v", p.ID, err)
		}

		cmds = Merge(cmds, dpcs)
	}

	return cmds, nil
}

// Merge merges two arrays of DMXCommands
func Merge(cmds []cntl.DMXCommands, cs []cntl.DMXCommands) []cntl.DMXCommands {
	return MergeAtOffset(cmds, cs, 0)
}

// MergeAtOffset merges two arrays of DMXCommands after a given offset
func MergeAtOffset(cmds []cntl.DMXCommands, cs []cntl.DMXCommands, offset int) []cntl.DMXCommands {
	for i, cs := range cs {
		index := i + offset
		if index > len(cmds)-1 {
			cmds = append(cmds, cs)
			continue
		}

		cmds[index] = append(cmds[index], cs...)
	}
	return cmds
}

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
			return []cntl.DMXCommands{}, fmt.Errorf("Unable to find DMXAnimation %q", dp.AnimationID)
		}

		return RenderAnimation(ds, dd, a)
	}

	if dp.Params != nil {
		c, err := RenderParams(ds, dd, *dp.Params)
		if err != nil {
			return []cntl.DMXCommands{}, err
		}

		return []cntl.DMXCommands{c}, nil
	}

	return []cntl.DMXCommands{}, errors.New("This code should be unreachable. If you see this message please reset the world spin.")
}

// RenderParams renders the given DMXParams to an array of DMXCommands to be sent to a DMX device
func RenderParams(ds *cntl.DataStore, dd []*cntl.DMXDevice, p cntl.DMXParams) (cmds cntl.DMXCommands, err error) {
	var channels cntl.DMXCommands

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

// RenderAnimation renders the given DMXAnimation to an array of DMXCommands to be sent to a DMX device
func RenderAnimation(ds *cntl.DataStore, dd []*cntl.DMXDevice, a *cntl.DMXAnimation) ([]cntl.DMXCommands, error) {
	cmds := make([]cntl.DMXCommands, a.Length)
	for _, f := range a.Frames {
		ps, err := RenderParams(ds, dd, f.Params)
		if err != nil {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to render animation %q: %v", a.ID, err)
		}

		cmds[f.At] = append(cmds[f.At], ps...)
	}

	return cmds, nil
}

// ResolveDeviceSelector returns all DMXDevices that match the given selector
func ResolveDeviceSelector(ds *cntl.DataStore, sel *cntl.DMXDeviceSelector) ([]*cntl.DMXDevice, error) {
	if sel.ID != "" && len(sel.Tags) > 0 {
		return []*cntl.DMXDevice{}, ErrDeviceSelectorCannotHaveTagsAndID
	}

	if sel.ID != "" {
		d, ok := ds.DMXDevices[sel.ID]
		if !ok {
			return []*cntl.DMXDevice{}, fmt.Errorf("Unable to find device by id %q", sel.ID)
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
