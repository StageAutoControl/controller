package playback

import (
	"testing"
	"time"

	"github.com/StageAutoControl/controller/pkg/cntl"
)

func TestCalcRenderSpeed(t *testing.T) {
	exp := []struct {
		bc    cntl.BarChange
		speed time.Duration
	}{
		{cntl.BarChange{Speed: 120, NoteValue: 4}, time.Minute / 1920},
		{cntl.BarChange{Speed: 120, NoteValue: 8}, time.Minute / 1920},
		{cntl.BarChange{Speed: 120, NoteValue: 16}, time.Minute / 1920},
		{cntl.BarChange{Speed: 120, NoteValue: 32}, time.Minute / 1920},

		{cntl.BarChange{Speed: 60, NoteValue: 4}, time.Minute / 960},
		{cntl.BarChange{Speed: 60, NoteValue: 8}, time.Minute / 960},
	}

	for i, e := range exp {
		res := CalcRenderSpeed(&e.bc)
		if res != e.speed {
			t.Errorf("Expected to get duration %q at index %d, got %q", e.speed, i, res)
		}
	}
}
