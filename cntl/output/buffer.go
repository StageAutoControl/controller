package output

import (
	"fmt"
	"io"

	"strings"

	"github.com/StageAutoControl/controller/cntl"
)

// BufferOutput is an output channel that can render to a buffer
type BufferOutput struct {
	w io.Writer
}

// NewBufferOutput returns a new BufferOutput instance
func NewBufferOutput(w io.Writer) *BufferOutput {
	return &BufferOutput{w}
}

// Write writes to the buffer
func (b *BufferOutput) Write(cmd cntl.Command) {
	fmt.Fprintf(b.w, "%+v, %+v, %s\n", cmd.BarChange, cmd.MIDICommands, renderCommands(cmd.DMXCommands))
}
func renderCommands(cmds cntl.DMXCommands) string {
	s := make([]string, len(cmds))

	for i, c := range cmds {
		s[i] = fmt.Sprintf("%d:%d -> %d", c.Universe, c.Channel, c.Value)
	}

	return strings.Join(s, " | ")
}
