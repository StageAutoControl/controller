package metronome

import "fmt"

type Bar struct {
	Beats     uint
	NoteValue uint
	Tempo     uint
}

func NewBar(beats, noteValue, tempo uint) *Bar {
	return &Bar{
		Beats:     beats,
		NoteValue: noteValue,
		Tempo:     tempo,
	}
}

func (b *Bar) String() string {
	return fmt.Sprintf("%d/%d @ %d BPM", b.Beats, b.NoteValue, b.Tempo)
}
