// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package artnet

import (
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	root "github.com/StageAutoControl/controller/cmd"
	artnetTransport "github.com/StageAutoControl/controller/cntl/transport/artnet"
	"github.com/jsimonetti/go-artnet"
	"github.com/spf13/cobra"
)

// Listen represents the ArtNetTest command
var Listen = &cobra.Command{
	Use:   "listen",
	Short: "ArtNet server to listen for devices and print them",
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

		host, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		c := artnet.NewController(host, ip, artnet.NewLogger(root.Logger))
		var wg sync.WaitGroup

		go func() {
			wg.Add(1)
			if err := c.Start(); err != nil {
				root.Logger.Fatal(err)
			}

			wg.Done()
		}()

		time.Sleep(10 * time.Second)

		cancel := make(chan os.Signal, 2)
		signal.Notify(cancel, syscall.SIGTERM, syscall.SIGKILL)
		var builder strings.Builder

	LOOP:
		for {
			select {
			case <-cancel:
				break LOOP
			default:
			}

			for _, n := range c.Nodes {
				builder.WriteString(artnetTransport.NodeToString(n))
			}

			root.Logger.Infof(builder.String())
			builder.Reset()

			time.Sleep(10 * time.Second)
		}

		c.Stop()
		wg.Wait()

		root.Logger.Infof("num: %d", runtime.NumGoroutine())
	},
}

func init() {
	ArtNetCmd.AddCommand(Listen)
}
