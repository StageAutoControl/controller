package waiter

import (
	"github.com/sirupsen/logrus"
	"github.com/gordonklaus/portaudio"
)

// Audio is a waiter that does nothing
type Audio struct {
	logger    *logrus.Entry
	threshold float32
	fanOut    []chan struct{}
	buf       []float32
	stream    *portaudio.Stream
	stop      chan struct{}
	err       chan error
}

// NewAudio creates a new Audio waiter
func NewAudio(logger *logrus.Entry, threshold float32) (*Audio, error) {
	portaudio.Initialize()

	buf := make([]float32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buf), buf)
	if err != nil {
		return nil, err
	}

	a := &Audio{
		logger:    logger,
		threshold: threshold,
		fanOut:    make([]chan struct{}, 0),
		buf:       buf,
		stream:    stream,
		stop:      make(chan struct{}, 1),
		err:       make(chan error, 1),
	}

	go a.readStream()

	return a, nil
}

func (a *Audio) readStream() {
	if err := a.stream.Start(); err != nil {
		a.err <- err
		return
	}

	defer func() {
		if err := a.stream.Stop(); err != nil {
			a.err <- err
			return
		}
	}()

	for {
		err := a.stream.Read()
		if err != nil {
			a.logger.Infof("Error reading audio stream: %s", err)
			return
		}

		a.checkForPeak()

		select {
		case <-a.stop:
			return
		default:
		}
	}
}

func (a *Audio) checkForPeak() {
	for _, i := range a.buf {
		if i >= a.threshold || i <= (a.threshold * -1) {
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

// Wait waits for a specific event to happen. In this case, nothing.
func (a *Audio) Wait(done chan struct{}, cancel chan struct{}, err chan error) error {
	waitForPeak := make(chan struct{}, 1)
	a.fanOut = append(a.fanOut, waitForPeak)

	// remove channel from fanout, we don't want to have further updates
	defer func() {
		a.fanOut = a.fanOut[:len(a.fanOut)-1]
	}()

	for {
		select {
		case <-waitForPeak:
			a.logger.Info("Found peak. Starting playback!")
			done <- struct{}{}
			return nil
		case <-cancel:
			return nil
		case err := <-a.err:
			return err
		}

		return nil
	}
}

// Stop stops the audio stream
func (a *Audio) Stop() (err error) {
	a.stop <- struct{}{}

	defer portaudio.Terminate()

	return a.stream.Close()
}
