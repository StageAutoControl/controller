// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package midi

import (
	"fmt"
	"log"
	"os"

	"github.com/rakyll/portmidi"
	"github.com/spf13/cobra"
)

// MidiDeviceCmd represents the MidiDevices command
var MidiDeviceCmd = &cobra.Command{
	Use:   "devices",
	Short: "Prints info about all devices",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := portmidi.Initialize(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func() {
			if err := portmidi.Terminate(); err != nil {
				log.Fatal(err)
			}
		}()

		num := portmidi.CountDevices()
		fmt.Printf("Found %d devices. \n", num)
		if num == 0 {
			return
		}

		for i := 0; i < num; i++ {
			info := portmidi.Info(portmidi.DeviceID(i))
			if info == nil {
				fmt.Println("Unable to read default output devices")
				os.Exit(1)
			}
			printDevice(portmidi.DeviceID(i), info)
		}

		fmt.Println("Default device: ")
		deviceID := portmidi.DefaultOutputDeviceID()
		info := portmidi.Info(deviceID)
		if info == nil {
			fmt.Println("Unable to read default output devices")
			os.Exit(1)
		}
		printDevice(deviceID, info)
	},
}

func init() {
	MidiCmd.AddCommand(MidiDeviceCmd)
}

func printDevice(deviceID portmidi.DeviceID, info *portmidi.DeviceInfo) {
	fmt.Printf("%d: %+v \n", deviceID, info)
}
