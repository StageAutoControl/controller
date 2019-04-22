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
	fanOut    []chan struct{}
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
		fanOut:    make([]chan struct{}, 0),
		buf:       make([]float32, 64),
		cancel:    make(chan struct{}, 1),
		err:       make(chan error, 5),
	}
}

func (a *Audio) start() (err error) {
	if err := portaudio.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize portaudio: %v", err)
	}

	a.stream, err = portaudio.OpenDefaultStream(1, 0, sampleRate, len(a.buf), a.buf)
	if err != nil {
		return fmt.Errorf("failed to open default portaudio stream: %v", err)
	}

	go a.readStream()

	return nil
}

func (a *Audio) readStream() {
	if err := a.stream.Start(); err != nil {
		a.err <- err
		return
	}

	for {
		err := a.stream.Read()
		if err != nil {
			a.logger.Infof("Error reading audio stream: %s", err)
			return
		}

		a.checkForPeak()

		select {
		case <-a.cancel:
			return

		default:
		}
	}
}

func (a *Audio) checkForPeak() {
	for _, i := range a.buf {
		if i >= a.threshold || i <= (a.threshold*-1) {
			a.notifyWait()
			return
		}
	}
}

func (a *Audio) notifyWait() {
	for _, c := range a.fanOut {
		c <- struct{}{}
	}
}

// Wait for a peak in the incoming audio stream
func (a *Audio) Wait(done chan struct{}, cancel chan struct{}) error {
	if err := a.start(); err != nil {
		return err
	}

	waitForPeak := make(chan struct{}, 1)
	a.fanOut = append(a.fanOut, waitForPeak)

loop:
	for {
		select {
		case <-waitForPeak:
			a.logger.Info("Found peak. Starting playback!")
			done <- struct{}{}
			break loop
		case <-cancel:
			break loop
		case e := <-a.err:
			return e
		}
	}

	a.fanOut = a.fanOut[:len(a.fanOut)-1]
	return a.stop()
}

// Stop stops the audio stream
func (a *Audio) stop() (err error) {
	a.cancel <- struct{}{}

	if err := a.stream.Abort(); err != nil {
		a.err <- err
		a.logger.Errorf("failed to stop portaudio stream: %v", err)
	}

	if err := a.stream.Close(); err != nil {
		a.err <- err
		a.logger.Errorf("failed to close portaudio stream: %v", err)
	}

	if err := portaudio.Terminate(); err != nil {
		a.logger.Errorf("failed to terminate portaudio: %v", err)
		a.err <- err
	}

	return nil
}
