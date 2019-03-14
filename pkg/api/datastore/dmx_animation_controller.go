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

// DMXAnimationController controls the DMXAnimation entity
type DMXAnimationController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXAnimationController returns a new DMXAnimationController instance
func NewDMXAnimationController(logger *logrus.Entry, storage api.Storage) *DMXAnimationController {
	return &DMXAnimationController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXAnimation
func (c *DMXAnimationController) Create(r *http.Request, entity *cntl.DMXAnimation, reply *cntl.DMXAnimation) error {
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

// Update a new DMXAnimation
func (c *DMXAnimationController) Update(r *http.Request, entity *cntl.DMXAnimation, reply *cntl.DMXAnimation) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXAnimation
func (c *DMXAnimationController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXAnimation) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXAnimation{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXAnimation
func (c *DMXAnimationController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXAnimation) error {
	for _, id := range c.storage.List(&cntl.DMXAnimation{}) {
		entity := &cntl.DMXAnimation{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXAnimation
func (c *DMXAnimationController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXAnimation{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXAnimation{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
