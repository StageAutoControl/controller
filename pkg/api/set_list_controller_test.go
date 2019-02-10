package api

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestSetListController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"
	entity := ds.SetLists[key]

	createReply := &cntl.SetList{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestSetListController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"
	entity := ds.SetLists[key]

	createEntity := &cntl.SetList{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.SetList{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestSetListController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"

	reply := &cntl.SetList{}

	idReq := &IDRequest{ID: key}
	if err := controller.Get(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestSetListController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"
	entity := ds.SetLists[key]

	createReply := &cntl.SetList{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.SetList{}
	idReq := &IDRequest{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestSetListController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"
	entity := ds.SetLists[key]

	reply := &cntl.SetList{}

	if err := controller.Update(req, entity, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestSetListController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"
	entity := ds.SetLists[key]

	createReply := &cntl.SetList{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.SetList{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestSetListController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"

	reply := &SuccessResponse{}
	idReq := &IDRequest{ID: key}
	if err := controller.Delete(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestSetListController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSetListController(logger, store)
	key := "f5b4be8a-0b18-11e7-b837-4bac99d86956"
	entity := ds.SetLists[key]

	createReply := &cntl.SetList{}
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
