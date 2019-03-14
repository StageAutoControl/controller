package playback

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/cntl/playback"
	"github.com/StageAutoControl/controller/pkg/process"
	"github.com/jinzhu/copier"
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

// StartRequest starts the Playback of a Song or SetList
type StartRequest struct {
	playback.Params
}

// Start a playback with either a song or a setlist
func (c *Controller) Start(r *http.Request, req *StartRequest, res *process.Status) error {
	if (req.Song.ID != "" && req.SetList.ID != "") || (req.Song.ID == "" && req.SetList.ID == "") {
		return errSongSetListNeedToBeDistinct
	}

	p, _, err := c.pm.GetProcess(playback.ProcessName)
	if err != nil {
		return fmt.Errorf("failed to fetch playback process: %v", err)
	}
	p.(*playback.Process).SetParams(req.Params)

	if s, err := c.pm.Start(playback.ProcessName); err != nil {
		return fmt.Errorf("failed to start playback: %v", err)

	} else if err := copier.Copy(res, s); err != nil {
		return fmt.Errorf("failed to write response body: %v", err)
	}

	return nil
}

// Stop a playback
func (c *Controller) Stop(r *http.Request, req *api.IDBody, res *process.Status) error {
	_, err := c.pm.Stop(playback.ProcessName)
	return err
}

// Status returns the current status of a playback
func (c *Controller) Status(r *http.Request, req *api.IDBody, res *process.Status) error {
	return nil
}
