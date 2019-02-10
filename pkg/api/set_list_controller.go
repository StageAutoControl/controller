package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type setListController struct {
	logger  *logrus.Entry
	storage storage
}

func newSetListController(logger *logrus.Entry, storage storage) *setListController {
	return &setListController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new SetList
func (c *setListController) Create(r *http.Request, entity *cntl.SetList, reply *cntl.SetList) error {
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

// Update a new SetList
func (c *setListController) Update(r *http.Request, entity *cntl.SetList, reply *cntl.SetList) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a SetList
func (c *setListController) Get(r *http.Request, idReq *IDRequest, reply *cntl.SetList) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.SetList{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of SetList
func (c *setListController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.SetList) error {
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
func (c *setListController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.SetList{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.SetList{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
