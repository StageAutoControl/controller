package files

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/internal/fixtures"
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

	for key, tr := range fix.DMXTransitions {
		dt, ok := data.DMXTransitions[key]
		if !ok {
			t.Fatalf("ID %q not found \n", key)
		}

		if dt.ID != tr.ID {
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
