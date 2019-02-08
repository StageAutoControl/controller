package midi

import (
	"github.com/StageAutoControl/controller/pkg/cntl"
)

// StreamlineMidiCommands streamlines a given array of MidiCommands into a map of beats to MidiCommands
func StreamlineMidiCommands(s *cntl.Song) map[uint64]cntl.MIDICommands {
	mcs := make(map[uint64]cntl.MIDICommands)
	for _, bc := range s.MIDICommands {
		mcs[bc.At] = append(mcs[bc.At], bc)
	}

	return mcs
}
