// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/apinnecke/go-exitcontext"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/cntl/playback"
	"github.com/StageAutoControl/controller/pkg/cntl/transport"
	"github.com/StageAutoControl/controller/pkg/cntl/waiter"
	"github.com/StageAutoControl/controller/pkg/visualizer"
)

const (
	playbackTypeSong    = "song"
	playbackTypeSetList = "setlist"
)

var (
	transportTypes = []string{
		transport.TypeStream,
		transport.TypeVisualizer,
		transport.TypeArtNet,
		transport.TypeMidi,
	}
	usedTransports []string

	viualizerEndpoint string
	midiDeviceID      int8

	waiterTypes = []string{
		waiter.TypeNone,
		waiter.TypeAudio,
	}
	usedWaiters          []string
	audioWaiterThreshold float32
)

// playbackCmd represents the playback command
var playbackCmd = &cobra.Command{
	Use:   "playback [song|setlist] song-valid-uuid-1",
	Short: "Plays a given Process or SetList by id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetLevel(logrus.WarnLevel)

		if len(args) != 2 {
			if err := cmd.Usage(); err != nil {
				logrus.Fatal(err)
			}
			os.Exit(1)
		}

		data, err := loader.Load()
		if err != nil {
			logrus.Fatal(err)
		}

		var writers []playback.TransportWriter
		for _, transportType := range usedTransports {
			switch transportType {

			case transport.TypeStream:
				writers = append(writers, transport.NewStream(logger.WithField(cntl.LoggerFieldTransport, transport.TypeStream), os.Stdout))

			case transport.TypeBarLogger:
				writers = append(writers, transport.NewBarLogger(logger.WithField(cntl.LoggerFieldTransport, transport.TypeBarLogger)))

			case transport.TypeVisualizer:
				w := visualizer.NewServer(logger.WithField(cntl.LoggerFieldTransport, transport.TypeVisualizer))
				writers = append(writers, w)

			case transport.TypeArtNet:
				controller, err := artnet.NewController(logger.WithField(cntl.LoggerFieldTransport, transport.TypeArtNet))
				if err != nil {
					logger.Fatal(err)
				}

				w, err := transport.NewArtNet(controller)
				if err != nil {
					logger.Fatalf("Unable to open art net controller: %v", err)
				}

				writers = append(writers, w)

			case transport.TypeMidi:
				w, err := transport.NewMIDI(logger.WithField(cntl.LoggerFieldTransport, transport.TypeMidi), midiDeviceID)
				if err != nil {
					logger.Fatalf("Unable to connect to midi device: %v", err)
				}

				writers = append(writers, w)

			default:
				logger.Fatalf("Transport %q is not supported", transportType)
			}
		}

		var waiters []playback.Waiter
		for _, waiterType := range usedWaiters {
			switch waiterType {
			case waiter.TypeNone:
				waiters = append(waiters, waiter.NewNone(logger.WithField(cntl.LoggerFieldWaiter, waiter.TypeNone)))

			case waiter.TypeAudio:
				waiters = append(waiters, waiter.NewAudio(logger.WithField(cntl.LoggerFieldWaiter, waiter.TypeAudio), audioWaiterThreshold))

			}
		}

		ctx := exitcontext.New()
		player := playback.NewPlayer(logger.Logger.WithField("player", "default"), data, writers, waiters)

		switch args[0] {
		case playbackTypeSong:
			songID := args[1]
			if err = player.PlaySong(ctx, songID); err != nil {
				logger.Fatal(err)
			}

		case playbackTypeSetList:
			setListID := args[1]
			if err = player.PlaySetList(ctx, setListID); err != nil {
				logger.Fatal(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(playbackCmd)

	playbackCmd.Flags().StringSliceVarP(&usedTransports, "transport", "t", []string{}, fmt.Sprintf("Which usedTransports to use from %s.", transportTypes))
	playbackCmd.Flags().StringVar(&viualizerEndpoint, "visualizer-endpoint", "localhost:1337", "Endpoint of the visualizer backend if visualizer transport is chosen.")
	playbackCmd.Flags().Int8VarP(&midiDeviceID, "midi-device-id", "m", -1, "DeviceID of MIDI output to use (On empty string the default device is used)")
	playbackCmd.Flags().StringSliceVarP(&usedWaiters, "wait-for", "w", []string{waiter.TypeNone}, fmt.Sprintf("Wait for a specific signal before playing a song (required to be used on stage, otherwise the next song would start immediately), one of %s", waiterTypes))
	playbackCmd.Flags().Float32Var(&audioWaiterThreshold, "audio-waiter-threshold", 0.9, "Threshold frequency for audio waiter to trigger a signal")
}
