// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package audio

import (
	"github.com/StageAutoControl/controller/cmd"
	"github.com/spf13/cobra"
)

// AudioCmd represents the audio/audio command
var AudioCmd = &cobra.Command{
	Use:   "audio",
	Short: "A brief description of your command",
	Long:  ``,
}

func init() {
	cmd.RootCmd.AddCommand(AudioCmd)
}
