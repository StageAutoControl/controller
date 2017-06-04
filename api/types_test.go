package api

import (
	"github.com/StageAutoControl/controller/cntl"
)

type testRepo struct {
	locked bool
}

func (r *testRepo) Load() (*cntl.DataStore, error) {
	return &cntl.DataStore{}, nil
}

func (r *testRepo) Save(store *cntl.DataStore) error {
	return nil
}

func (r *testRepo) Lock() {
	r.locked = true
}

func (r *testRepo) Unlock() {
	r.locked = false
}
