package dmx

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

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
