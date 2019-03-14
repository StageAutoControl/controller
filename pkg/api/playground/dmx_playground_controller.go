package playground

import (
	"errors"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

var (
	errControllerDisabled = errors.New("the ArtNet controller is not set, most likely it is disabled in your current instance")
)

// DMXPlaygroundController to play around and test DMX settings
type DMXPlaygroundController struct {
	logger     logging.Logger
	controller artnet.Controller
}

// NewDMXPlaygroundController returns a new DMXPlaygroundController instance
func NewDMXPlaygroundController(logger logging.Logger, controller artnet.Controller) *DMXPlaygroundController {
	return &DMXPlaygroundController{
		logger:     logger,
		controller: controller,
	}
}

// SetChannelValue sets a single artnet/dmx value
func (c *DMXPlaygroundController) SetChannelValue(r *http.Request, value *artnet.ChannelValue, response *api.Empty) error {
	if c.controller == nil {
		return errControllerDisabled
	}

	c.controller.SetDMXChannelValue(*value)
	return nil
}

// SetChannelValues sets multiple artnet/dmx values
func (c *DMXPlaygroundController) SetChannelValues(r *http.Request, values *[]artnet.ChannelValue, response *api.Empty) error {
	if c.controller == nil {
		return errControllerDisabled
	}

	c.controller.SetDMXChannelValues(*values)
	return nil
}
