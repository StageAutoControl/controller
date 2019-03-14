package datastore

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestDMXColorVariableController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"
	entity := ds.DMXColorVariables[key]

	createReply := &cntl.DMXColorVariable{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXColorVariableController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"
	entity := ds.DMXColorVariables[key]

	createEntity := &cntl.DMXColorVariable{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.DMXColorVariable{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXColorVariableController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"

	reply := &cntl.DMXColorVariable{}

	idReq := &api.IDBody{ID: key}
	if err := controller.Get(req, idReq, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXColorVariableController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"
	entity := ds.DMXColorVariables[key]

	createReply := &cntl.DMXColorVariable{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXColorVariable{}
	idReq := &api.IDBody{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestDMXColorVariableController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"
	entity := ds.DMXColorVariables[key]

	reply := &cntl.DMXColorVariable{}

	if err := controller.Update(req, entity, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXColorVariableController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"
	entity := ds.DMXColorVariables[key]

	createReply := &cntl.DMXColorVariable{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXColorVariable{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestDMXColorVariableController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"

	reply := &api.SuccessResponse{}
	idReq := &api.IDBody{ID: key}
	if err := controller.Delete(req, idReq, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXColorVariableController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXColorVariableController(logger, store)
	key := "4b848ea8-5094-4509-a067-09a0e568220d"
	entity := ds.DMXColorVariables[key]

	createReply := &cntl.DMXColorVariable{}
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
