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

// DMXSceneController controls the DMXScene entity
type DMXSceneController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewDMXSceneController returns a new DMXSceneController instance
func NewDMXSceneController(logger *logrus.Entry, storage api.Storage) *DMXSceneController {
	return &DMXSceneController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXScene
func (c *DMXSceneController) Create(r *http.Request, entity *cntl.DMXScene, reply *cntl.DMXScene) error {
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

// Update a new DMXScene
func (c *DMXSceneController) Update(r *http.Request, entity *cntl.DMXScene, reply *cntl.DMXScene) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXScene
func (c *DMXSceneController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.DMXScene) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXScene{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXScene
func (c *DMXSceneController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.DMXScene) error {
	*reply = []*cntl.DMXScene{}
	for _, id := range c.storage.List(&cntl.DMXScene{}) {
		entity := &cntl.DMXScene{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a DMXScene
func (c *DMXSceneController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXScene{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXScene{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
