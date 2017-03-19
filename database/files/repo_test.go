package files

import (
	"reflect"
	"testing"

	"fmt"

	"github.com/StageAutoControl/controller/fixtures"
)

func TestRepository_Load(t *testing.T) {
	dir := "./fixtures"
	fix := fixtures.DataStore()

	loader := New(dir)
	data, err := loader.Load()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	for key, sl := range fix.SetLists {
		dsl, ok := data.SetLists[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if dsl.ID != sl.ID {
			t.Errorf("ID %q is not equal \n", key)
		}
	}

	for key, s := range fix.Songs {
		ds, ok := data.Songs[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if ds.ID != s.ID {
			t.Errorf("ID %q is not equal \n", key)
		}
	}

	for key, s := range fix.DMXScenes {
		ds, ok := data.DMXScenes[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if ds.ID != s.ID {
			t.Errorf("ID %q is not equal \n", key)
		}
	}

	for key, p := range fix.DMXPresets {
		dp, ok := data.DMXPresets[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if dp.ID != p.ID {
			t.Errorf("ID %q is not equal \n", key)
		}
	}

	for key, a := range fix.DMXAnimations {
		da, ok := data.DMXAnimations[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if da.ID != a.ID {
			t.Errorf("ID %q is not equal \n", key)
		}
	}

	for key, dg := range fix.DMXDeviceGroups {
		ddg, ok := data.DMXDeviceGroups[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if ddg.ID != dg.ID {
			t.Errorf("ID %q is not equal \n", key)
		}
	}

	for key, d := range fix.DMXDevices {
		dd, ok := data.DMXDevices[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if dd.ID != d.ID {
			t.Errorf("ID %q is not equal \n", key)
		}
	}

}

func compare(t *testing.T, key string, expected, actual interface{}, ok bool) {
	if !ok {
		t.Fatalf("Cannot find key %q in given data object", key)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Given objects at key %q are not equal. Expected %#v, got %#v", key, expected, actual)
	}

	fmt.Printf("ID %q is equal \n", key)
}
