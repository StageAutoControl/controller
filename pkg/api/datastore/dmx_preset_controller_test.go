package datastore

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestDMXPresetController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"
	entity := ds.DMXPresets[key]

	createReply := &cntl.DMXPreset{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXPresetController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"
	entity := ds.DMXPresets[key]

	createEntity := &cntl.DMXPreset{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.DMXPreset{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXPresetController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"

	reply := &cntl.DMXPreset{}

	idReq := &api.IDBody{ID: key}
	if err := controller.Get(req, idReq, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXPresetController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"
	entity := ds.DMXPresets[key]

	createReply := &cntl.DMXPreset{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXPreset{}
	idReq := &api.IDBody{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestDMXPresetController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"
	entity := ds.DMXPresets[key]

	reply := &cntl.DMXPreset{}

	if err := controller.Update(req, entity, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXPresetController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"
	entity := ds.DMXPresets[key]

	createReply := &cntl.DMXPreset{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXPreset{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestDMXPresetController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"

	reply := &api.SuccessResponse{}
	idReq := &api.IDBody{ID: key}
	if err := controller.Delete(req, idReq, reply); err != api.ErrNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXPresetController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXPresetController(logger, store)
	key := "0de258e0-0e7b-11e7-afd4-ebf6036983dc"
	entity := ds.DMXPresets[key]

	createReply := &cntl.DMXPreset{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &api.SuccessResponse{}
	idReq := &api.IDBody{ID: key}
	if err := controller.Delete(req, idReq, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if !reply.Success {
		t.Error("Expected to get result true, but got false")
	}
}
