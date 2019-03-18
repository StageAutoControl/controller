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

// DMXDeviceGroupController controls the DMXDeviceGroup entity
type DMXDeviceGroupController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXDeviceGroupController returns a new DMXDeviceGroupController instance
func NewDMXDeviceGroupController(logger *logrus.Entry, storage api.Storage) *DMXDeviceGroupController {
	return &DMXDeviceGroupController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXDeviceGroup
func (c *DMXDeviceGroupController) Create(r *http.Request, entity *cntl.DMXDeviceGroup, reply *cntl.DMXDeviceGroup) error {
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

// Update a new DMXDeviceGroup
func (c *DMXDeviceGroupController) Update(r *http.Request, entity *cntl.DMXDeviceGroup, reply *cntl.DMXDeviceGroup) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXDeviceGroup
func (c *DMXDeviceGroupController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXDeviceGroup) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceGroup{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXDeviceGroup
func (c *DMXDeviceGroupController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXDeviceGroup) error {
	*reply = []*cntl.DMXDeviceGroup{}
	for _, id := range c.storage.List(&cntl.DMXDeviceGroup{}) {
		entity := &cntl.DMXDeviceGroup{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXDeviceGroup
func (c *DMXDeviceGroupController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceGroup{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXDeviceGroup{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
