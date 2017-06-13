// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
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
)

var (
	transports        = []string{bufferTransport, visualizerTransport, artnetTransport}
	transportTypes    []string
	viualizerEndpoint string
	songID            string
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
			panic(fmt.Errorf("Failed to load %q data: %v", loaderType, err))
		}

		fmt.Printf("Loaded %d set lists, %d songs, %d scenes, %d presets %d animations, %d device types, %d device groups and %d devices\n",
			len(data.SetLists), len(data.Songs), len(data.DMXScenes), len(data.DMXPresets), len(data.DMXAnimations),
			len(data.DMXDeviceTypes), len(data.DMXDeviceGroups), len(data.DMXDevices))

		_, ok := data.Songs[songID]
		if !ok {
			fmt.Printf("Unable to find song %q.\n", songID)
			os.Exit(1)
		}

		var writers []song.TransportWriter
		for _, transportType := range transportTypes {
			switch transportType {
			case bufferTransport:
				writers = append(writers, transport.NewBufferTransport(os.Stdout))
				break

			case visualizerTransport:
				w, err := transport.NewVisualizerTransport(viualizerEndpoint)
				if err != nil {
					fmt.Printf("Unable to connect to the visualizer: %v \n", err)
					os.Exit(1)
				}

				writers = append(writers, w)

			case artnetTransport:
				w, err := transport.NewArtNet("stage-auto-control")
				if err != nil {
					fmt.Printf("Unable to connect to the visualizer: %v \n", err)
					os.Exit(1)
				}

				writers = append(writers, w)

			default:
				fmt.Printf("Transport %q is not supported. \n", transportTypes)
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
}
