package process

import "errors"

var (
	errProcessNotFound       = errors.New("the process with given name was not found or isn't running")
	errProcessAlreadyExists  = errors.New("the process with the given name already exists")
	errProcessAlreadyRunning = errors.New("the process with the given name is already running")
	errProcessNotRunning     = errors.New("the process with the given name is not running")
)
