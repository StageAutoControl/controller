package output

import (
	"fmt"
	"io"
)

type BufferOutput struct {
	out io.Writer
}

func NewBufferOutput(out io.Writer) *BufferOutput {
	return &BufferOutput{out}
}

func (o *BufferOutput) PlayStrong() {
	fmt.Fprintln(o.out, "BEEP")
}

func (o *BufferOutput) PlayWeak() {
	fmt.Fprintln(o.out, "beep")
}
