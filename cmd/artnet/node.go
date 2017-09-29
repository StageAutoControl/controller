// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package artnet

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"

	root "github.com/StageAutoControl/controller/cmd"
	artnetTransport "github.com/StageAutoControl/controller/cntl/transport/artnet"
	"github.com/jsimonetti/go-artnet"
	"github.com/jsimonetti/go-artnet/packet/code"
	"github.com/spf13/cobra"
)

// Node represents the ArtNetTest command
var Node = &cobra.Command{
	Use:   "node",
	Short: "ArtNet node to test network communication",
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

		n := artnet.NewNode(fmt.Sprintf("node-%d", rand.Int()), code.StNode, ip, artnet.NewLogger(root.Logger))

		if err := n.Start(); err != nil {
			log.Fatal(err)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c

		log.Println("Stopping node ...")
		n.Stop()
	},
}

func init() {
	ArtNetCmd.AddCommand(Node)
}
