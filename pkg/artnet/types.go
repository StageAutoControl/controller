package artnet

import (
	"github.com/jsimonetti/go-artnet"

	"github.com/StageAutoControl/controller/pkg/cntl"
)

// Sender is an artnet controller abstraction of the base implementation of jsimonetti
type Sender interface {
	SendDMXToAddress(dmx [512]byte, address artnet.Address)
	Start() error
	Stop()
}

// ChannelValue defines an ArtNet Universe and the value of the DMX channel
type ChannelValue struct {
	Universe       uint16
	Channel, Value uint8
}

// Controller is a convenience interface to use within this application
type Controller interface {
	Write(cntl.Command) error
	SetDMXChannelValue(value ChannelValue)
	SetDMXChannelValues(values []ChannelValue)
	Start() error
	Stop()
}
