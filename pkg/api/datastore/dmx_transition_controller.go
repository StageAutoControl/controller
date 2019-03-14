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

// DMXTransitionController controls the DMXTransition entity
type DMXTransitionController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXTransitionController returns a new DMXTransitionController instance
func NewDMXTransitionController(logger *logrus.Entry, storage api.Storage) *DMXTransitionController {
	return &DMXTransitionController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXTransition
func (c *DMXTransitionController) Create(r *http.Request, entity *cntl.DMXTransition, reply *cntl.DMXTransition) error {
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

// Update a new DMXTransition
func (c *DMXTransitionController) Update(r *http.Request, entity *cntl.DMXTransition, reply *cntl.DMXTransition) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXTransition
func (c *DMXTransitionController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXTransition) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXTransition{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXTransition
func (c *DMXTransitionController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXTransition) error {
	for _, id := range c.storage.List(&cntl.DMXTransition{}) {
		entity := &cntl.DMXTransition{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXTransition
func (c *DMXTransitionController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXTransition{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXTransition{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
