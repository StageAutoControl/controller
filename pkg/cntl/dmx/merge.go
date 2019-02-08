package dmx

import "github.com/StageAutoControl/controller/pkg/cntl"

// Merge merges two arrays of DMXCommands
func Merge(cmds []cntl.DMXCommands, cs []cntl.DMXCommands) []cntl.DMXCommands {
	return MergeAtOffset(cmds, cs, 0)
}

// MergeAtOffset merges two arrays of DMXCommands after a given offset
func MergeAtOffset(cmds []cntl.DMXCommands, cs []cntl.DMXCommands, offset int) []cntl.DMXCommands {
	for sourceIndex, cs := range cs {
		targetIndex := sourceIndex + offset
		targetLength := targetIndex + 1

		if targetLength > len(cmds) {
			cmds = resize(uint(targetLength), cmds)
		}

		cmds[targetIndex] = append(cmds[targetIndex], cs...)
	}
	return cmds
}

// MergeWithFrameChange merges two given DMXCommand slices, respecting that the second one is not in the right renderFrames setting
// Like, cs is in 16th and RenderFrames ist 64th we have to add spacing while merging
func MergeWithFrameChange(cmds []cntl.DMXCommands, cs []cntl.DMXCommands, paramsNoteValue uint8) []cntl.DMXCommands {
	spacing := int(cntl.RenderFrames / paramsNoteValue)
	for sourceIndex, cmd := range cs {
		cmds = MergeAtOffset(cmds, []cntl.DMXCommands{cmd}, sourceIndex*spacing)
	}

	return cmds
}
