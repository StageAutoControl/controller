// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/cntl/output"
	"github.com/StageAutoControl/controller/cntl/song"
	"github.com/StageAutoControl/controller/database/files"
	"github.com/spf13/cobra"
)

const (
	directoryLoader = "directory"
	databaseLoader  = "database"
)

var (
	loaders    = []string{directoryLoader, databaseLoader}
	loaderType string
	dataDir    string
	songID     string
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

		s, ok := data.Songs[songID]
		if !ok {
			fmt.Printf("Unable to find song %q.\n", s)
			os.Exit(1)
		}

		output := output.NewBufferOutput(os.Stdout)
		player := song.NewPlayer(data, output)

		err = player.PlayAll(songID)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(playbackCmd)

	playbackCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "d", "", "Data directory to load (when loader is set to directory)")
	playbackCmd.PersistentFlags().StringVarP(&loaderType, "loader", "l", directoryLoader, fmt.Sprintf("Which loader to use %s.", loaders))
}
