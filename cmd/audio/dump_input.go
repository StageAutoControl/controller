// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package audio

import (
	"fmt"
	"log"
	"os"

	"os/signal"

	"github.com/gordonklaus/portaudio"
	"github.com/spf13/cobra"
)

var (
	averageSamples int
	threshold      float32
)

// DumpInputCmd represents the DumpInputs command
var DumpInputCmd = &cobra.Command{
	Use:   "dump-input",
	Short: "Dumps the audio input of a device to console",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		buf := make([]float32, averageSamples)
		s, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buf), buf)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := s.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		if err := s.Start(); err != nil {
			panic(err)
		}
		defer func() {
			if err := s.Stop(); err != nil {
				log.Fatal(err)
			}
		}()

		var frame int64
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Interrupt)

		fmt.Println("Started listener, dumping input.")

		for {
			if err := s.Read(); err != nil {
				panic(err)
			}

			go calcAverage(frame, buf)
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

	DumpInputCmd.PersistentFlags().IntVarP(&averageSamples, "average-samples", "a", 1000, "How many samples to calc the average from")
	DumpInputCmd.PersistentFlags().Float32VarP(&threshold, "threshold", "t", 0.8, "Threshold of tick")

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

	tick := ""
	if min < (threshold*-1) || max > threshold {
		tick = "tick"
	}

	avg = sum / float32(len(buf))

	go fmt.Printf(
		"%10d %14s %14s %14s %10s\n",
		frame,
		fmt.Sprintf("%5.10f", avg),
		fmt.Sprintf("%5.10f", min),
		fmt.Sprintf("%5.10f", max),
		tick,
	)
}
