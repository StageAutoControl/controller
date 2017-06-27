// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package midi

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rakyll/portmidi"
	"github.com/spf13/cobra"
)

var (
	deviceID string
)

// MidiDumpCmd represents the MidiDevices command
var MidiDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump incoming midi commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := portmidi.Initialize(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer portmidi.Terminate()

		var d portmidi.DeviceID
		if deviceID == "" {
			d = portmidi.DefaultOutputDeviceID()
		} else {
			i, err := strconv.Atoi(deviceID)
			if err != nil {
				fmt.Printf("Failed to transform deviceID %q to int: %v\n", deviceID, err)
			}
			d = portmidi.DeviceID(i)
		}

		i := portmidi.Info(d)
		if i == nil {
			fmt.Printf("Unable to find input interface %d \n", d)
			os.Exit(1)
		}

		fmt.Printf("Using midi device %d \n", d)

		in, err := portmidi.NewInputStream(d, 1024)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Starting dump for midi interface %d ... \n", d)

		// or alternatively listen events
		events := in.Listen()
		for event := range events {
			fmt.Printf("%x\t%d\t%d\n", event.Status, event.Data1, event.Data2)
		}
	},
}

func init() {
	MidiCmd.AddCommand(MidiDumpCmd)

	MidiDumpCmd.PersistentFlags().StringVar(&deviceID, "device-id", "", "DeviceID of MIDI output to use (On empty string the default device is used)")
}
