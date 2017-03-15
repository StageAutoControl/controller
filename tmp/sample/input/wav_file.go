package input

import (
	"io"
	"os"

	"github.com/apinnecke/dmx-auto-control/sample"
	"github.com/cryptix/wav"
)

// WavFileReader reads the samples from a wav file
type WavFileReader struct {
	in        *sample.Buffer
	wavReader *wav.Reader
}

// NewWavFileReader returns a new NewWavFileReader and checks weather the given file is even valid
func NewWavFileReader(file string, in *sample.Buffer) (r *WavFileReader, err error) {
	testInfo, err := os.Stat(file)
	if err != nil {
		return
	}

	testWav, err := os.Open(file)
	if err != nil {
		return
	}

	wavReader, err := wav.NewReader(testWav, testInfo.Size())
	if err != nil {
		return
	}

	r = &WavFileReader{
		in:        in,
		wavReader: wavReader,
	}
	return
}

// Read reads the files content to the buffer
func (r *WavFileReader) Read(sig chan os.Signal) (err error) {
	meta := r.wavReader.GetFile()
	bytesPerSample := uint32(meta.SignificantBits / 8)

	for {
		var b []byte
		b, err = r.wavReader.ReadRawSample()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return
		}

		i := toInt(b, bytesPerSample)
		s := checkNegative(i, bytesPerSample)
		r.in.Write(s)

		select {
		case <-sig:
			return
		default:
		}
	}

	return
}

// ReadOrPanic reads the content, panic in case of error
func (r *WavFileReader) ReadOrPanic() {
	err := r.Read()
	if err != nil {
		panic(err)
	}
}
