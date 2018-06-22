package song

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/cntl/dmx"
	"github.com/StageAutoControl/controller/cntl/midi"
)

// Render renders a given SongID to a list of Commands
func Render(ds *cntl.DataStore, songID string) ([]cntl.Command, error) {
	s, ok := ds.Songs[songID]
	if !ok {
		return []cntl.Command{}, fmt.Errorf("cannot find Song %q", songID)
	}

	if s.Length == 0 {
		return []cntl.Command{}, fmt.Errorf("song %q has no length set", songID)
	}

	scs, err := dmx.StreamlineScenes(ds, s)
	if err != nil {
		return []cntl.Command{}, err
	}

	bcs := streamlineBarChanges(s)
	mcs := midi.StreamlineMidiCommands(s)

	fb := &frameBrain{}
	cs := make([]cntl.Command, s.Length)

	for frame := uint64(0); frame < s.Length; frame++ {
		cs[frame].MIDICommands = make(cntl.MIDICommands, 0)
		cs[frame].DMXCommands = make(cntl.DMXCommands, 0)

		if bc, ok := bcs[frame]; ok {
			cs[frame].BarChange = &bc
			fb.setBarChange(&bc)
		}

		fb.update(frame, &cs[frame])

		if mc, ok := mcs[frame]; ok {
			cs[frame].MIDICommands = append(cs[frame].MIDICommands, mc...)
		}

		if scs, ok := scs[frame]; ok {
			for _, sc := range scs {
				dcs, err := dmx.RenderScene(ds, sc)
				if err != nil {
					return []cntl.Command{}, err
				}

				for j, dc := range dcs {
					if len(dc) == 0 {
						continue
					}

					cmdIndex := uint64(j) + frame

					if cmdIndex >= uint64(len(cs)) {
						cs = append(cs, cntl.Command{DMXCommands: dc})
					} else {
						cs[cmdIndex].DMXCommands = append(
							cs[cmdIndex].DMXCommands,
							dc...,
						)
					}
				}
			}
		}
	}

	return cs, nil
}

func streamlineBarChanges(s *cntl.Song) map[uint64]cntl.BarChange {
	bcs := make(map[uint64]cntl.BarChange)
	for _, bc := range s.BarChanges {
		bcs[bc.At] = bc
	}

	return bcs
}

// CalcBarLength calculates the length of a bar by given BarChange
func CalcBarLength(bc *cntl.BarChange) uint64 {
	return uint64(bc.NoteCount) * CalcNoteLength(bc)
}

// CalcNoteLength calculates the amount of frames in a single note of given barChange
func CalcNoteLength(bc *cntl.BarChange) uint64 {
	return uint64(cntl.RenderFrames / bc.NoteValue)
}
