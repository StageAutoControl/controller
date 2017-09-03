package transport

import (
	"fmt"
	"io"

	"strings"

	"github.com/StageAutoControl/controller/cntl"
)

// Buffer is an output channel that can render to a buffer
type Buffer struct {
	w io.Writer
	i uint64
}

// NewBuffer returns a new Buffer instance
func NewBuffer(w io.Writer) *Buffer {
	fmt.Fprint(w, "          Position  [      BarChange     ] [ Midi ] [ DMX ] \n")

	return &Buffer{w, 0}
}

// Write writes to the buffer
func (b *Buffer) Write(cmd cntl.Command) error {
	fmt.Fprintf(
		b.w, "%s [%s] [%s] [%s]\n",
		renderPosition(b.i),
		renderBarChangeCommand(cmd.BarChange),
		renderMidiCommands(cmd.MIDICommands),
		renderDmxCommands(cmd.DMXCommands),
	)

	b.i++
	return nil
}

func renderPosition(i uint64) string {
	return fmt.Sprintf("%19d", i)
}

func renderBarChangeCommand(bc *cntl.BarChange) string {
	if bc == nil {
		return strings.Repeat(" ", 20)
	}

	return fmt.Sprintf("%20s", fmt.Sprintf("#%d %d/%d @%d bpm", bc.At, bc.NoteCount, bc.NoteValue, bc.Speed))
}

func renderMidiCommands(cmds cntl.MIDICommands) string {
	s := make([]string, len(cmds))

	for i, c := range cmds {
		s[i] = fmt.Sprintf("%d:%d:%d", c.Status, c.Data1, c.Data2)
	}

	return strings.Join(s, " | ")
}

func renderDmxCommands(cmds cntl.DMXCommands) string {
	s := make([]string, len(cmds))

	for i, c := range cmds {
		s[i] = fmt.Sprintf("%d:%d -> %d", c.Universe, c.Channel, c.Value)
	}

	return strings.Join(s, " | ")
}
