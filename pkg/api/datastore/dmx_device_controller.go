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

// DMXDeviceController controls the DMXDevice entity
type DMXDeviceController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXDeviceController returns a new DMXDeviceController instance
func NewDMXDeviceController(logger *logrus.Entry, storage api.Storage) *DMXDeviceController {
	return &DMXDeviceController{
		logger:  logger,
		storage: storage,
	}
}

func (c *DMXDeviceController) validate(entity *cntl.DMXDevice) error {
	if entity.Tags == nil {
		entity.Tags = make([]cntl.Tag, 0)
	}

	if !c.storage.Has(entity.TypeID, &cntl.DMXDeviceType{}) {
		return fmt.Errorf("cannot save DMXDevice with non-existing DMXDeviceType %q", entity.TypeID)
	}

	return nil
}

// Create a new DMXDevice
func (c *DMXDeviceController) Create(r *http.Request, entity *cntl.DMXDevice, reply *cntl.DMXDevice) error {
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

// Update a new DMXDevice
func (c *DMXDeviceController) Update(r *http.Request, entity *cntl.DMXDevice, reply *cntl.DMXDevice) error {
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

// Get a DMXDevice
func (c *DMXDeviceController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXDevice) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	fmt.Println(idReq.ID)
	if !c.storage.Has(idReq.ID, &cntl.DMXDevice{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXDevice
func (c *DMXDeviceController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXDevice) error {
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
func (c *DMXDeviceController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXDevice{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXDevice{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
