package dmx

import (
	"reflect"
	"testing"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
)

func TestStreamlineScenes(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		s *cntl.Song
		m map[uint64]*cntl.DMXScene
	}{
		{
			s: ds.Songs["3c1065c8-0b14-11e7-96eb-5b134621c411"],
			m: map[uint64]*cntl.DMXScene{
				0:    ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"],
				32:   ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"],
				64:   ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"],
				96:   ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"],
				512:  ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"],
				528:  ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"],
				544:  ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"],
				1408: ds.DMXScenes["99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6"],
				1472: ds.DMXScenes["99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6"],
				1920: ds.DMXScenes["b82f4750-0e7a-11e7-9522-0f9d6d69958a"],
			},
		},
	}

	for i, e := range exp {
		res, err := StreamlineScenes(ds, e.s)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		for k, v := range e.m {
			resv, ok := res[k]
			if !ok {
				t.Errorf("Expected to have key %d at index %d", k, i)
				continue
			}

			if !reflect.DeepEqual(resv, v) {
				t.Errorf("Expected to get value %+v, got %+v", v, resv)
			}

			t.Logf("Found correct key %d at index %d \n", k, i)
		}

	}
}

func TestCalcSceneLength(t *testing.T) {
	exp := []struct {
		sc     *cntl.DMXScene
		length uint64
	}{
		{&cntl.DMXScene{NoteCount: 3, NoteValue: 4}, 24},
		{&cntl.DMXScene{NoteCount: 12, NoteValue: 8}, 48},
		{&cntl.DMXScene{NoteCount: 11, NoteValue: 4}, 88},
		{&cntl.DMXScene{NoteCount: 4, NoteValue: 4}, 32},
		{&cntl.DMXScene{NoteCount: 9, NoteValue: 8}, 36},
	}

	for i, e := range exp {
		if l := CalcSceneLength(e.sc); l != e.length {
			t.Errorf("Expected to get length %d, got %d for index %d", e.length, l, i)
		}
	}
}

func TestRenderScene(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		s   *cntl.DMXScene
		c   []cntl.DMXCommands
		err error
	}{
		{
			s: ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"],
			c: []cntl.DMXCommands{
				{{Universe: 1, Channel: 222, Value: 255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: 255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: 255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: 255}},
				{}, {}, {}, {}, {}, {}, {},
			},
			err: nil,
		},
		{
			s: ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"],
			c: []cntl.DMXCommands{
				{{Universe: 1, Channel: 224, Value: 255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 224, Value: 255}},
				{}, {}, {}, {}, {}, {}, {},
			},
			err: nil,
		},
		{
			s: ds.DMXScenes["99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6"],
			c: []cntl.DMXCommands{
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: 31}},
				{{Universe: 1, Channel: 228, Value: 63}},
				{{Universe: 1, Channel: 228, Value: 127}},
				{{Universe: 1, Channel: 228, Value: 255}},
				{}, {}, {}, {},
			},
			err: nil,
		},
	}

	for i, e := range exp {
		c, err := RenderScene(ds, e.s)

		t.Log(c)

		if e.err != nil && (err == nil || err.Error() != e.err.Error()) {
			t.Fatalf("Expected to get error %v, got %v at index %d", e.err, err, i)
		}

		if len(c) != len(e.c) {
			t.Errorf("Expected to get %d commands, got %d at index %d", len(e.c), len(c), i)
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
					{Universe: 1, Channel: 222, Value: 255},
				},
			},
			err: nil,
		},
		{
			p: ds.DMXPresets["11adf93e-0e7b-11e7-998c-5bd2bd0df396"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 224, Value: 255},
				},
			},
			err: nil,
		},
		{
			p: ds.DMXPresets["652e716a-0e7b-11e7-b92a-8f2ff28ba235"],
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 223, Value: 255},
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
			t.Errorf("Expected to get %d commands, got %d at index %d", len(e.c), len(c), i)
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

func TestMerge(t *testing.T) {

}
func TestMergeAtOffset(t *testing.T) {
	cmds := []cntl.DMXCommands{
		{{Universe: 0, Channel: 255, Value: 12}},
		{{Universe: 45, Channel: 200, Value: 15}},
		{{Universe: 12, Channel: 0, Value: 255}},
		{{Universe: 44, Channel: 55, Value: 66}},
		{{Universe: 41, Channel: 210, Value: 115}},
	}
	cs := []cntl.DMXCommands{
		{{Universe: 10, Channel: 15, Value: 1}},
		{{Universe: 11, Channel: 16, Value: 15}},
	}
	e := []cntl.DMXCommands{
		{{Universe: 0, Channel: 255, Value: 12}},
		{{Universe: 45, Channel: 200, Value: 15}},
		{{Universe: 12, Channel: 0, Value: 255}, {Universe: 10, Channel: 15, Value: 1}},
		{{Universe: 44, Channel: 55, Value: 66}, {Universe: 11, Channel: 16, Value: 15}},
		{{Universe: 41, Channel: 210, Value: 115}},
	}

	res := MergeAtOffset(cmds, cs, 2)
	for i, c := range e {
		if !c.Equals(res[i]) {
			t.Errorf("Expected %+v to equal %+v at index %d but doesn't.", c, res[i], i)
		}
	}
}

func TestRenderDeviceParams(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		dp  *cntl.DMXDeviceParams
		c   []cntl.DMXCommands
		err error
	}{
		{
			dp: &cntl.DMXDeviceParams{
				Device: &cntl.DMXDeviceSelector{ID: "5e0335e0-0b17-11e7-ad6c-63a7138d926c"},
				Params: &cntl.DMXParams{Red: 255, Green: 255, Blue: 255, LED: 0},
			},
			c: []cntl.DMXCommands{
				{
					{Universe: 2, Channel: 26, Value: 255},
					{Universe: 2, Channel: 27, Value: 255},
					{Universe: 2, Channel: 28, Value: 255},
				},
			},
			err: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Group:  &cntl.DMXDeviceGroupSelector{ID: "475b71a0-0b16-11e7-9406-e3f678e8b788"},
				Params: &cntl.DMXParams{Red: 255, Green: 255, Blue: 255, LED: 0},
			},
			c: []cntl.DMXCommands{
				{
					{Universe: 2, Channel: 10, Value: 255},
					{Universe: 2, Channel: 11, Value: 255},
					{Universe: 2, Channel: 12, Value: 255},
					{Universe: 2, Channel: 14, Value: 255},
					{Universe: 2, Channel: 15, Value: 255},
					{Universe: 2, Channel: 16, Value: 255},
				},
			},
			err: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Group:  &cntl.DMXDeviceGroupSelector{ID: "cb58bc10-0b16-11e7-b45a-7bee591b0adb"},
				Params: &cntl.DMXParams{Preset: 200},
			},
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 222, Value: 200},
				},
			},
			err: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:      &cntl.DMXDeviceSelector{ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e"},
				AnimationID: "a51f7b2a-0e7b-11e7-bfc8-57da167865d7",
			},
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 228, Value: 31},
				},
				{
					{Universe: 1, Channel: 228, Value: 63},
				},
				{
					{Universe: 1, Channel: 228, Value: 127},
				},
				{
					{Universe: 1, Channel: 228, Value: 255},
				},
			},
			err: nil,
		},
	}

	for i, e := range exp {
		c, err := RenderDeviceParams(ds, e.dp)
		if e.err != nil && (err == nil || err.Error() != e.err.Error()) {
			t.Fatalf("Expected to get error %v, got %v at index %d", e.err, err, i)
		}

		if len(c) != len(e.c) {
			t.Errorf("Expected to get %d commands, got %d at index %d", len(e.c), len(c), i)
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

func TestRenderParams(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		ds []*cntl.DMXDevice
		p  cntl.DMXParams
		c  int
	}{
		{
			ds: []*cntl.DMXDevice{
				ds.DMXDevices["4a545466-0b17-11e7-9c61-d3c0693099ab"],
			},
			p: cntl.DMXParams{Red: 255},
			c: 1,
		},
		{
			ds: []*cntl.DMXDevice{
				ds.DMXDevices["4a545466-0b17-11e7-9c61-d3c0693099ab"],
				ds.DMXDevices["5e0335e0-0b17-11e7-ad6c-63a7138d926c"],
			},
			p: cntl.DMXParams{Red: 255, Green: 255, Blue: 255},
			c: 6,
		},
	}

	for _, e := range exp {
		c, err := RenderParams(ds, e.ds, e.p)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if len(c) != e.c {
			t.Errorf("Expected to get %d commands, got %d.", e.c, len(c))
		}
	}
}

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
					{Universe: 1, Channel: 228, Value: 31},
				},
				{
					{Universe: 1, Channel: 228, Value: 63},
				},
				{
					{Universe: 1, Channel: 228, Value: 127},
				},
				{
					{Universe: 1, Channel: 228, Value: 255},
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

func TestResolveDeviceSelectorByID(t *testing.T) {
	ds := fixtures.DataStore()
	sel := &cntl.DMXDeviceSelector{
		ID: "4a545466-0b17-11e7-9c61-d3c0693099ab",
	}

	dd, err := ResolveDeviceSelector(ds, sel)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(dd) != 1 {
		t.Errorf("Expected to get 1 devices, got %d", len(dd))
	}

	if dd[0].ID != sel.ID {
		t.Errorf("Expected to get device %q, got %q.", sel.ID, dd[0].ID)
	}
}

func TestResolveDeviceSelectorByTags(t *testing.T) {
	ds := fixtures.DataStore()
	sel := &cntl.DMXDeviceSelector{
		Tags: []cntl.Tag{
			"inner",
			"drums-left",
		},
	}

	dd, err := ResolveDeviceSelector(ds, sel)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(dd) != 2 {
		t.Errorf("Expected to get 2 devices, got %d", len(dd))
	}
}

func TestResolveDevicesByTags(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		count int
		tags  []cntl.Tag
	}{
		{0, []cntl.Tag{"par", "inner", "stand-left"}},
		{2, []cntl.Tag{"par", "inner", "drums-left"}},
		{4, []cntl.Tag{"par"}},
		{1, []cntl.Tag{"strobe-back"}},
	}

	for _, e := range exp {
		res := ResolveDevicesByTags(ds, e.tags)
		if len(res) != e.count {
			t.Errorf("Expected to get %d devices for tags %s, got %d", e.count, e.tags, len(res))
		}
	}
}

func TestResolveDevicesByTag(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		c int
		t cntl.Tag
	}{
		{1, cntl.Tag("bar")},
		{4, cntl.Tag("par")},
		{2, cntl.Tag("right")},
		{2, cntl.Tag("drums-left")},
		{1, cntl.Tag("vocs")},
	}

	for _, e := range exp {
		d := ResolveDevicesByTag(ds, e.t)

		if len(d) != e.c {
			t.Errorf("Expected to get %d devices, got %d", e.c, len(d))
		}
	}
}

func TestHas(t *testing.T) {
	ds := []*cntl.DMXDevice{
		{ID: "1"},
		{ID: "2"},
	}

	exp := []struct {
		d   *cntl.DMXDevice
		has bool
	}{
		{&cntl.DMXDevice{ID: "0"}, false},
		{&cntl.DMXDevice{ID: "1"}, true},
		{&cntl.DMXDevice{ID: "2"}, true},
		{&cntl.DMXDevice{ID: "3"}, false},
	}

	for _, e := range exp {
		ok := has(ds, e.d)
		if ok != e.has {
			t.Errorf("Expected to get %s for ID %q, got %s", e.has, e.d.ID, ok)
		}
	}
}
