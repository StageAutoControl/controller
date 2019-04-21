package api

import (
	"errors"

	"github.com/StageAutoControl/controller/pkg/cntl"
)

var (
	// RPCPath to where the RPC server should listen on
	RPCPath = "/api"

	// ErrNoIDGiven is returned when the request did not contain a valid ID
	ErrNoIDGiven = errors.New("no ID was given with request")
	// ErrExists is returned when the entity which is tried to create already exists
	ErrExists = errors.New("entity with given ID already exists")
	// ErrNotExists is returned when the entity tried to manage dies not exist
	ErrNotExists = errors.New("entity with given ID does not exist")
)

// Storage interface for abstraction in api usage
type Storage interface {
	Has(key string, kind interface{}) bool
	Write(key string, value interface{}) error
	Read(key string, value interface{}) error
	List(kind interface{}) []string
	Delete(key string, kind interface{}) error
}

// Loader interface for abstraction in api usage
type Loader interface {
	Load() (*cntl.DataStore, error)
}

// IDBody is a request object only storing an ID
type IDBody struct {
	ID string `json:"id"`
}

// SuccessResponse returns a simple bool to state weather the operation was successful
type SuccessResponse struct {
	Success bool `json:"success"`
}

// Empty is ... yah, an empty request :shrug:
type Empty struct{}
