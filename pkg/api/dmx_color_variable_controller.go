package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type dmxColorVariableController struct {
	logger  *logrus.Entry
	storage storage
}

func newDMXColorVariableController(logger *logrus.Entry, storage storage) *dmxColorVariableController {
	return &dmxColorVariableController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new DMXColorVariable
func (c *dmxColorVariableController) Create(r *http.Request, entity *cntl.DMXColorVariable, reply *cntl.DMXColorVariable) error {
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

// Update a new DMXColorVariable
func (c *dmxColorVariableController) Update(r *http.Request, entity *cntl.DMXColorVariable, reply *cntl.DMXColorVariable) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a DMXColorVariable
func (c *dmxColorVariableController) Get(r *http.Request, idReq *IDRequest, reply *cntl.DMXColorVariable) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	fmt.Println(idReq.ID)
	if !c.storage.Has(idReq.ID, &cntl.DMXColorVariable{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of DMXColorVariable
func (c *dmxColorVariableController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.DMXColorVariable) error {
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
func (c *dmxColorVariableController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.DMXColorVariable{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.DMXColorVariable{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
