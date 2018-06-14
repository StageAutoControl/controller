// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/StageAutoControl/controller/cmd/internal"
	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/cntl/playback"
	"github.com/StageAutoControl/controller/cntl/transport"
	"github.com/StageAutoControl/controller/cntl/waiter"
	"github.com/StageAutoControl/controller/database/files"
	"github.com/spf13/cobra"
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
	usedTransports    []string
	viualizerEndpoint string
	midiDeviceID      string

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
	Short: "Plays a given Song or SetList by id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetLevel(logrus.WarnLevel)

		if len(args) != 2 {
			cmd.Usage()
			os.Exit(1)
		}

		var loader cntl.Loader
		switch loaderType {
		case directoryLoader:
			Logger.Infof("Loading data directory %q ...", dataDir)
			loader = files.New(dataDir)

		case databaseLoader:
			//loader = database.New(),
			Logger.Fatal("Database loader is not yet supported.")

		default:
			Logger.Fatalf("Loader %q is not supported. Choose one of %s", loader, loaders)
		}

		data, err := loader.Load()
		if err != nil {
			Logger.Fatalf("Failed to load data from %q: %v", loaderType, err)
		}

		var writers []playback.TransportWriter
		for _, transportType := range usedTransports {
			switch transportType {

			case transport.TypeStream:
				writers = append(writers, transport.NewStream(Logger.WithField(cntl.LoggerFieldTransport, transport.TypeStream), os.Stdout))
				break

			case transport.TypeBarLogger:
				writers = append(writers, transport.NewBarLogger(Logger.WithField(cntl.LoggerFieldTransport, transport.TypeBarLogger)))
				break

			case transport.TypeVisualizer:
				w, err := transport.NewVisualizer(Logger.WithField(cntl.LoggerFieldTransport, transport.TypeVisualizer), viualizerEndpoint)
				if err != nil {
					Logger.Fatalf("Unable to connect to the visualizer: %v", err)
				}

				writers = append(writers, w)
				break

			case transport.TypeArtNet:
				w, err := transport.NewArtNet(Logger.WithField(cntl.LoggerFieldTransport, transport.TypeArtNet), "stage-auto-control")
				if err != nil {
					Logger.Fatalf("Unable to connect to the visualizer: %v", err)
				}

				writers = append(writers, w)
				break

			case transport.TypeMidi:
				w, err := transport.NewMIDI(Logger.WithField(cntl.LoggerFieldTransport, transport.TypeMidi), midiDeviceID)
				if err != nil {
					Logger.Fatalf("Unable to connect to midi device: %v", err)
				}

				writers = append(writers, w)
				break

			default:
				Logger.Fatalf("Transport %q is not supported.", transportType)
			}
		}

		var waiters []playback.Waiter
		for _, waiterType := range usedWaiters {
			switch waiterType {
			case waiter.TypeNone:
				waiters = append(waiters, waiter.NewNone(Logger.WithField(cntl.LoggerFieldWaiter, waiter.TypeNone)))

				break

			case waiter.TypeAudio:
				a, err := waiter.NewAudio(Logger.WithField(cntl.LoggerFieldWaiter, waiter.TypeAudio), audioWaiterThreshold)
				if err != nil {
					Logger.Fatal(err)
				}

				waiters = append(waiters, a)

				break
			}
		}

		ctx := internal.NewExitHandlerContext(Logger.Logger)
		player := playback.NewPlayer(Logger.Logger.WithField("player", "default"), data, writers, waiters)

		switch args[0] {
		case playbackTypeSong:
			songID := args[1]
			if err = player.PlaySong(ctx, songID); err != nil {
				Logger.Fatal(err)
			}

			break
		case playbackTypeSetList:
			setListID := args[1]
			if err = player.PlaySetList(ctx, setListID); err != nil {
				Logger.Fatal(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(playbackCmd)

	playbackCmd.PersistentFlags().StringSliceVarP(&usedTransports, "transport", "t", []string{}, fmt.Sprintf("Which usedTransports to use from %s.", transportTypes))
	playbackCmd.PersistentFlags().StringVar(&viualizerEndpoint, "visualizer-endpoint", "localhost:1337", "Endpoint of the visualizer backend if visualizer transport is chosen.")
	playbackCmd.PersistentFlags().StringVarP(&midiDeviceID, "midi-device-id", "m", "", "DeviceID of MIDI output to use (On empty string the default device is used)")

	playbackCmd.PersistentFlags().StringSliceVarP(&usedWaiters, "wait-for", "w", []string{waiter.TypeNone}, fmt.Sprintf("Wait for a specific signal before playing a song (required to be used on stage, otherwise the next song would start immediately), one of %s", waiterTypes))
	playbackCmd.PersistentFlags().Float32Var(&audioWaiterThreshold, "audio-waiter-threshold", 0.9, "Threshold frequency for audio waiter to trigger a signal")
}
