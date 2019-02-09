package api

import (
	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/internal/fixtures"
	internalTesting "github.com/StageAutoControl/controller/pkg/internal/testing"
	disk "github.com/StageAutoControl/controller/pkg/storage"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

var (
	ds  = fixtures.DataStore()
	key = "35cae00a-0b17-11e7-8bca-bbf30c56f20e"
)

func TestDmxDeviceController_Create_WithID(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	entity := ds.DMXDevices[key]
	reply := &cntl.DMXDevice{}

	if err := controller.Create(req, entity, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID == "" {
		t.Errorf("exptected to get reply with an ID set, but isn't")
	}
}

func TestDmxDeviceController_Create_WithoutID(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	entity := ds.DMXDevices[key]
	entity.ID = ""
	reply := &cntl.DMXDevice{}

	if err := controller.Create(req, entity, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID == "" {
		t.Errorf("exptected to get reply with an ID set, but isn't")
	}
}

func TestDmxDeviceController_Get_NotExisting(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	reply := &cntl.DMXDevice{}

	if err := controller.Get(req, &IDRequest{ID: key}, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDmxDeviceController_Get_Existing(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	entity := ds.DMXDevices[key]

	createReply := &cntl.DMXDevice{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	reply := &cntl.DMXDevice{}
	if err := controller.Get(req, &IDRequest{ID: key}, reply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	if reply.ID == "" {
		t.Errorf("exptected to get reply with an ID set, but isn't")
	}
}

func TestDmxDeviceController_Update_NotExisting(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	entity := ds.DMXDevices[key]
	reply := &cntl.DMXDevice{}

	if err := controller.Update(req, entity, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDmxDeviceController_Update_Existing(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	entity := ds.DMXDevices[key]

	createReply := &cntl.DMXDevice{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	reply := &cntl.DMXDevice{}
	if err := controller.Update(req, entity, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}
}
func TestDmxDeviceController_Delete_NotExisting(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	reply := &SuccessResponse{}

	if err := controller.Delete(req, &IDRequest{ID: key}, reply); err != errNotExists {
		t.Errorf("expected to get errNotExists, but got %v", err)
	}
}

func TestDmxDeviceController_Delete_Existing(t *testing.T) {
	logger := logrus.New().WithFields(logrus.Fields{})
	path := filepath.Join(os.TempDir(), "api_test", string(rand.Int63()))
	defer internalTesting.Cleanup(t, path)
	store := disk.New(path)
	controller := newDMXDeviceController(logger, store)
	req := httptest.NewRequest(http.MethodPost, rpcPath, nil)
	entity := ds.DMXDevices[key]

	createReply := &cntl.DMXDevice{}
	if err := controller.Create(req, entity, createReply); err != nil {
		t.Errorf("failed to call controller: %v", err)
	}

	reply := &SuccessResponse{}
	if err := controller.Delete(req, &IDRequest{ID: key}, reply); err != nil {
		t.Errorf("expected to get no error, but got %v", err)
	}

	if !reply.Success {
		t.Error("Expected to get result true, but got false")
	}
}
