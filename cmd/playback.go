// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
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
)

// playbackCmd represents the playback command
var playbackCmd = &cobra.Command{
	Use:   "playback",
	Short: "Plays a given songname",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("playback called")
		var loader cntl.Loader

		switch loaderType {
		case directoryLoader:
			loader = files.New(dataDir)
		case databaseLoader:
			//loader = database.New(),
		default:
			panic(fmt.Errorf("Loader %q is not supported. Choose one of %s", loader, loaders))
		}
	},
}

func init() {
	RootCmd.AddCommand(playbackCmd)

	playbackCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "d", "", "Data directory to load (when loader is set to directory)")
	playbackCmd.PersistentFlags().StringVarP(&loaderType, "loader", "l", directoryLoader, fmt.Sprintf("Which loader to use %s.", loaders))
}
