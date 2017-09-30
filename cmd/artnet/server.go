// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package artnet

import (
	"net"
	"runtime"
	"sync"
	"time"

	root "github.com/StageAutoControl/controller/cmd"
	artnetTransport "github.com/StageAutoControl/controller/cntl/transport/artnet"
	"github.com/jsimonetti/go-artnet"
	"github.com/spf13/cobra"
)

// Server represents the ArtNetTest command
var Server = &cobra.Command{
	Use:   "server",
	Short: "ArtNet server to test network communication",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var ip net.IP
		var err error

		root.Logger.Info("InterfaceName is empty, searching for suitable one ...")
		ip, err = artnetTransport.FindArtNetIP()
		if err != nil {
			root.Logger.Fatal(err)
		}

		root.Logger.Infof("Using interface with IP %s", ip.String())

		if len(ip) == 0 {
			root.Logger.Fatal("No IP found")
		}

		c := artnet.NewController("controller-1", ip, artnet.NewLogger(root.Logger))
		var wg sync.WaitGroup

		go func() {
			wg.Add(1)
			if err := c.Start(); err != nil {
				root.Logger.Fatal(err)
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

		root.Logger.Infof("num: %d", runtime.NumGoroutine())
	},
}

func init() {
	ArtNetCmd.AddCommand(Server)
}
