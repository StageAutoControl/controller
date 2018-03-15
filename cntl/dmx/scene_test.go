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
		m map[uint64][]*cntl.DMXScene
	}{
		{
			s: ds.Songs["3c1065c8-0b14-11e7-96eb-5b134621c411"],
			m: map[uint64][]*cntl.DMXScene{
				0:    {ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"]},
				32:   {ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"]},
				64:   {ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"]},
				96:   {ds.DMXScenes["492cef2e-0b14-11e7-be89-c3fa25f9cabb"]},
				512:  {ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"]},
				528:  {ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"]},
				544:  {ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"]},
				1408: {ds.DMXScenes["99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6"]},
				1472: {ds.DMXScenes["99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6"]},
				1920: {ds.DMXScenes["b82f4750-0e7a-11e7-9522-0f9d6d69958a"]},
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
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 222, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {},
			},
			err: nil,
		},
		{
			s: ds.DMXScenes["a44f8dee-0b14-11e7-b5b9-bf1015384192"],
			c: []cntl.DMXCommands{
				{{Universe: 1, Channel: 224, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {},
				{{Universe: 1, Channel: 224, Value: *fixtures.Value255}},
				{}, {}, {}, {}, {}, {}, {},
			},
			err: nil,
		},
		{
			s: ds.DMXScenes["99b86a5e-0e7a-11e7-a01a-5b5fbdeba3d6"],
			c: []cntl.DMXCommands{
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
				{}, {}, {}, {},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value31}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value63}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value127}},
				{{Universe: 1, Channel: 228, Value: *fixtures.Value255}},
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
