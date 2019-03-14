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

// DMXColorVariableController controls the DMXColorVariable entity
type DMXColorVariableController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXColorVariableController returns a new MXColorVariableController instance
func NewDMXColorVariableController(logger *logrus.Entry, storage api.Storage) *DMXColorVariableController {
	return &DMXColorVariableController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXColorVariable
func (c *DMXColorVariableController) Create(r *http.Request, entity *cntl.DMXColorVariable, reply *cntl.DMXColorVariable) error {
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

// Update a new DMXColorVariable
func (c *DMXColorVariableController) Update(r *http.Request, entity *cntl.DMXColorVariable, reply *cntl.DMXColorVariable) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXColorVariable
func (c *DMXColorVariableController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXColorVariable) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	fmt.Println(idReq.ID)
	if !c.storage.Has(idReq.ID, &cntl.DMXColorVariable{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXColorVariable
func (c *DMXColorVariableController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXColorVariable) error {
	for _, id := range c.storage.List(&cntl.DMXColorVariable{}) {
		entity := &cntl.DMXColorVariable{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXColorVariable
func (c *DMXColorVariableController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXColorVariable{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXColorVariable{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
