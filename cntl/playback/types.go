package playback

import "github.com/StageAutoControl/controller/cntl"

// TransportWriter is a writer to an output stream, for example a websocket or Stdout.
type TransportWriter interface {
	Write(cntl.Command) error
}

// Waiter waits for a trigger to happen
type Waiter interface {
	Wait(done chan struct{}, cancel chan struct{}, err chan error) error
}
