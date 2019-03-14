package datastore

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestDMXAnimationController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"
	entity := ds.DMXAnimations[key]

	createReply := &cntl.DMXAnimation{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXAnimationController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"
	entity := ds.DMXAnimations[key]

	createEntity := &cntl.DMXAnimation{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.DMXAnimation{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXAnimationController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"

	reply := &cntl.DMXAnimation{}

	idReq := &api.IDBody{ID: key}
	if err := controller.Get(req, idReq, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXAnimationController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"
	entity := ds.DMXAnimations[key]

	createReply := &cntl.DMXAnimation{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXAnimation{}
	idReq := &api.IDBody{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestDMXAnimationController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"
	entity := ds.DMXAnimations[key]

	reply := &cntl.DMXAnimation{}

	if err := controller.Update(req, entity, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXAnimationController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"
	entity := ds.DMXAnimations[key]

	createReply := &cntl.DMXAnimation{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call apiController: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXAnimation{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestDMXAnimationController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"

	reply := &api.SuccessResponse{}
	idReq := &api.IDBody{ID: key}
	if err := controller.Delete(req, idReq, reply); err != api.ErrNotExists {
		t.Errorf("expected to get api.ErrNotExists, but got %v", err)
	}
}

func TestDMXAnimationController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := NewDMXAnimationController(logger, store)
	key := "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"
	entity := ds.DMXAnimations[key]

	createReply := &cntl.DMXAnimation{}
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
