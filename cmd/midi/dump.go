// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package midi

import (
	"fmt"
	"os"

	"github.com/rakyll/portmidi"
	"github.com/spf13/cobra"
)

// MidiDumpCmd represents the MidiDevices command
var MidiDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump incoming midi commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := portmidi.Initialize(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		defer portmidi.Terminate()

		deviceID := portmidi.DefaultInputDeviceID()
		if deviceID == 0 {
			fmt.Fprintln(os.Stderr, "Unable to find default input interface")
			os.Exit(1)
		}

		in, err := portmidi.NewInputStream(deviceID, 1024)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		// or alternatively listen events
		events := in.Listen()
		for event := range events {
			fmt.Printf("%x\t%d\t%d\n", event.Status, event.Data1, event.Data2)
		}
	},
}

func init() {
	MidiCmd.AddCommand(MidiDumpCmd)
}
