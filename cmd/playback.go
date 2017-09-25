// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

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
		transport.TypeBuffer,
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
	audioWaiterThreshold int32
)

// playbackCmd represents the playback command
var playbackCmd = &cobra.Command{
	Use:   "playback [song|setlist] song-valid-uuid-1",
	Short: "Plays a given Song or SetList by id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Usage()
			os.Exit(1)
		}

		var loader cntl.Loader
		switch loaderType {
		case directoryLoader:
			fmt.Printf("Loading data directoy %q ... \n", dataDir)
			loader = files.New(dataDir)
		case databaseLoader:
			//loader = database.New(),
			fmt.Println("Database loader is not yet supported.")
			os.Exit(1)

		default:
			fmt.Printf("Loader %q is not supported. Choose one of %s \n", loader, loaders)
			os.Exit(1)
		}

		data, err := loader.Load()
		if err != nil {
			fmt.Printf("Failed to load data from %q: %v \n", loaderType, err)
			os.Exit(1)
		}

		var writers []playback.TransportWriter
		for _, transportType := range usedTransports {
			switch transportType {
			case transport.TypeBuffer:
				writers = append(writers, transport.NewBuffer(os.Stdout))
				break

			case transport.TypeVisualizer:
				w, err := transport.NewVisualizer(viualizerEndpoint)
				if err != nil {
					fmt.Printf("Unable to connect to the visualizer: %v \n", err)
					os.Exit(1)
				}

				writers = append(writers, w)
				break

			case transport.TypeArtNet:
				w, err := transport.NewArtNet("stage-auto-control")
				if err != nil {
					fmt.Printf("Unable to connect to the visualizer: %v \n", err)
					os.Exit(1)
				}

				writers = append(writers, w)
				break

			case transport.TypeMidi:
				w, err := transport.NewMIDI(midiDeviceID)
				if err != nil {
					fmt.Printf("Unable to connect to midi device: %v \n", err)
					os.Exit(1)
				}

				writers = append(writers, w)
				break

			default:
				fmt.Printf("Transport %q is not supported. \n", transportType)
				os.Exit(1)

				break
			}
		}

		var waiters []playback.Waiter
		for _, waiterType := range usedWaiters {
			switch waiterType {
			case waiter.TypeNone:
				waiters = append(waiters, waiter.NewNone())

				break

			case waiter.TypeAudio:
				a, err := waiter.NewAudio(audioWaiterThreshold)
				if err != nil {
					panic(err)
				}

				waiters = append(waiters, a)

				break
			}
		}

		player := playback.NewPlayer(data, writers, waiters)

		switch args[0] {
		case playbackTypeSong:
			songID := args[1]
			if err = player.PlaySong(songID); err != nil {
				panic(err)
			}

			break
		case playbackTypeSetList:
			setListID := args[2]
			if err = player.PlaySetList(setListID); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(playbackCmd)

	playbackCmd.PersistentFlags().StringSliceVarP(&usedTransports, "transport", "t", []string{transport.TypeBuffer}, fmt.Sprintf("Which usedTransports to use from %s.", transportTypes))
	playbackCmd.PersistentFlags().StringVar(&viualizerEndpoint, "visualizer-endpoint", "localhost:1337", "Endpoint of the visualizer backend if visualizer transport is chosen.")
	playbackCmd.PersistentFlags().StringVarP(&midiDeviceID, "midi-device-id", "m", "", "DeviceID of MIDI output to use (On empty string the default device is used)")

	playbackCmd.PersistentFlags().StringSliceVarP(&usedWaiters, "wait-for", "w", []string{waiter.TypeNone}, fmt.Sprintf("Wait for a specific signal before playing a song (required to be used on stage, otherwise the next song would start immediately), one of %s", waiterTypes))
	playbackCmd.PersistentFlags().Int32Var(&audioWaiterThreshold, "audio-waiter-threshold", 15000, "Threshold frequency for audio waiter to trigger a signal")
}
