package api

import (
	"sync"

	"github.com/StageAutoControl/controller/cntl"
)

// EmptyRequest is used e.g. for Get request where no args need to be passed.
type EmptyRequest struct{}

// BaseResponse is the basic response for all handlers. Might be extended for detailed responses.
type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type repo interface {
	sync.Locker
	Load() (*cntl.DataStore, error)
	Save(store *cntl.DataStore) error
}
