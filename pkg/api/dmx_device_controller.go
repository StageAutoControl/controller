package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
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

func (c *dmxDeviceController) validate(entity *cntl.DMXDevice) error {
	if entity.Tags == nil {
		entity.Tags = make([]cntl.Tag, 0)
	}

	if !c.storage.Has(entity.TypeID, &cntl.DMXDeviceType{}) {
		return fmt.Errorf("cannot save DMXDevice with non-existing DMXDeviceType %q", entity.TypeID)
	}

	return nil
}

// Create a new DMXDevice
func (c *dmxDeviceController) Create(r *http.Request, entity *cntl.DMXDevice, reply *cntl.DMXDevice) error {
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

// Update a new DMXDevice
func (c *dmxDeviceController) Update(r *http.Request, entity *cntl.DMXDevice, reply *cntl.DMXDevice) error {
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

// Get a DMXDevice
func (c *dmxDeviceController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXDevice) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	fmt.Println(idReq.ID)
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
