package song

import (
	"errors"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"log"
	"reflect"
)

// max returns the bigger of two given uint64 values
func max(x, y uint64) uint64 {
	if x < y {
		return y
	}
	return x
}

func uint64Keys(v interface{}) []uint64 {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		panic(errors.New("not a map"))
	}

	t := rv.Type()
	if t.Key().Kind() != reflect.Uint64 {
		panic(errors.New("not Uint64 key"))
	}

	var result []uint64
	for _, kv := range rv.MapKeys() {
		result = append(result, kv.Uint())
	}

	return result
}

// maxKey receives a map index with uint64 and returns the biggest key
func maxKey(search interface{}) uint64 {
	keys := uint64Keys(search)
	var biggest uint64

	for _, key := range keys {
		if key > biggest {
			biggest = key
		}
	}

	return biggest
}

func makeCommand() cntl.Command {
	return cntl.Command{
		MIDICommands: make([]cntl.MIDICommand, 0),
		DMXCommands:  make([]cntl.DMXCommand, 0),
	}
}

func makeCommandArray(length uint64) []cntl.Command {
	log.Print(length)
	cmds := make([]cntl.Command, length)

	for i := range cmds {
		cmds[i].MIDICommands = make([]cntl.MIDICommand, 0)
		cmds[i].DMXCommands = make([]cntl.DMXCommand, 0)
	}
	return cmds
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
