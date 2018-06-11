package dmx

import "github.com/StageAutoControl/controller/cntl"

func repeat(count uint, cmds []cntl.DMXCommands) []cntl.DMXCommands {
	result := make([]cntl.DMXCommands, 0)

	for i := uint(0); i < count; i++ {
		result = append(result, cmds...)
	}

	return result
}
