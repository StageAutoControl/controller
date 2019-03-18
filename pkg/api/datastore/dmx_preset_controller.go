package datastore

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// DMXPresetController controls the DMXPreset entity
type DMXPresetController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXPresetController returns a new DMXPresetController instance
func NewDMXPresetController(logger *logrus.Entry, storage api.Storage) *DMXPresetController {
	return &DMXPresetController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXPreset
func (c *DMXPresetController) Create(r *http.Request, entity *cntl.DMXPreset, reply *cntl.DMXPreset) error {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}

	if c.storage.Has(entity.ID, entity) {
		return api.ErrExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to write to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Update a new DMXPreset
func (c *DMXPresetController) Update(r *http.Request, entity *cntl.DMXPreset, reply *cntl.DMXPreset) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXPreset
func (c *DMXPresetController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXPreset) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXPreset{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXPreset
func (c *DMXPresetController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXPreset) error {
	*reply = []*cntl.DMXPreset{}
	for _, id := range c.storage.List(&cntl.DMXPreset{}) {
		entity := &cntl.DMXPreset{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXPreset
func (c *DMXPresetController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXPreset{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXPreset{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
