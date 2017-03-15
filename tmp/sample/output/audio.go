package output

import (
	"sync"

	"github.com/gordonklaus/portaudio"
)

type AudioOutput struct {
	sampleRate uint
	out        []int32
	m          sync.Mutex
}

func NewAudio() {

}

func (a *AudioOutput) Start() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	stream, err := portaudio.OpenDefaultStream(0, 1, float64(a.sampleRate), 0, a.processAudio)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		panic(err)
	}
	defer stream.Stop()
}

func (a *AudioOutput) processAudio(out []int32) {
	for i := range out {
		if len(a.out) == 0 {
			out[i] = 0
			continue
		}

	}
}
