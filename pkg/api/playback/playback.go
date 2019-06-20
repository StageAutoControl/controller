package playback

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/cntl/playback"
	"github.com/StageAutoControl/controller/pkg/process"
)

var (
	errSongSetListNeedToBeDistinct = errors.New("either a SetList ID or a Song ID must be given, neither both nor none")
)

// Controller handles management of playback processes
type Controller struct {
	pm process.Manager
}

// NewController returns a new playback controller instance
func NewController(pm process.Manager) *Controller {
	return &Controller{
		pm: pm,
	}
}

// Status Response of the playback process
type Status struct {
	Process process.Status  `json:"process"`
	Params  playback.Params `json:"params"`
}

// Start a playback with either a song or a setlist
func (c *Controller) Start(r *http.Request, req *playback.Params, res *Status) error {
	if (req.Song.ID != "" && req.SetList.ID != "") || (req.Song.ID == "" && req.SetList.ID == "") {
		return errSongSetListNeedToBeDistinct
	}

	p, _, err := c.pm.GetProcess(playback.ProcessName)
	if err != nil {
		return fmt.Errorf("failed to fetch playback process: %v", err)
	}
	p.(*playback.Process).SetParams(*req)

	s, err := c.pm.Start(playback.ProcessName)
	if err != nil {
		return fmt.Errorf("failed to start playback: %v", err)
	}

	res.Process = *s
	res.Params = *req

	return nil
}

// Stop a playback
func (c *Controller) Stop(r *http.Request, req *api.IDBody, res *Status) error {
	p, _, err := c.pm.GetProcess(playback.ProcessName)
	if err != nil {
		return fmt.Errorf("failed to get playback status: %v", err)
	}

	s, err := c.pm.Stop(playback.ProcessName)
	if err != nil {
		return fmt.Errorf("failed to stop playback: %v", err)
	}

	res.Process = *s
	res.Params = p.(*playback.Process).GetParams()

	return err
}

// Status returns the current status of a playback
func (c *Controller) Status(r *http.Request, req *api.IDBody, res *Status) error {
	p, s, err := c.pm.GetProcess(playback.ProcessName)
	if err != nil {
		return fmt.Errorf("failed to get playback status: %v", err)
	}

	res.Process = *s
	res.Params = p.(*playback.Process).GetParams()

	return nil
}
