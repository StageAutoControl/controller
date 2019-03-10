package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
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

func (c *dmxDeviceTypeController) validate(entity *cntl.DMXDeviceType) error {
	if entity.LEDs == nil {
		entity.LEDs = make([]cntl.LED, 0)
	}

	return nil
}

// Create a new DMXDeviceType
func (c *dmxDeviceTypeController) Create(r *http.Request, entity *cntl.DMXDeviceType, reply *cntl.DMXDeviceType) error {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}

	if c.storage.Has(entity.ID, entity) {
		return errExists
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
func (c *dmxDeviceTypeController) Update(r *http.Request, entity *cntl.DMXDeviceType, reply *cntl.DMXDeviceType) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
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
func (c *dmxDeviceTypeController) GetAll(r *http.Request, idReq *Empty, reply *[]*cntl.DMXDeviceType) error {
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
