package transport

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

type logger interface {
	Print(args ...interface{})
}

// BarLogger logs every new bar to given logger
type BarLogger struct {
	logger logger
	state cntl.FrameState
}

// NewBarLogger returns a new logger instance
func NewBarLogger(logger logger) *BarLogger {
	return &BarLogger{
		logger: logger,
	}
}

func (l *BarLogger) Write(cmd cntl.Command) error {
	l.state.Frame = cmd.Frame

	if cmd.Bar != l.state.Bar {
		l.state.Bar = cmd.Bar
	}

	if cmd.Note != l.state.Note {
		l.state.Note = cmd.Note

		l.logger.Print(fmt.Sprintf("[Frame %10d] %4d.%d", l.state.Frame, l.state.Bar, l.state.Note))
	}


	return nil
}
