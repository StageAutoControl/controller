package api

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestDMXTransitionController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"
	entity := ds.DMXTransitions[key]

	createReply := &cntl.DMXTransition{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXTransitionController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"
	entity := ds.DMXTransitions[key]

	createEntity := &cntl.DMXTransition{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.DMXTransition{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXTransitionController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"

	reply := &cntl.DMXTransition{}

	idReq := &IDRequest{ID: key}
	if err := controller.Get(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXTransitionController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"
	entity := ds.DMXTransitions[key]

	createReply := &cntl.DMXTransition{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXTransition{}
	idReq := &IDRequest{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestDMXTransitionController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"
	entity := ds.DMXTransitions[key]

	reply := &cntl.DMXTransition{}

	if err := controller.Update(req, entity, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXTransitionController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"
	entity := ds.DMXTransitions[key]

	createReply := &cntl.DMXTransition{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXTransition{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestDMXTransitionController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"

	reply := &SuccessResponse{}
	idReq := &IDRequest{ID: key}
	if err := controller.Delete(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXTransitionController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXTransitionController(logger, store)
	key := "a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"
	entity := ds.DMXTransitions[key]

	createReply := &cntl.DMXTransition{}
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
