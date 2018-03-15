package dmx

import "github.com/StageAutoControl/controller/cntl"

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
