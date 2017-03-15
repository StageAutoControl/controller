package input

import (
	"os"

	"github.com/apinnecke/dmx-auto-control/sample"
	"github.com/gordonklaus/portaudio"
)

type AudioReader struct {
	in *sample.Buffer
}

func NewAudioReader(in *sample.Buffer) *AudioReader {
	return &AudioReader{
		in: in,
	}
}

func (r *AudioReader) Read(sig chan os.Signal) (err error) {
	portaudio.Initialize()
	defer portaudio.Terminate()

	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, 64, in)
	if err != nil {
		return err
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		return err
	}

	for {
		err = stream.Read()
		if err != nil {
			return err
		}

		r.in.WriteAll(in)
		if err != nil {
			return err
		}

		select {
		case <-sig:
			return
		default:
		}
	}

	err = stream.Stop()
	return
}
