package api

import (
	"net/http"

	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

type DMXPlaygroundController struct {
	logger     logging.Logger
	controller artnet.Controller
}

func newDMXPlaygroundController(logger logging.Logger, controller artnet.Controller) *DMXPlaygroundController {
	return &DMXPlaygroundController{
		logger:     logger,
		controller: controller,
	}
}

// SetChannelValue sets a single artnet/dmx value
func (c *DMXPlaygroundController) SetChannelValue(r *http.Request, value *artnet.ChannelValue, request Empty) error {
	c.controller.SetDMXChannelValue(*value)
	return nil
}

// SetChannelValues sets multiple artnet/dmx values
func (c *DMXPlaygroundController) SetChannelValues(r *http.Request, values *[]artnet.ChannelValue, request Empty) error {
	c.controller.SetDMXChannelValues(*values)
	return nil
}
