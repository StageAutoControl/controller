package playback

import "github.com/StageAutoControl/controller/pkg/cntl"

// TransportWriter is a writer to an output stream, for example a websocket or Stdout.
type TransportWriter interface {
	Write(cntl.Command) error
}

// Waiter waits for a trigger to happen
type Waiter interface {
	Wait(done chan struct{}, cancel chan struct{}, err chan error) error
}

type storage interface {
	Has(key string, kind interface{}) bool
	Write(key string, value interface{}) error
	Read(key string, value interface{}) error
	List(kind interface{}) []string
	Delete(key string, kind interface{}) error
}

type loader interface {
	Load() (*cntl.DataStore, error)
}

// Params specifies how to run a playback
type Params struct {
	Song struct {
		ID string
	}
	SetList struct {
		ID string
	}
}

type parsedConfig struct {
	waiters []Waiter
	writers []TransportWriter
}

// Config stores the information on which waiters and/or transport writers are enabled and what their config is
type Config struct {
	Waiters struct {
		Audio struct {
			Enabled   bool
			Threshold float32
		}
	}
	TransportWriters struct {
		ArtNet struct {
			Enabled bool
		}
		MIDI struct {
			Enabled        bool
			OutputDeviceID string
		}
	}
}
