package waiter

import (
	"log"

	"github.com/gordonklaus/portaudio"
)

// Audio is a waiter that does nothing
type Audio struct {
	threshold int32
	fanOut    []chan struct{}
	in        []int32
	stream    *portaudio.Stream
	stop      chan struct{}
}

// NewAudio creates a new Audio waiter
func NewAudio(threshold int32) (*Audio, error) {
	portaudio.Initialize()

	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	if err != nil {
		return nil, err
	}

	err = stream.Start()
	if err != nil {
		return nil, err
	}

	a := &Audio{
		threshold: threshold,
		fanOut:    make([]chan struct{}, 0),
		in:        in,
		stream:    stream,
		stop:      make(chan struct{}, 1),
	}

	go a.readStream()

	return a, nil
}

func (a *Audio) readStream() {
	for {
		err := a.stream.Read()
		if err != nil {
			log.Printf("Error reading audio stream: %s\n", err)
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
	for _, i := range a.in {
		if i >= a.threshold {
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
	defer func() {
		a.fanOut = a.fanOut[:len(a.fanOut)-1]
	}()

	for {
		select {
		case <-waitForPeak:
			done <- struct{}{}
		case <-cancel:
			return nil
		}

		return nil
	}
}

// Stop stops the audio stream
func (a *Audio) Stop() (err error) {
	defer portaudio.Terminate()

	return a.stream.Close()
}
