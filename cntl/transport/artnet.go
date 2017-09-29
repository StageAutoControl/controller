package transport

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/StageAutoControl/controller/cntl"
	artnetTransport "github.com/StageAutoControl/controller/cntl/transport/artnet"
	artnet "github.com/jsimonetti/go-artnet"
)

// ArtNet is a transport for the ArtNet protocol (DMX over UDP/IP)
type ArtNet struct {
	name string
	c    *artnet.Controller
}

// NewArtNet returns a new ArtNet transport instance
func NewArtNet(name string, logger *logrus.Entry) (*ArtNet, error) {
	ip, err := artnetTransport.FindArtNetIP()
	if err != nil {
		return nil, fmt.Errorf("failed to find the art-net IP: %v", err)
	}

	if len(ip) == 0 {
		return nil, errors.New("failed to find the art-net IP: No interface found")
	}

	c := artnet.NewController(name, ip, artnet.NewLogger(logger))
	if err := c.Start(); err != nil {
		return nil, fmt.Errorf("failed to start controller: %v", err)
	}

	return &ArtNet{
		name: name,
		c:    c,
	}, nil
}

func (a *ArtNet) Write(cmd cntl.Command) error {
	for _, c := range cmd.DMXCommands {
		dmx := [512]byte{}
		dmx[c.Channel] = c.Value.Byte()
		a.c.SendDMXToAddress(dmx, universeToAddress(c.Universe))
	}

	return nil
}

func universeToAddress(u cntl.DMXUniverse) artnet.Address {
	// https://play.golang.org/p/pdQPC5u7JX

	v := make([]uint8, 2)
	binary.BigEndian.PutUint16(v, uint16(u))

	return artnet.Address{
		Net:    v[0],
		SubUni: v[1],
	}
}
