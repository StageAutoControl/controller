package dmx

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

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
