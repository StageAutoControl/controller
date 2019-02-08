package dmx

import (
	"github.com/StageAutoControl/controller/pkg/cntl"
	"testing"

	"github.com/StageAutoControl/controller/pkg/internal/fixtures"
)

func TestRenderAnimation(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		d []*cntl.DMXDevice
		a *cntl.DMXAnimation
		c []cntl.DMXCommands
	}{
		{
			d: []*cntl.DMXDevice{
				ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			},
			a: ds.DMXAnimations["a51f7b2a-0e7b-11e7-bfc8-57da167865d7"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 228, Value: *fixtures.Value31},
				},
				{
					{Universe: 1, Channel: 228, Value: *fixtures.Value63},
				},
				{
					{Universe: 1, Channel: 228, Value: *fixtures.Value127},
				},
				{
					{Universe: 1, Channel: 228, Value: *fixtures.Value255},
				},
			},
		},
	}

	for i, e := range exp {
		c, err := RenderAnimation(ds, e.d, e.a)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(c) != len(e.c) {
			t.Fatalf("Expected to get length %d, got %d", len(e.c), len(c))
		}

		for j := range e.c {
			if len(e.c[j]) != len(c[j]) {
				t.Fatalf("Expected to get length %d at command index %d, got %d at index %d", len(e.c[j]), j, len(c[j]), i)
			}

			for _, cmd := range e.c[j] {
				if !c[j].Contains(cmd) {
					t.Errorf("Expected %+v to have %+v, but hasn't index %d", c[j], cmd, i)
				}
			}
		}
	}
}
