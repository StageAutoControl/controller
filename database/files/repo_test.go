package files

import (
	"reflect"
	"testing"

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
		compare(t, key, sl, dsl, ok)
	}

	for key, s := range fix.Songs {
		ds, ok := data.Songs[key]
		compare(t, key, s, ds, ok)
	}

	for key, s := range fix.DmxScenes {
		ds, ok := data.DmxScenes[key]
		compare(t, key, s, ds, ok)
	}

	for key, p := range fix.DmxPresets {
		dp, ok := data.DmxPresets[key]
		compare(t, key, p, dp, ok)
	}

	for key, a := range fix.DmxAnimations {
		da, ok := data.DmxAnimations[key]
		compare(t, key, a, da, ok)
	}

	for key, dg := range fix.DmxDeviceGroups {
		ddg, ok := data.DmxDeviceGroups[key]
		compare(t, key, dg, ddg, ok)
	}

	for key, d := range fix.DmxDevices {
		dd, ok := data.DmxDevices[key]
		compare(t, key, d, dd, ok)
	}

}

func compare(t *testing.T, key string, expected, actual interface{}, ok bool) {
	if !ok {
		t.Fatalf("Cannot find key %q in given data object", key)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Given objects at key %q are not equal. Expected %#v, got %#v", key, expected, actual)
	}
}
