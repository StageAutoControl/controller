package transport

import (
	art "github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/cntl"
)

// ArtNet is a transport for the ArtNet protocol (DMX over UDP/IP)
type ArtNet struct {
	controller art.Controller
}

// NewArtNet returns a new ArtNet transport instance
func NewArtNet(controller art.Controller) (*ArtNet, error) {
	return &ArtNet{
		controller: controller,
	}, nil
}

func (a *ArtNet) Write(cmd cntl.Command) error {
	values := make([]art.ChannelValue, 0)

	for _, c := range cmd.DMXCommands {
		values = append(values, art.ChannelValue{
			Universe: uint16(c.Universe),
			Channel:  uint16(c.Channel),
			Value:    c.Value.Uint8(),
		})
	}

	a.controller.SetDMXChannelValues(values)

	return nil
}
