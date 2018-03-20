package artnet

import (
	"testing"
)

func TestState_Set(t *testing.T) {
	cases := []struct {
		before, after State
		u             uint16
		c, v          uint8
	}{
		{
			before: State(map[uint16][512]byte{}),
			after:  State(map[uint16][512]byte{12: {14: 16}}),
			u:      12,
			c:      14,
			v:      16,
		},
		{
			before: State(map[uint16][512]byte{12: {14: 16}}),
			after:  State(map[uint16][512]byte{12: {14: 16}, 2: {4: 6}}),
			u:      2,
			c:      4,
			v:      6,
		},
	}

	for i, c := range cases {
		c.before.Set(c.u, c.c, c.v)
		if !c.before.Equals(c.after) {
			t.Errorf("Expected to have state %+v at case %v, got %+v", c.after, i, c.before)
		}
	}
}
