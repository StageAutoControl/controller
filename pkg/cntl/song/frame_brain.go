package song

import "github.com/StageAutoControl/controller/pkg/cntl"

type frameBrain struct {
	bar    uint16
	note   uint8
	lastBC *cntl.BarChange

	framesEachNote uint64
	currentFrame   uint64
}

func (f *frameBrain) setBarChange(bc *cntl.BarChange) {
	f.lastBC = bc
	f.framesEachNote = CalcNoteLength(bc)
	f.bar++
	f.note = 1
	f.currentFrame = 0
}

func (f *frameBrain) update(frame uint64, cmd *cntl.Command) {
	f.currentFrame++

	if f.currentFrame >= f.framesEachNote {
		f.note++
		f.currentFrame = 0

		if f.note > f.lastBC.NoteCount {
			f.note = 0
			f.bar++
		}
	}

	cmd.Frame = frame
	cmd.Bar = f.bar
	cmd.Note = f.note
}
