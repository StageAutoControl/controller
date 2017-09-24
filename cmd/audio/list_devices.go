// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package audio

import (
	"fmt"
	"os"

	"github.com/gordonklaus/portaudio"
	"github.com/spf13/cobra"
)

// DeviceCmd represents the Devices command
var DeviceCmd = &cobra.Command{
	Use:   "devices",
	Short: "Prints info about all devices",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := portaudio.Initialize(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer portaudio.Terminate()

		devices, err := portaudio.Devices()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Found %d devices. \n", len(devices))
		if len(devices) == 0 {
			return
		}

		fmt.Println("\n\nID Name Input Output SampleRate")

		for i, device := range devices {
			fmt.Printf("%v ", i)
			printDevice(device)
		}

		fmt.Println("\n\nDefault input device: ")
		defaultInputDevice, err := portaudio.DefaultInputDevice()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printDevice(defaultInputDevice)

		fmt.Println("\n\nDefault output device: ")
		defaultOutputDevice, err := portaudio.DefaultOutputDevice()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printDevice(defaultOutputDevice)
	},
}

func printDevice(info *portaudio.DeviceInfo) {
	fmt.Printf("%v %v %v %v\n", info.Name, info.MaxInputChannels, info.MaxOutputChannels, info.DefaultSampleRate)
}

func init() {
	AudioCmd.AddCommand(DeviceCmd)
}
