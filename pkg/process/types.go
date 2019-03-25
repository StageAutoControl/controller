package process

import (
	"context"

	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

// Status of a process as handled by the manager
type Status struct {
	Name      string    `json:"name"`
	Running   bool      `json:"running"`
	StartedAt *JSONTime `json:"startedAt"`
	StoppedAt *JSONTime `json:"stoppedAt"`
	Error     error     `json:"error"`
	Logs      []Log     `json:"logs"`
	Verbose   bool      `json:"verbose"`
}

// Log represents a log line printed by the process
type Log struct {
	Time    JSONTime `json:"time"`
	Level   string   `json:"level"`
	Message string   `json:"message"`
}

// Process carries the information how and what to manage as a process, it implements the custom logic
type Process interface {
	// SetLogger sets the logger of the process, which in fact is a buffering logger
	SetLogger(logger logging.Logger)

	// Start should care about starting the process, including handling errors if the process does not come up.
	// When Start executed without an error the process manager assumes that the process is up and running.
	Start(ctx context.Context) error

	// Stop should fully stop the process and also clean up any leftovers (state, files, go routines, ...)
	Stop() error
}

// Manager to handle a set of processes as pets
type Manager interface {
	AddProcess(name string, process Process, verbose bool) error
	GetProcess(name string) (Process, *Status, error)
	Start(name string) (*Status, error)
	Stop(name string) (*Status, error)
}
