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
		return []cntl.Command{}, fmt.Errorf("Cannot find Song %q", songID)
	}

	if s.Length == 0 {
		return []cntl.Command{}, fmt.Errorf("Song %q has no length set", songID)
	}

	scs, err := dmx.StreamlineScenes(ds, s)
	if err != nil {
		return []cntl.Command{}, err
	}

	bcs := streamlineBarChanges(s)
	mcs := midi.StreamlineMidiCommands(s)

	cs := make([]cntl.Command, s.Length)
	for i := uint64(0); i < s.Length; i++ {

		if bc, ok := bcs[i]; ok {
			cs[i].BarChange = &bc
		}

		if mc, ok := mcs[i]; ok {
			cs[i].MIDICommands = append(cs[i].MIDICommands, mc...)
		}

		if scs, ok := scs[i]; ok {
			for _, sc := range scs {
				dcs, err := dmx.RenderScene(ds, sc)
				if err != nil {
					return []cntl.Command{}, err
				}

				for j, dc := range dcs {
					if len(dc) == 0 {
						continue
					}

					cmdIndex := uint64(j) + i

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
	return uint64(bc.NoteCount * (cntl.RenderFrames / bc.NoteValue))
}
