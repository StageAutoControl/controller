package waiter

import (
	"fmt"

	"github.com/gordonklaus/portaudio"

	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

// Audio is a waiter that does nothing
type Audio struct {
	logger    logging.Logger
	threshold float32
	notify    chan struct{}
	buf       []float32
	stream    *portaudio.Stream
	cancel    chan struct{}
	err       chan error
}

// NewAudio creates a new Audio waiter
func NewAudio(logger logging.Logger, threshold float32) *Audio {
	return &Audio{
		logger:    logger,
		threshold: threshold,
	}
}

func (a *Audio) start() (err error) {
	a.notify = make(chan struct{}, 1)
	a.buf = make([]float32, 64)
	a.cancel = make(chan struct{}, 1)
	a.err = make(chan error, 5)

	a.stream, err = portaudio.OpenDefaultStream(1, 0, sampleRate, len(a.buf), a.buf)
	if err != nil {
		return fmt.Errorf("failed to open default portaudio stream: %v", err)
	}

	if err := a.stream.Start(); err != nil {
		return fmt.Errorf("failed to start portaudio stream: %v", err)
	}

	go a.readStream()

	return nil
}

func (a *Audio) readStream() {
	for {
		err := a.stream.Read()
		if err != nil {
			a.err <- err
			a.logger.Infof("Error reading portaudio stream: %s", err)
			return
		}

		if a.checkForPeak() {
			return
		}

		select {
		case <-a.cancel:
			return
		default:
		}
	}
}

func (a *Audio) checkForPeak() bool {
	for _, i := range a.buf {
		if i >= a.threshold || i <= (a.threshold*-1) {
			a.notify <- struct{}{}
			return true
		}
	}

	return false
}

// Wait for a peak in the incoming audio stream
func (a *Audio) Wait(done chan struct{}, cancel chan struct{}) error {
	if err := a.start(); err != nil {
		return err
	}

loop:
	for {
		select {
		case <-a.notify:
			done <- struct{}{}
			break loop
		case <-cancel:
			break loop
		case e := <-a.err:
			return e
		}
	}

	return a.stop()
}

// Stop stops the audio stream
func (a *Audio) stop() (err error) {
	a.cancel <- struct{}{}

	if err := a.stream.Stop(); err != nil {
		a.err <- err
		a.logger.Errorf("failed to stop portaudio stream: %v", err)
		// don't return the error, stream.Close has to be called
		// return err
	}

	if err := a.stream.Close(); err != nil {
		a.err <- err
		a.logger.Errorf("failed to close portaudio stream: %v", err)
		return err
	}

	close(a.notify)
	close(a.cancel)
	close(a.err)

	return nil
}
