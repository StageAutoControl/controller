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

// DMXDeviceTypeController controls the DMXDeviceType entity
type DMXDeviceTypeController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXDeviceTypeController returns a new DMXDeviceTypeController instance
func NewDMXDeviceTypeController(logger *logrus.Entry, storage api.Storage) *DMXDeviceTypeController {
	return &DMXDeviceTypeController{
		logger:  logger,
		storage: storage,
	}
}

func (c *DMXDeviceTypeController) validate(entity *cntl.DMXDeviceType) error {
	if entity.LEDs == nil {
		entity.LEDs = make([]cntl.LED, 0)
	}

	return nil
}

// Create a new DMXDeviceType
func (c *DMXDeviceTypeController) Create(r *http.Request, entity *cntl.DMXDeviceType, reply *cntl.DMXDeviceType) error {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}

	if c.storage.Has(entity.ID, entity) {
		return api.ErrExists
	}

	if err := c.validate(entity); err != nil {
		return fmt.Errorf("failed to validate entity: %v", err)
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to write to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Update a new DMXDeviceType
func (c *DMXDeviceTypeController) Update(r *http.Request, entity *cntl.DMXDeviceType, reply *cntl.DMXDeviceType) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.validate(entity); err != nil {
		return fmt.Errorf("failed to validate entity: %v", err)
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXDeviceType
func (c *DMXDeviceTypeController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXDeviceType) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceType{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXDeviceType
func (c *DMXDeviceTypeController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXDeviceType) error {
	*reply = []*cntl.DMXDeviceType{}
	for _, id := range c.storage.List(&cntl.DMXDeviceType{}) {
		entity := &cntl.DMXDeviceType{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXDeviceType
func (c *DMXDeviceTypeController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceType{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXDeviceType{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
