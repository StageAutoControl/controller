// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package midi

import (
	"github.com/StageAutoControl/controller/cmd"
	"github.com/spf13/cobra"
)

// MidiCmd represents the midi/midi command
var MidiCmd = &cobra.Command{
	Use:   "midi",
	Short: "A brief description of your command",
	Long:  ``,
}

func init() {
	cmd.RootCmd.AddCommand(MidiCmd)
}
