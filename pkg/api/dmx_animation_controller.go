package api

import (
	"fmt"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type dmxAnimationController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXAnimationController(logger *logrus.Entry, storage storage) *dmxAnimationController {
	return &dmxAnimationController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXAnimation
func (c *dmxAnimationController) Create(r *http.Request, entity *cntl.DMXAnimation, reply *cntl.DMXAnimation) error {
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

// Update a new DMXAnimation
func (c *dmxAnimationController) Update(r *http.Request, entity *cntl.DMXAnimation, reply *cntl.DMXAnimation) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	*reply = *entity
	return nil
}

// Get a DMXAnimation
func (c *dmxAnimationController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXAnimation) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXAnimation{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXAnimation
func (c *dmxAnimationController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.DMXAnimation) error {
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
func (c *dmxAnimationController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXAnimation{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXAnimation{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
