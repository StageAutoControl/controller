package artnet

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jsimonetti/go-artnet"
	"github.com/sirupsen/logrus"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

// Controller is a transport for the ArtNet protocol (DMX over UDP/IP)
type controller struct {
	logger      logging.Logger
	sender      *artnet.Controller
	state       *State
	sendTrigger chan UniverseStateMap
	context     context.Context
}

// NewController returns a artnet Controller as an anonymous interface
func NewController(logger logging.Logger) (Controller, error) {
	ip, err := FindArtNetIP()
	if err != nil {
		return nil, fmt.Errorf("failed to find the art-net IP: %v", err)
	}

	if len(ip) == 0 {
		return nil, errors.New("failed to find the art-net IP: No interface found")
	}

	host, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve hostname: %v", err)
	}

	host = strings.ToLower(strings.Split(host, ".")[0])
	logger.Infof("Using ArtNet IP %s and hostname %s", ip.String(), host)

	senderLogger := artnet.NewLogger(logger.(*logrus.Entry).WithField("module", "artnet"))
	control := &controller{
		logger:      logger,
		sender:      artnet.NewController(host, ip, senderLogger, artnet.MaxFPS(40)),
		state:       NewState(),
		sendTrigger: make(chan UniverseStateMap, 100),
	}

	return control, nil
}

// Start the controller
func (c *controller) Start(ctx context.Context) error {
	if err := c.sender.Start(); err != nil {
		return fmt.Errorf("failed to start Controller: %v", err)
	}

	c.context = ctx
	go c.sendBackground()
	go c.debugDevices()

	return nil
}

// Stop the controller
func (c *controller) Stop() {
	close(c.sendTrigger)
	c.sender.Stop()
}

func (c *controller) SetDMXChannelValue(value ChannelValue) {
	c.state.SetChannel(value.Universe, value.Channel, value.Value)
	c.triggerSend()
}

func (c *controller) SetDMXChannelValues(values []ChannelValue) {
	c.state.SetChannelValues(values)
	c.triggerSend()
}

// Write implements the playback.TransportWriter interface to compatibility :)
func (c *controller) Write(cmd cntl.Command) error {
	values := make([]ChannelValue, len(cmd.DMXCommands))
	for i, dmxCmd := range cmd.DMXCommands {
		values[i].Universe = uint16(dmxCmd.Universe)
		values[i].Channel = uint16(dmxCmd.Channel)
		values[i].Value = dmxCmd.Value.Uint8()
	}

	c.SetDMXChannelValues(values)

	return nil
}

func (c *controller) triggerSend() {
	c.sendTrigger <- c.state.Get()
}

func (c *controller) sendBackground() {
	for data := range c.sendTrigger {
		for u, dmx := range data {
			c.sender.SendDMXToAddress(dmx.toByteSlice(), c.universeToAddress(u))
		}
	}
}

// universeToAddress converts a dmx universe to a artnet address
// Code stolen from https://play.golang.org/p/pdQPC5u7JX
func (c *controller) universeToAddress(universe uint16) artnet.Address {
	v := make([]uint8, 2)
	binary.BigEndian.PutUint16(v, universe)

	return artnet.Address{
		Net:    v[0],
		SubUni: v[1],
	}
}

func (c *controller) debugDevices() {
	t := time.NewTicker(30 * time.Second)
	for range t.C {
		c.logger.Debugf("Currently %d devices are registered: %+s", len(c.sender.Nodes), ips(c.sender.Nodes))
	}
}

func ips(nodes []*artnet.ControlledNode) (ips []string) {
	for _, n := range nodes {
		ips = append(ips, NodeToString(n))
	}
	return
}
