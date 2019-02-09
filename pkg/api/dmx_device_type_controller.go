package api

import (
	"fmt"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type dmxDeviceTypeController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXDeviceTypeController(logger *logrus.Entry, storage storage) *dmxDeviceTypeController {
	return &dmxDeviceTypeController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXDeviceType
func (c *dmxDeviceTypeController) Create(r *http.Request, entity *cntl.DMXDeviceType, reply *cntl.DMXDeviceType) error {
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

// Update a new DMXDeviceType
func (c *dmxDeviceTypeController) Update(r *http.Request, entity *cntl.DMXDeviceType, reply *cntl.DMXDeviceType) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	*reply = *entity
	return nil
}

// Get a DMXDeviceType
func (c *dmxDeviceTypeController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXDeviceType) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceType{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXDeviceType
func (c *dmxDeviceTypeController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.DMXDeviceType) error {
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
func (c *dmxDeviceTypeController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDeviceType{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXDeviceType{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
