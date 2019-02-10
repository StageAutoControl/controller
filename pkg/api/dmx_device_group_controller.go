package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type dmxDeviceGroupController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXDeviceGroupController(logger *logrus.Entry, storage storage) *dmxDeviceGroupController {
	return &dmxDeviceGroupController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXDeviceGroup
func (c *dmxDeviceGroupController) Create(r *http.Request, entity *cntl.DMXDeviceGroup, reply *cntl.DMXDeviceGroup) error {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}

	if c.storage.Has(entity.ID, entity) {
		return errExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to write to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Update a new DMXDeviceGroup
func (c *dmxDeviceGroupController) Update(r *http.Request, entity *cntl.DMXDeviceGroup, reply *cntl.DMXDeviceGroup) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXDeviceGroup
func (c *dmxDeviceGroupController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXDeviceGroup) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceGroup{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXDeviceGroup
func (c *dmxDeviceGroupController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.DMXDeviceGroup) error {
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
func (c *dmxDeviceGroupController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceGroup{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXDeviceGroup{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
