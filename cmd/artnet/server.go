// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package artnet

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"sync"
	"time"

	"github.com/StageAutoControl/controller/cmd"
	artnetTransport "github.com/StageAutoControl/controller/cntl/transport/artnet"
	"github.com/jsimonetti/go-artnet"
	"github.com/spf13/cobra"
)

// ArtNetServer represents the ArtNetTest command
var ArtNetServer = &cobra.Command{
	Use:   "artnet-server",
	Short: "ArtNet server to test network communication",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var ip net.IP
		var err error

		log.Println("InterfaceName is empty, searching for suitable one ...")
		ip, err = artnetTransport.FindArtNetIP()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Using interface with IP %s", ip.String())

		if len(ip) == 0 {
			log.Fatal("No IP found")
		}

		c := artnet.NewController("controller-1", ip)
		var wg sync.WaitGroup

		go func() {
			wg.Add(1)
			if err := c.Start(); err != nil {
				log.Fatal(err)
			}

			wg.Done()
		}()

		time.Sleep(10 * time.Second)
		c.SendDMXToAddress([512]byte{0x00, 0xff, 0x00, 0xff, 0x00}, artnet.Address{Net: 0, SubUni: 0})
		time.Sleep(2 * time.Second)
		c.SendDMXToAddress([512]byte{0xff, 0x00, 0x00, 0xff, 0x00}, artnet.Address{Net: 0, SubUni: 0})
		time.Sleep(2 * time.Second)
		c.SendDMXToAddress([512]byte{0x00, 0x00, 0xff, 0xff, 0x00}, artnet.Address{Net: 0, SubUni: 0})
		time.Sleep(2 * time.Second)
		c.SendDMXToAddress([512]byte{}, artnet.Address{Net: 0, SubUni: 0})
		time.Sleep(2 * time.Second)

		c.Stop()
		wg.Wait()

		fmt.Printf("num: %d", runtime.NumGoroutine())
	},
}

func init() {
	cmd.RootCmd.AddCommand(ArtNetServer)
}
