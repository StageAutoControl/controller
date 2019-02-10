package api

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestSongController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"
	entity := ds.Songs[key]

	createReply := &cntl.Song{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestSongController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"
	entity := ds.Songs[key]

	createEntity := &cntl.Song{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.Song{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestSongController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"

	reply := &cntl.Song{}

	idReq := &IDRequest{ID: key}
	if err := controller.Get(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestSongController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"
	entity := ds.Songs[key]

	createReply := &cntl.Song{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.Song{}
	idReq := &IDRequest{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestSongController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"
	entity := ds.Songs[key]

	reply := &cntl.Song{}

	if err := controller.Update(req, entity, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestSongController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"
	entity := ds.Songs[key]

	createReply := &cntl.Song{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.Song{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestSongController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"

	reply := &SuccessResponse{}
	idReq := &IDRequest{ID: key}
	if err := controller.Delete(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestSongController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newSongController(logger, store)
	key := "3c1065c8-0b14-11e7-96eb-5b134621c411"
	entity := ds.Songs[key]

	createReply := &cntl.Song{}
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
