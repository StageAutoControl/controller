package api

import "errors"

var (
	errNoIDGiven = errors.New("no ID was given with request")
	errExists = errors.New("entity with given ID already exists")
	errNotExists = errors.New("entity with given ID does not exist")
)

type storage interface {
	Has(key string, kind interface{}) bool
	Write(key string, value interface{}) error
	Read(key string, value interface{}) error
	List(kind interface{}) []string
	Delete(key string, kind interface{}) error
}

type IDRequest struct {
	ID string `json:"id"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type EmptyRequest struct {}
