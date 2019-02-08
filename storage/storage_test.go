package storage

import (
	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	ds = fixtures.DataStore()
	device = ds.DMXDevices[key]
	path = filepath.Join(os.TempDir(), "storage_test")
	key = "35cae00a-0b17-11e7-8bca-bbf30c56f20e"
	expectedFileName = filepath.Join(path, "DMXDevice", "DMXDevice_35cae00a-0b17-11e7-8bca-bbf30c56f20e.json")
	expectedContent = "{\"id\":\"35cae00a-0b17-11e7-8bca-bbf30c56f20e\",\"name\":\"LED-Bar below drums front\",\"typeId\":\"1555d67e-1187-11e7-8135-9b41038b5b75\",\"universe\":1,\"startChannel\":222,\"tags\":[\"bar\",\"drums\"]}"
)

func cleanup(t *testing.T, path string) {
	if err := os.RemoveAll(path); err != nil {
		t.Errorf("failed to remove test storage dir: %v", err)
	}
}

func TestStorage_buildFileName (t *testing.T) {
	storage := New(path)
	generated := storage.buildFileName(key, &cntl.DMXDevice{})
	expected := "DMXDevice_35cae00a-0b17-11e7-8bca-bbf30c56f20e.json"
	if generated != expected {
		t.Errorf("Expected go get fileKey %q, got %q", expected, generated)
	}
}

func TestStorage_getType (t *testing.T) {
	storage := New(path)

	name := storage.getType(&cntl.DMXDevice{})
	exp := "DMXDevice"
	if name != exp {
		t.Errorf("expected to get type name %q, got %q", exp, name)
	}
}

func TestStorage_Write(t *testing.T) {
	defer cleanup(t, path)
	storage := New(path)

	err := storage.Write(key, device)
	if err != nil {
		t.Error(err)
		return
	}

	if _, err := os.Stat(expectedFileName); err != nil {
		t.Errorf("expected file %s not found: %v", expectedFileName, err)
		return
	}

	b, err := ioutil.ReadFile(expectedFileName)
	if err != nil {
		t.Errorf("failed to read storage file %v: %v", expectedFileName, err)
		return
	}

	if string(b) != expectedContent {
		t.Errorf("Expected to find content %q, but got content %q", expectedContent, string(b))
	}
}

func TestStorage_Read(t *testing.T) {
	defer cleanup(t, path)
	storage := New(path)

	if err := os.MkdirAll(filepath.Dir(expectedFileName), 0755); err != nil {
		t.Fatalf("failed to prepare disk directory path: %v", err)
	}

	if err := ioutil.WriteFile(expectedFileName, []byte(expectedContent), 0755); err != nil {
		t.Fatalf("failed to prepare disk file: %v", err)
	}

	expDevice := &cntl.DMXDevice{}
	err := storage.Read(key, expDevice)
	if err != nil {
		t.Error(err)
		return
	}

	if expDevice.ID != key {
		t.Errorf("expected device to have id %q, but has %q", key, expDevice.ID)
		return
	}
}

func TestStorage_List(t *testing.T) {
	defer cleanup(t, path)
	storage := New(path)

	for k, dev := range ds.DMXDevices {
		if err := storage.Write(k, dev); err != nil {
			t.Error(err)
			return
		}
	}

	keys := storage.List(&cntl.DMXDevice{})
	if len(keys) != len(ds.DMXDevices) {
		t.Errorf("expected to get %d keys, got %d keys", len(ds.DMXDevices), len(keys))
	}
}

func TestStorage_Delete(t *testing.T) {
	defer cleanup(t, path)
	storage := New(path)

	if err := os.MkdirAll(filepath.Dir(expectedFileName), 0755); err != nil {
		t.Fatalf("failed to prepare disk directory path: %v", err)
	}

	if err := ioutil.WriteFile(expectedFileName, []byte(expectedContent), 0755); err != nil {
		t.Fatalf("failed to prepare disk file: %v", err)
	}

	err := storage.Delete(key, &cntl.DMXDevice{})
	if err != nil {
		t.Error(err)
		return
	}

	if _, err := os.Stat(expectedFileName); err != nil && !os.IsNotExist(err) {
		t.Error(err)
		return
	}
}
