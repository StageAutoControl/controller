package api

import (
	"fmt"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type dmxDeviceController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXDeviceController(logger *logrus.Entry, storage storage) *dmxDeviceController {
	return &dmxDeviceController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXDevice
func (c *dmxDeviceController) Create(r *http.Request, entity *cntl.DMXDevice, reply *cntl.DMXDevice) error {
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

// Update a new DMXDevice
func (c *dmxDeviceController) Update(r *http.Request, entity *cntl.DMXDevice, reply *cntl.DMXDevice) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	*reply = *entity
	return nil
}

// Get a DMXDevice
func (c *dmxDeviceController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXDevice) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDevice{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXDevice
func (c *dmxDeviceController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.DMXDevice) error {
	for _, id := range c.storage.List(&cntl.DMXDevice{}) {
		entity := &cntl.DMXDevice{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXDevice
func (c *dmxDeviceController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDevice{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXDevice{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
