package api

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestDMXSceneController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"
	entity := ds.DMXScenes[key]

	createReply := &cntl.DMXScene{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXSceneController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"
	entity := ds.DMXScenes[key]

	createEntity := &cntl.DMXScene{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.DMXScene{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXSceneController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"

	reply := &cntl.DMXScene{}

	idReq := &IDRequest{ID: key}
	if err := controller.Get(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXSceneController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"
	entity := ds.DMXScenes[key]

	createReply := &cntl.DMXScene{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXScene{}
	idReq := &IDRequest{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestDMXSceneController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"
	entity := ds.DMXScenes[key]

	reply := &cntl.DMXScene{}

	if err := controller.Update(req, entity, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXSceneController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"
	entity := ds.DMXScenes[key]

	createReply := &cntl.DMXScene{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXScene{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestDMXSceneController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"

	reply := &SuccessResponse{}
	idReq := &IDRequest{ID: key}
	if err := controller.Delete(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXSceneController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXSceneController(logger, store)
	key := "492cef2e-0b14-11e7-be89-c3fa25f9cabb"
	entity := ds.DMXScenes[key]

	createReply := &cntl.DMXScene{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &SuccessResponse{}
	idReq := &IDRequest{ID: key}
	if err := controller.Delete(req, idReq, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if !reply.Success {
		t.Error("Expected to get result true, but got false")
	}
}
