package dmx

import (
	"testing"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
)

func TestRepeat(t *testing.T) {
	exp := []struct {
		count uint
		cmds  []cntl.DMXCommands
		res   []cntl.DMXCommands
	}{
		{
			count: 4,
			cmds: []cntl.DMXCommands{
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			},
			res: []cntl.DMXCommands{
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			},
		},
	}

	for i, e := range exp {
		res := repeat(e.count, e.cmds)

		if len(res) != len(e.res) {
			t.Fatalf("Expected to get %d commands, got %d at case index %d", len(e.res), len(res), i)
		}

		for j := range e.res {
			if len(e.res[j]) != len(res[j]) {
				t.Fatalf("Expected to get length %d at command index %d, got %d at case index %d", len(e.res[j]), j, len(res[j]), i)
			}

			for _, cmd := range e.res[j] {
				if !res[j].Contains(cmd) {
					t.Errorf("Expected %+v to have %+v, but hasn't index %d", res[j], cmd, i)
				}
			}
		}
	}
}


func TestResize(t *testing.T) {
	exp := []struct {
		length uint
		cmds  []cntl.DMXCommands
		res   []cntl.DMXCommands
	}{
		{
			length: 10,
			cmds: []cntl.DMXCommands{
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
			},
			res: []cntl.DMXCommands{
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {},
			},
		},
		{
			length: 10,
			cmds: []cntl.DMXCommands{
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {},
				{}, {}, {}, {}, {}, {}, {}, {}, {},
			},
			res: []cntl.DMXCommands{
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {}, {}, {},
				{}, {}, {}, {}, {}, {}, {}, {}, {},
			},
		},
	}

	for i, e := range exp {
		res := resize(e.length, e.cmds)
		t.Logf("case %d result %+v", i, res)

		if len(res) != len(e.res) {
			t.Fatalf("Expected to get %d commands, got %d at case index %d", len(e.res), len(res), i)
		}

		for j := range e.res {
			if len(e.res[j]) != len(res[j]) {
				t.Fatalf("Expected to get length %d at command index %d, got %d at case index %d", len(e.res[j]), j, len(res[j]), i)
			}

			for _, cmd := range e.res[j] {
				if !res[j].Contains(cmd) {
					t.Errorf("Expected %+v to have %+v, but hasn't index %d", res[j], cmd, i)
				}
			}
		}
	}
}
