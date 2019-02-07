package transport

import (
	"fmt"
	"io"

	"strings"

	"github.com/sirupsen/logrus"
	"github.com/StageAutoControl/controller/cntl"
)

// Stream is an output channel that can render to a buffer
type Stream struct {
	logger *logrus.Entry
	w io.Writer
	i uint64
}

// NewStream returns a new Stream instance
func NewStream(logger *logrus.Entry, w io.Writer) *Stream {
	fmt.Fprint(w, "          Position  [      BarChange     ] [ Midi ] [ DMX ] \n")

	return &Stream{logger,w, 0}
}

// Write writes to the buffer
func (b *Stream) Write(cmd cntl.Command) error {
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
		s[i] = fmt.Sprintf("%d:%d -> %d", c.Universe, c.Channel, c.Value.Value)
	}

	return strings.Join(s, " | ")
}
