// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/apinnecke/dmx-auto-control/metronome"
	"github.com/apinnecke/dmx-auto-control/metronome/output"
	"github.com/apinnecke/dmx-auto-control/metronome/utils"
	"github.com/spf13/cobra"
)

const (
	outputTypeAudio  = "audio"
	outputTypeStdOut = "stdout"
)

var (
	outputType string
	strongFreq float64
	weakFreq   float64
	limit      uint

	outputTypes = []string{outputTypeStdOut, outputTypeAudio}
)

// metronomeCmd represents the metronome command
var metronomeCmd = &cobra.Command{
	Use:     "metronome [speed beats noteValue]",
	Example: "metronome 160 3 4",
	Short:   "Simple metronome with cli and audio output",
	Long:    `A very simple but flexible metronome using time.Ticker and channels for communication, mainly used for testing some stuff.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			panic(fmt.Errorf("Need to pass 3 args, got %d", len(args)))
		}

		var speed, beats, noteValue uint64
		var err error

		if speed, err = strconv.ParseUint(args[0], 10, 64); err != nil {
			panic(fmt.Errorf("Unable to parse %q as speed", args[0]))
		}
		if beats, err = strconv.ParseUint(args[1], 10, 64); err != nil {
			panic(fmt.Errorf("Unable to parse %q as beats", args[1]))
		}
		if noteValue, err = strconv.ParseUint(args[2], 10, 64); err != nil {
			panic(fmt.Errorf("Unable to parse %q as noteValue", args[2]))
		}

		sig := utils.GetSignal()
		var out metronome.Output

		switch outputType {
		case outputTypeAudio:
			o := output.NewAudioOutput(strongFreq, weakFreq)
			if err := o.Start(); err != nil {
				panic(err)
			}

			defer o.Stop()
			out = o
			break
		case outputTypeStdOut:
			out = output.NewBufferOutput(os.Stdout)
			break
		default:
			panic(fmt.Errorf("Invalid output type %q, valid are %v", outputType, outputTypes))
		}

		m := metronome.NewPlayer(out)
		b := metronome.NewBar(uint(beats), uint(noteValue), uint(speed))

		if err := m.PlayBarUntilSignalOrLimit(b, sig, limit); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(metronomeCmd)

	metronomeCmd.Flags().StringVarP(&outputType, "output", "o", "audio", fmt.Sprintf("Which output should be used %v", outputTypes))
	metronomeCmd.Flags().Float64Var(&strongFreq, "strongFreq", 1760, "Which frequency should be used to render the sin wave for the strong bar accent click (audio only)")
	metronomeCmd.Flags().Float64Var(&weakFreq, "weakFreq", 1320, "Which frequency should be used to render the sin wave for the weak mediate click (audio only)")
	metronomeCmd.Flags().UintVar(&limit, "limit", 0, "How many clicks to play. Used to limit to a single bar for example")
}
