package api

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/cntl"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	"github.com/jinzhu/copier"
)

func TestDMXDeviceTypeController_Create_WithID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"
	entity := ds.DMXDeviceTypes[key]

	createReply := &cntl.DMXDeviceType{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXDeviceTypeController_Create_WithoutID(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"
	entity := ds.DMXDeviceTypes[key]

	createEntity := &cntl.DMXDeviceType{}
	if err := copier.Copy(createEntity, entity); err != nil {
		t.Fatal(err)
	}

	createEntity.ID = ""

	createReply := &cntl.DMXDeviceType{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}
}

func TestDMXDeviceTypeController_Get_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"

	reply := &cntl.DMXDeviceType{}

	idReq := &IDRequest{ID: key}
	if err := controller.Get(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXDeviceTypeController_Get_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"
	entity := ds.DMXDeviceTypes[key]

	createReply := &cntl.DMXDeviceType{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXDeviceType{}
	idReq := &IDRequest{ID: key}
	t.Log("idReq has ID:", idReq.ID)
	if err := controller.Get(req, idReq, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}

func TestDMXDeviceTypeController_Update_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"
	entity := ds.DMXDeviceTypes[key]

	reply := &cntl.DMXDeviceType{}

	if err := controller.Update(req, entity, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXDeviceTypeController_Update_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"
	entity := ds.DMXDeviceTypes[key]

	createReply := &cntl.DMXDeviceType{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if createReply.ID != key {
		t.Errorf("Expected createReply to have id %s, but has %s", key, createReply.ID)
	}

	reply := &cntl.DMXDeviceType{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if reply.ID != key {
		t.Errorf("Expected reply to have id %s, but has %s", key, reply.ID)
	}
}
func TestDMXDeviceTypeController_Delete_NotExisting(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"

	reply := &SuccessResponse{}
	idReq := &IDRequest{ID: key}
	if err := controller.Delete(req, idReq, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDMXDeviceTypeController_Delete_Existing(t *testing.T) {
	defer internalTesting.Cleanup(t, path)
	controller := newDMXDeviceTypeController(logger, store)
	key := "628fc3ea-1188-11e7-8824-5f72d80c17b6"
	entity := ds.DMXDeviceTypes[key]

	createReply := &cntl.DMXDeviceType{}
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
