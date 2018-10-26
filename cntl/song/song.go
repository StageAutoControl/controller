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

	scs, err := dmx.StreamlineScenes(ds, s)
	if err != nil {
		return []cntl.Command{}, err
	}

	bcs := streamlineBarChanges(s)
	mcs := midi.StreamlineMidiCommands(s)

	fb := &frameBrain{}
	numFrames := max(maxKey(scs), maxKey(mcs))
	cs := makeCommandArray(numFrames)

	for frame := uint64(0); frame < numFrames; frame++ {
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
						indexDiff := cmdIndex - uint64(len(cs)) + 1
						cs = append(cs, makeCommandArray(indexDiff)...)
					}

					cs[cmdIndex].DMXCommands = append(
						cs[cmdIndex].DMXCommands,
						dc...,
					)
				}
			}
		}
	}

	return cs, nil
}
