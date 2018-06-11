package song

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
)

func TestRender(t *testing.T) {
	ds := fixtures.DataStore()
	_, err := Render(ds, "3c1065c8-0b14-11e7-96eb-5b134621c411")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

}

func TestStreamlineBarChanges(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		s *cntl.Song
		m map[uint64]cntl.BarChange
	}{
		{
			s: ds.Songs["3c1065c8-0b14-11e7-96eb-5b134621c411"],
			m: map[uint64]cntl.BarChange{
				0:    {At: 0, NoteCount: 4, NoteValue: 4, Speed: 160},
				512:  {At: 512, NoteCount: 3, NoteValue: 4},
				1184: {At: 1184, NoteCount: 7, NoteValue: 8},
				1632: {At: 1632, NoteCount: 4, NoteValue: 4},
			},
		},
	}

	for i, e := range exp {
		res := streamlineBarChanges(e.s)

		for k, v := range e.m {
			resv, ok := res[k]
			if !ok {
				t.Errorf("Expected to have key %d at index %d", k, i)
				continue
			}

			if !reflect.DeepEqual(resv, v) {
				t.Errorf("Expected to get value %+v, got %+v", v, resv)
			}

			fmt.Printf("Found correct key %d at index %d \n", k, i)
		}

	}
}

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
