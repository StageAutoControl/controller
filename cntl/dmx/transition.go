package dmx

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

// RenderTransition renders the given DMXTransition to an array of DMXCommands to be sent to a DMX device
func RenderTransition(ds *cntl.DataStore, dd []*cntl.DMXDevice, t *cntl.DMXTransition) ([]cntl.DMXCommands, error) {
	cmds := make([]cntl.DMXCommands, t.Length)

	for i, p := range t.Params {
		if p.From.LED != p.To.LED {
			return []cntl.DMXCommands{}, ErrTransitionDeviceParamsMustMatchLED
		}

		paramCMDs, err := RenderTransitionParams(ds, dd, p)
		if err != nil {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to render animation transition %q param %d: %v", t.ID, i, err)
		}

		cmds = Merge(cmds, paramCMDs)
	}

	return cmds, nil
}

func RenderTransitionParams(ds *cntl.DataStore, dd []*cntl.DMXDevice, p cntl.DMXTransitionParams) ([]cntl.DMXCommands, error) {

	return []cntl.DMXCommands{}, nil
}
