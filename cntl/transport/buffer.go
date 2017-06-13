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
	return &Buffer{w, 0}
}

// Write writes to the buffer
func (b *Buffer) Write(cmd cntl.Command) error {
	fmt.Fprintf(b.w, "%d %+v, %+v, %s\n", b.i, cmd.BarChange, cmd.MIDICommands, renderCommands(cmd.DMXCommands))
	b.i++
	return nil
}

func renderCommands(cmds cntl.DMXCommands) string {
	s := make([]string, len(cmds))

	for i, c := range cmds {
		s[i] = fmt.Sprintf("%d:%d -> %d", c.Universe, c.Channel, c.Value)
	}

	return strings.Join(s, " | ")
}
