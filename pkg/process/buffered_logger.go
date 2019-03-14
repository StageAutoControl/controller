package process

import (
	"fmt"
	"time"
)

// BufferedLogger appends the logs to a given slice of Log entries which is passed by reference
type BufferedLogger struct {
	logs    *[]Log
	verbose bool
}

// NewBufferedLogger returns a new BufferedLogger instance
func NewBufferedLogger(logs *[]Log, verbose bool) *BufferedLogger {
	return &BufferedLogger{
		logs:    logs,
		verbose: verbose,
	}
}

func (l *BufferedLogger) log(level, msg string) {
	if !l.verbose && level == "debug" {
		return
	}

	*l.logs = append(*l.logs, Log{
		Time:    JSONTime{Time: time.Now()},
		Level:   level,
		Message: msg,
	})
}

// Debugf log method
func (l *BufferedLogger) Debugf(format string, args ...interface{}) {
	l.log("debug", fmt.Sprintf(format, args...))
}

// Infof log method
func (l *BufferedLogger) Infof(format string, args ...interface{}) {
	l.log("info", fmt.Sprintf(format, args...))
}

// Warnf log method
func (l *BufferedLogger) Warnf(format string, args ...interface{}) {
	l.log("warn", fmt.Sprintf(format, args...))
}

// Warningf log method
func (l *BufferedLogger) Warningf(format string, args ...interface{}) {
	l.log("warning", fmt.Sprintf(format, args...))
}

// Errorf log method
func (l *BufferedLogger) Errorf(format string, args ...interface{}) {
	l.log("error", fmt.Sprintf(format, args...))
}

// Debug log method
func (l *BufferedLogger) Debug(args ...interface{}) {
	l.log("debug", fmt.Sprint(args...))
}

// Info log method
func (l *BufferedLogger) Info(args ...interface{}) {
	l.log("info", fmt.Sprint(args...))
}

// Warn log method
func (l *BufferedLogger) Warn(args ...interface{}) {
	l.log("warn", fmt.Sprint(args...))
}

// Warning log method
func (l *BufferedLogger) Warning(args ...interface{}) {
	l.log("warning", fmt.Sprint(args...))
}

// Error log method
func (l *BufferedLogger) Error(args ...interface{}) {
	l.log("error", fmt.Sprint(args...))
}
