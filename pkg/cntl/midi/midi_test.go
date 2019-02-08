package midi

import (
	"github.com/StageAutoControl/controller/pkg/cntl"
	"testing"
)

func TestStreamlineMidiCommands(t *testing.T) {
	s := &cntl.Song{
		MIDICommands: []cntl.MIDICommand{
			{At: 0, Status: 230, Data1: 123, Data2: 231},
			{At: 200, Status: 100, Data1: 100},
			{At: 200, Status: 200, Data1: 200},
		},
	}

	mcs := StreamlineMidiCommands(s)

	if len(mcs) != 2 {
		t.Errorf("Expected rendered midi map to have length 2, got %d", len(mcs))
	}

	if len(mcs[0]) != 1 {
		t.Errorf("Expected map to have length 1 at index 0, got %d", len(mcs[0]))
	}

	if len(mcs[200]) != 2 {
		t.Errorf("Expected map to have length 2 at index 200, got %d", len(mcs[0]))
	}
}
