package api

import (
	"fmt"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type dmxPresetController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXPresetController(logger *logrus.Entry, storage storage) *dmxPresetController {
	return &dmxPresetController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXPreset
func (c *dmxPresetController) Create(r *http.Request, entity *cntl.DMXPreset, reply *cntl.DMXPreset) error {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	} else {
		if c.storage.Has(entity.ID, entity) {
			return errExists
		}
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to write to disk: %v", err)
	}

	*reply = *entity
	return nil
}

// Update a new DMXPreset
func (c *dmxPresetController) Update(r *http.Request, entity *cntl.DMXPreset, reply *cntl.DMXPreset) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	*reply = *entity
	return nil
}

// Get a DMXPreset
func (c *dmxPresetController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXPreset) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXPreset{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXPreset
func (c *dmxPresetController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.DMXPreset) error {
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
func (c *dmxPresetController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXPreset{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXPreset{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
