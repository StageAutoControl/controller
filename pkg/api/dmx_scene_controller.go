package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type dmxSceneController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXSceneController(logger *logrus.Entry, storage storage) *dmxSceneController {
	return &dmxSceneController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXScene
func (c *dmxSceneController) Create(r *http.Request, entity *cntl.DMXScene, reply *cntl.DMXScene) error {
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

// Update a new DMXScene
func (c *dmxSceneController) Update(r *http.Request, entity *cntl.DMXScene, reply *cntl.DMXScene) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXScene
func (c *dmxSceneController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXScene) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXScene{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXScene
func (c *dmxSceneController) GetAll(r *http.Request, idReq *Empty, reply *[]*cntl.DMXScene) error {
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
func (c *dmxSceneController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXScene{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXScene{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
