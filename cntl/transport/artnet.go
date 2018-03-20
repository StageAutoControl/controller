package transport

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/StageAutoControl/controller/cntl"
	artnetTransport "github.com/StageAutoControl/controller/cntl/transport/artnet"
	artnet "github.com/jsimonetti/go-artnet"
)

// ArtNet is a transport for the ArtNet protocol (DMX over UDP/IP)
type ArtNet struct {
	name string
	c    *artnet.Controller
	state artnetTransport.State
}

// NewArtNet returns a new ArtNet transport instance
func NewArtNet(logger *logrus.Entry, name string) (*ArtNet, error) {
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

	logger.Warn("Waiting 5 seconds for nodes to register")
	time.Sleep(5 * time.Second)

	return &ArtNet{
		name: name,
		c:    c,
		state: artnetTransport.NewState(),
	}, nil
}

func (a *ArtNet) Write(cmd cntl.Command) error {
	for _, c := range cmd.DMXCommands {
		a.state.Set(uint16(c.Universe), uint8(c.Channel), c.Value.Uint8())
	}

	for u, dmx := range a.state {
		a.c.SendDMXToAddress(dmx, universeToAddress(cntl.DMXUniverse(u)))
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
