package api

import (
	"errors"
	"testing"
)

func TestMakeRepoReadError(t *testing.T) {
	s := errors.New("str")
	exp := "Unable to read repository: str"

	err := makeRepoReadError(s)

	if err.Error() != exp {
		t.Errorf("Expected to get %q, got %q", exp, err.Error())
	}
}

func TestMakeRepoSaveError(t *testing.T) {
	s := errors.New("str")
	exp := "Unable to save repository: str"

	err := makeRepoSaveError(s)

	if err.Error() != exp {
		t.Errorf("Expected to get %q, got %q", exp, err.Error())
	}
}
