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

// SetListController controls the SetList entity
type SetListController struct {
	logger  *logrus.Entry
	storage api.Storage
}

// NewSetListController returns a new SetListController instance
func NewSetListController(logger *logrus.Entry, storage api.Storage) *SetListController {
	return &SetListController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new SetList
func (c *SetListController) Create(r *http.Request, entity *cntl.SetList, reply *cntl.SetList) error {
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

// Update a new SetList
func (c *SetListController) Update(r *http.Request, entity *cntl.SetList, reply *cntl.SetList) error {
	if !c.storage.Has(entity.ID, entity) {
		return api.ErrNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a SetList
func (c *SetListController) Get(r *http.Request, idReq *api.IDBody, reply *cntl.SetList) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.SetList{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of SetList
func (c *SetListController) GetAll(r *http.Request, idReq *api.Empty, reply *[]*cntl.SetList) error {
	for _, id := range c.storage.List(&cntl.SetList{}) {
		entity := &cntl.SetList{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a SetList
func (c *SetListController) Delete(r *http.Request, idReq *api.IDBody, reply *api.SuccessResponse) error {
	if idReq.ID == "" {
		return api.ErrNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.SetList{}) {
		return api.ErrNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.SetList{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
