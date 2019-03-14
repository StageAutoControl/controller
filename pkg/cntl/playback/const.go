package playback

import "errors"

// Player errors
var (
	ErrCancelled = errors.New("playback cancelled")
)

const (
	// ProcessName defines the name of the playback process
	ProcessName      = "playback"
	paramsStorageKey = "playback_process"
)
