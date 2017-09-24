// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package audio

import (
	"fmt"
	"os"

	"os/signal"

	"github.com/gordonklaus/portaudio"
	"github.com/spf13/cobra"
)

var (
	averageSamples int
)

// DumpInputCmd represents the DumpInputs command
var DumpInputCmd = &cobra.Command{
	Use:   "dump-input",
	Short: "Dumps the audio input of a device to console",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := portaudio.Initialize(); err != nil {
			panic(err)
		}
		defer portaudio.Terminate()

		buf := make([]float32, averageSamples)
		s, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buf), buf)
		if err != nil {
			panic(err)
		}
		defer s.Close()

		if err := s.Start(); err != nil {
			panic(err)
		}
		defer s.Stop()

		var frame int64
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

		fmt.Println("Started listener, dumping input.")

		for {
			if err := s.Read(); err != nil {
				panic(err)
			}

			calcAverage(frame, buf)
			frame++

			select {
			case <-c:
				fmt.Println("cancelled.")
				return
			default:
			}
		}
	},
}

func init() {
	AudioCmd.AddCommand(DumpInputCmd)

	AudioCmd.PersistentFlags().IntVarP(&averageSamples, "average-samples", "a", 1000, "How many samples to calc the average from")
}

func calcAverage(frame int64, buf []float32) {
	var avg, min, max, sum float32
	for _, s := range buf {
		sum += s

		if min == 0 || s < min {
			min = s
		} else if min == 0 || s > max {
			max = s
		}
	}

	avg = sum / float32(len(buf))

	fmt.Printf(
		"%10d %14s %14s %14s\n",
		frame,
		fmt.Sprintf("%5.10f", avg),
		fmt.Sprintf("%5.10f", min),
		fmt.Sprintf("%5.10f", max),
	)
}
