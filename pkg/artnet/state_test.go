package artnet

import (
	"reflect"
	"testing"
)

func TestState_Set(t *testing.T) {
	cases := []struct {
		before, after *State
		u             uint16
		c, v          uint8
	}{
		{
			before: NewStateFromData(map[uint16]Universe{}),
			after:  NewStateFromData(map[uint16]Universe{12: {14: 16}}),
			u:      12,
			c:      14,
			v:      16,
		},
		{
			before: NewStateFromData(map[uint16]Universe{12: {14: 16}}),
			after:  NewStateFromData(map[uint16]Universe{12: {14: 16}, 2: {4: 6}}),
			u:      2,
			c:      4,
			v:      6,
		},
	}

	for i, c := range cases {
		c.before.SetChannel(c.u, c.c, c.v)

		bu := c.before.GetUniverses()
		au := c.after.GetUniverses()

		if !reflect.DeepEqual(bu, au) {
			t.Errorf("Expected to get universes %+v at case %v, got %+v", au, c, bu)
			continue
		}

		for _, u := range au {
			if !reflect.DeepEqual(c.before.GetUniverse(u), c.after.GetUniverse(u)) {
				t.Errorf("Expected to have state %+v at case %v, got %+v", c.after.GetUniverse(u), i, c.before.GetUniverse(u))
			}
		}

	}
}
