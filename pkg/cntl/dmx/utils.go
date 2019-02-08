package dmx

import "github.com/StageAutoControl/controller/pkg/cntl"

func repeat(count uint, cmds []cntl.DMXCommands) []cntl.DMXCommands {
	result := make([]cntl.DMXCommands, 0)

	for i := uint(0); i < count; i++ {
		result = append(result, cmds...)
	}

	return result
}

func resize(length uint, cmds []cntl.DMXCommands) []cntl.DMXCommands {
	if length <= uint(len(cmds)) {
		return cmds
	}

	res := make([]cntl.DMXCommands, length)
	copy(res, cmds)

	return res
}
