package api

import (
	"fmt"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type dmxTransitionController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXTransitionController(logger *logrus.Entry, storage storage) *dmxTransitionController {
	return &dmxTransitionController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXTransition
func (c *dmxTransitionController) Create(r *http.Request, entity *cntl.DMXTransition, reply *cntl.DMXTransition) error {
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

// Update a new DMXTransition
func (c *dmxTransitionController) Update(r *http.Request, entity *cntl.DMXTransition, reply *cntl.DMXTransition) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	*reply = *entity
	return nil
}

// Get a DMXTransition
func (c *dmxTransitionController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXTransition) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXTransition{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXTransition
func (c *dmxTransitionController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.DMXTransition) error {
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
func (c *dmxTransitionController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXTransition{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXTransition{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
