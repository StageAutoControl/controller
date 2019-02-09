package api

import "errors"

var (
	rpcPath = "/rpc"
	errNoIDGiven = errors.New("no ID was given with request")
	errExists    = errors.New("entity with given ID already exists")
	errNotExists = errors.New("entity with given ID does not exist")
)

type storage interface {
	Has(key string, kind interface{}) bool
	Write(key string, value interface{}) error
	Read(key string, value interface{}) error
	List(kind interface{}) []string
	Delete(key string, kind interface{}) error
}

// IDRequest is a request object only storing an ID
type IDRequest struct {
	ID string `json:"id"`
}

// SuccessResponse returns a simple bool to state weather the operation was successful
type SuccessResponse struct {
	Success bool `json:"success"`
}

// EmptyRequest is ... yah, an empty request :shrug:
type EmptyRequest struct{}
