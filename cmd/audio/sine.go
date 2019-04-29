// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package audio

import (
	"math"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/spf13/cobra"
)

var (
	frequency int
	length    int
)

// SineCmd represents the Sines command
var SineCmd = &cobra.Command{
	Use:   "sine",
	Short: "Creates a sin curved audio",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := newStereoSine(float64(frequency), sampleRate)
		defer s.Close()

		if err := s.Start(); err != nil {
			panic(err)
		}
		defer s.Stop()

		time.Sleep(time.Duration(length) * time.Millisecond)
	},
}

func init() {
	AudioCmd.AddCommand(SineCmd)

	SineCmd.PersistentFlags().IntVarP(&frequency, "frequency", "f", 18000, "Frequency of the sin")
	SineCmd.PersistentFlags().IntVarP(&length, "length", "l", 100, "length of the sin in milliseconds")
}

type stereoSine struct {
	*portaudio.Stream
	step, phase float64
}

func newStereoSine(freq, sampleRate float64) *stereoSine {
	s := &stereoSine{nil, freq / sampleRate, 0}

	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, s.processAudio)
	if err != nil {
		panic(err)
	}

	return s
}

func (g *stereoSine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phase))
		_, g.phase = math.Modf(g.phase + g.step)
	}
}
