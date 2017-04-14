package transport

import (
	"fmt"
	"io"

	"strings"

	"github.com/StageAutoControl/controller/cntl"
)

// BufferOutput is an output channel that can render to a buffer
type BufferOutput struct {
	w io.Writer
	i uint64
}

// NewBufferTransport returns a new BufferOutput instance
func NewBufferTransport(w io.Writer) *BufferOutput {
	return &BufferOutput{w, 0}
}

// Write writes to the buffer
func (b *BufferOutput) Write(cmd cntl.Command) error {
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
