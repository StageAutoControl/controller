package playground

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/cntl/dmx"
	"github.com/StageAutoControl/controller/pkg/cntl/playback"
	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

var (
	errControllerDisabled = errors.New("the ArtNet controller is not set, most likely it is disabled in your current instance")
)

// DMXPlaygroundController to play around and test DMX settings
type DMXPlaygroundController struct {
	logger     logging.Logger
	controller artnet.Controller
	loader     api.Loader
}

// NewDMXPlaygroundController returns a new DMXPlaygroundController instance
func NewDMXPlaygroundController(logger logging.Logger, controller artnet.Controller, loader api.Loader) *DMXPlaygroundController {
	return &DMXPlaygroundController{
		loader:     loader,
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

// PlayOnceRequest is a request body to play a single entity (Scene, Preset) once
type PlayOnceRequest struct {
	api.IDBody
	cntl.BarParams
}

func (c *DMXPlaygroundController) defaultBarParams(bp *cntl.BarParams) {
	if bp.Speed == 0 {
		bp.Speed = 140
	}

	if bp.NoteCount == 0 {
		bp.NoteCount = 4
	}

	if bp.NoteValue == 0 {
		bp.NoteValue = 4
	}
}

// PlayScene plays the given Scene once
func (c *DMXPlaygroundController) PlayScene(r *http.Request, req *PlayOnceRequest, response *api.Empty) error {
	ds, err := c.loader.Load()
	if err != nil {
		return err
	}

	scene, ok := ds.DMXScenes[req.ID]
	if !ok {
		return fmt.Errorf("failed to find scene with id %s", req.ID)
	}

	dmxCommands, err := dmx.RenderScene(ds, scene)
	if err != nil {
		return fmt.Errorf("failed to render scene %s: %v", req.ID, err)
	}

	c.defaultBarParams(&req.BarParams)
	commands := playback.ToPlayable(req.BarParams, dmxCommands)
	if err := playback.Play(context.Background(), c.logger, []playback.TransportWriter{c.controller}, commands); err != nil {
		return fmt.Errorf("failed to start playback: %v", err)
	}

	return nil
}

// PlayPreset plays the given Preset once
func (c *DMXPlaygroundController) PlayPreset(r *http.Request, req *PlayOnceRequest, response *api.Empty) error {
	ds, err := c.loader.Load()
	if err != nil {
		return err
	}

	preset, ok := ds.DMXPresets[req.ID]
	if !ok {
		return fmt.Errorf("failed to find preset with id %s", req.ID)
	}

	dmxCommands, err := dmx.RenderPreset(ds, preset)
	if err != nil {
		return fmt.Errorf("failed to render preset %s: %v", req.ID, err)
	}

	c.defaultBarParams(&req.BarParams)
	commands := playback.ToPlayable(req.BarParams, dmxCommands)
	if err := playback.Play(context.Background(), c.logger, []playback.TransportWriter{c.controller}, commands); err != nil {
		return fmt.Errorf("failed to start playback: %v", err)
	}

	return nil
}
