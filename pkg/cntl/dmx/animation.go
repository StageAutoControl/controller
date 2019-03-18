package dmx

import (
	"fmt"

	"github.com/StageAutoControl/controller/pkg/cntl"
)

// RenderAnimation renders the given DMXAnimation to an array of DMXCommands to be sent to a DMX device
func RenderAnimation(ds *cntl.DataStore, dd []*cntl.DMXDevice, a *cntl.DMXAnimation) ([]cntl.DMXCommands, error) {
	cmds := make([]cntl.DMXCommands, maxFrame(a))
	for _, f := range a.Frames {
		ps, err := RenderParams(ds, dd, f.Params)
		if err != nil {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to render animation %q: %v", a.ID, err)
		}

		cmds[f.At] = append(cmds[f.At], ps...)
	}

	return cmds, nil
}

func maxFrame(a *cntl.DMXAnimation) uint8 {
	var max uint8
	for _, f := range a.Frames {
		if f.At > max {
			max = f.At
		}
	}

	return max
}
