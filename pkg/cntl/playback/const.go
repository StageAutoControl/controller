package playback

import "errors"

// Player errors
var (
	ErrCancelled                = errors.New("playback cancelled")
	ErrNoSongIDOrSetListIDGiven = errors.New("no songID or setListID given")
)

const (
	// ProcessName defines the name of the playback process
	ProcessName      = "playback"
	paramsStorageKey = "playback_process"
)

var defaultConfig = `
{
  "waiters": {
    "audio": {
      "enabled": true,
			"threshold": 0.7
    }
  },
  "transportWriters": {
    "artNet": {
      "enabled": true
    },
    "visualizer": {
      "enabled": true
    },
    "midi": {
      "enabled": false,
      "outputDeviceId": 0
    }
  }
}
`
