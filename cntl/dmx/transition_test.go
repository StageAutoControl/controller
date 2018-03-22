package dmx

import (
	"log"
	"testing"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
)

func TestRenderTransition(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		d []*cntl.DMXDevice
		t *cntl.DMXTransition
		c []cntl.DMXCommands
	}{
		{
			d: []*cntl.DMXDevice{
				ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			},
			t: ds.DMXTransitions["a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 0}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 0}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 10}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 10}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 41}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 41}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 93}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 93}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 161}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 161}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 213}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 213}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 244}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 244}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 255}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 255}},
				},
			},
		},
		{
			d: []*cntl.DMXDevice{
				ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			},
			t: ds.DMXTransitions["525eaa7e-fb2d-4608-b413-559d284b3c85"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 255}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 255}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 244}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 244}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 213}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 213}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 161}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 161}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 93}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 93}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 41}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 41}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 10}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 10}},
				},
				{
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 0}},
					{Universe: 1, Channel: 232, Value: cntl.DMXValue{Value: 0}},
				},
			},
		},
		{
			d: []*cntl.DMXDevice{
				ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			},
			t: ds.DMXTransitions["e683873b-20da-4fd4-ac62-271925c68047"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 255}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 0}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 31}},
				},
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 244}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 10}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 34}},
				},
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 213}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 41}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 46}},
				},
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 161}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 93}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 66}},
				},
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 93}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 161}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 91}},
				},
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 41}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 213}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 111}},
				},
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 10}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 244}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 123}},
				},
				{
					{Universe: 1, Channel: 226, Value: cntl.DMXValue{Value: 0}},
					{Universe: 1, Channel: 227, Value: cntl.DMXValue{Value: 255}},
					{Universe: 1, Channel: 228, Value: cntl.DMXValue{Value: 127}},
				},
			},
		},
	}

	for i, e := range exp {
		c, err := RenderTransition(ds, e.d, e.t)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		log.Printf("%+v\n", c)

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
