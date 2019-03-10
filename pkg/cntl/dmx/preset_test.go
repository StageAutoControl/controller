package dmx

import (
	"testing"

	"github.com/StageAutoControl/controller/pkg/cntl"

	"github.com/StageAutoControl/controller/pkg/internal/fixtures"
)

func TestRenderPreset(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		p   *cntl.DMXPreset
		c   []cntl.DMXCommands
		err error
	}{
		{
			p: ds.DMXPresets["0de258e0-0e7b-11e7-afd4-ebf6036983dc"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 222, Value: *fixtures.Value255},
				},
			},
			err: nil,
		},
		{
			p: ds.DMXPresets["11adf93e-0e7b-11e7-998c-5bd2bd0df396"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 224, Value: *fixtures.Value255},
				},
			},
			err: nil,
		},
		{
			p: ds.DMXPresets["652e716a-0e7b-11e7-b92a-8f2ff28ba235"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 223, Value: *fixtures.Value255},
				},
			},
			err: nil,
		},
	}

	for i, e := range exp {
		c, err := RenderPreset(ds, e.p)
		if e.err != nil && (err == nil || err.Error() != e.err.Error()) {
			t.Fatalf("Expected to get error %v, got %v at index %d", e.err, err, i)
		}

		if len(c) != len(e.c) {
			t.Fatalf("Expected to get %d commands, got %d at index %d", len(e.c), len(c), i)
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
