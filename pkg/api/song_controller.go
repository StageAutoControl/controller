package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type songController struct {
	logger  *logrus.Entry
	storage storage
}

func newSongController(logger *logrus.Entry, storage storage) *songController {
	return &songController{
		logger:  logger,
		storage: storage,
	}
}

// Create a new Song
func (c *songController) Create(r *http.Request, entity *cntl.Song, reply *cntl.Song) error {
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

// Update a new Song
func (c *songController) Update(r *http.Request, entity *cntl.Song, reply *cntl.Song) error {
	if !c.storage.Has(entity.ID, entity) {
		return errNotExists
	}

	if err := c.storage.Write(entity.ID, entity); err != nil {
		return fmt.Errorf("failed to update to disk: %v", err)
	}

	return copier.Copy(reply, entity)
}

// Get a Song
func (c *songController) Get(r *http.Request, idReq *IDRequest, reply *cntl.Song) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.Song{}) {
		return errNotExists
	}

	if err := c.storage.Read(idReq.ID, reply); err != nil {
		return fmt.Errorf("failed to read entity: %v", err)
	}

	return nil
}

// GetAll returns all entities of Song
func (c *songController) GetAll(r *http.Request, idReq *EmptyRequest, reply *[]*cntl.Song) error {
	for _, id := range c.storage.List(&cntl.Song{}) {
		entity := &cntl.Song{}
		if err := c.storage.Read(id, entity); err != nil {
			return fmt.Errorf("failed to read entity %s: %v", id, err)
		}
		*reply = append(*reply, entity)
	}

	return nil
}

// Delete a Song
func (c *songController) Delete(r *http.Request, idReq *IDRequest, reply *SuccessResponse) error {
	if idReq.ID == "" {
		return errNoIDGiven
	}

	if !c.storage.Has(idReq.ID, &cntl.Song{}) {
		return errNotExists
	}

	if err := c.storage.Delete(idReq.ID, &cntl.Song{}); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	reply.Success = true
	return nil
}
