package song

import (
	"testing"

	"github.com/StageAutoControl/controller/cntl"
)

func TestCalcBarLength(t *testing.T) {
	exp := []struct {
		bc     cntl.BarChange
		length uint64
	}{
		{cntl.BarChange{At: 0, NoteCount: 3, NoteValue: 4}, 48},
		{cntl.BarChange{At: 63, NoteCount: 12, NoteValue: 8}, 96},
		{cntl.BarChange{At: 10, NoteCount: 11, NoteValue: 4}, 176},
		{cntl.BarChange{At: 104, NoteCount: 4, NoteValue: 4}, 64},
		{cntl.BarChange{At: 5, NoteCount: 9, NoteValue: 8}, 72},
	}

	for i, e := range exp {
		if l := CalcBarLength(&e.bc); l != e.length {
			t.Errorf("Expected to get length %d, got %d for index %d", e.length, l, i)
		}
	}
}
