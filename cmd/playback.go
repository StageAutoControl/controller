// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/cntl/song"
	"github.com/StageAutoControl/controller/cntl/transport"
	"github.com/StageAutoControl/controller/database/files"
	"github.com/spf13/cobra"
)

const (
	bufferTransport     = "buffer"
	visualizerTransport = "visualizer"
	artnetTransport     = "artnet"
	midiTransport       = "midi"
)

var (
	transports        = []string{bufferTransport, visualizerTransport, artnetTransport, midiTransport}
	transportTypes    []string
	viualizerEndpoint string
	songID            string
	midiDeviceID      string
)

// playbackCmd represents the playback command
var playbackCmd = &cobra.Command{
	Use:   "playback song-uuid",
	Short: "Plays a given songname",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		songID = args[0]

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
			panic(fmt.Errorf("Loader %q is not supported. Choose one of %s", loader, loaders))
		}

		data, err := loader.Load()
		if err != nil {
			panic(fmt.Errorf("Failed to load data from %q: %v", loaderType, err))
		}

		fmt.Printf("Loaded %d set lists, %d songs, %d scenes, %d presets %d animations, %d device types, %d device groups and %d devices\n",
			len(data.SetLists), len(data.Songs), len(data.DMXScenes), len(data.DMXPresets), len(data.DMXAnimations),
			len(data.DMXDeviceTypes), len(data.DMXDeviceGroups), len(data.DMXDevices))

		s, ok := data.Songs[songID]
		if !ok {
			fmt.Printf("Unable to find song %q.\n", songID)
			os.Exit(1)
		}

		log.Printf("Playing song %q (%s) ...", s.Name, songID)

		var writers []song.TransportWriter
		for _, transportType := range transportTypes {
			switch transportType {
			case bufferTransport:
				writers = append(writers, transport.NewBuffer(os.Stdout))
				break

			case visualizerTransport:
				w, err := transport.NewVisualizer(viualizerEndpoint)
				if err != nil {
					fmt.Printf("Unable to connect to the visualizer: %v \n", err)
					os.Exit(1)
				}

				writers = append(writers, w)
				break

			case artnetTransport:
				w, err := transport.NewArtNet("stage-auto-control")
				if err != nil {
					fmt.Printf("Unable to connect to the visualizer: %v \n", err)
					os.Exit(1)
				}

				writers = append(writers, w)
				break

			case midiTransport:
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
			}
		}
		player := song.NewPlayer(data, writers)
		if err = player.PlayAll(songID); err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(playbackCmd)

	playbackCmd.PersistentFlags().StringSliceVarP(&transportTypes, "transport", "t", []string{bufferTransport}, fmt.Sprintf("Which transports to use from %s.", transports))
	playbackCmd.PersistentFlags().StringVar(&viualizerEndpoint, "visualizer-endpoint", "localhost:1337", "Endpoint of the visualizer backend if visualizer transport is chosen.")
	playbackCmd.PersistentFlags().StringVarP(&midiDeviceID, "midi-device-id", "m", "", "DeviceID of MIDI output to use (On empty string the default device is used)")
}
