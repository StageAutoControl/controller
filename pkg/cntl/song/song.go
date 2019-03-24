package song

import (
	"fmt"

	"github.com/StageAutoControl/controller/pkg/cntl"

	"github.com/StageAutoControl/controller/pkg/cntl/dmx"
	"github.com/StageAutoControl/controller/pkg/cntl/midi"
)

// Render renders a given SongID to a list of Commands
func Render(ds *cntl.DataStore, songID string) ([]cntl.Command, error) {
	s, ok := ds.Songs[songID]
	if !ok {
		return nil, fmt.Errorf("cannot find Song %q", songID)
	}

	scs, err := dmx.StreamlineScenes(ds, s)
	if err != nil {
		return nil, err
	}

	bcs := StreamlineBarChanges(s)
	if err := ValidateBarChanges(bcs); err != nil {
		return nil, fmt.Errorf("failed to validate bar changes: %v", err)
	}

	mcs := midi.StreamlineMidiCommands(s)

	fb := &frameBrain{}
	numFrames := max(maxKey(scs), maxKey(mcs)) + 1
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
					return nil, err
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
