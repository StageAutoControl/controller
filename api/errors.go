package api

import (
	"errors"
	"fmt"
)

var (
	errDuplicateName = errors.New("Name is already known")
)

func makeRepoReadError(err error) error {
	return fmt.Errorf("Unable to read repository: %v", err)
}

func makeRepoSaveError(err error) error {
	return fmt.Errorf("Unable to save repository: %v", err)
}
