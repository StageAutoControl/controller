package dmx

import (
	"fmt"
	"testing"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
)

func TestCheckDeviceParams(t *testing.T) {
	testCases := []struct {
		dp          *cntl.DMXDeviceParams
		expectedErr error
	}{
		{
			dp: &cntl.DMXDeviceParams{
				Device:     &cntl.DMXDeviceSelector{ID: "asdf"},
				Group:      &cntl.DMXDeviceGroupSelector{ID: "asdf2"},
				Transition: &cntl.TransitionSelector{ID: "anim1"},
				Animation:  &cntl.AnimationSelector{ID: "anim2"},
			},
			expectedErr: ErrDeviceParamsDevicesInvalid,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:     &cntl.DMXDeviceSelector{ID: "asdf"},
				Group:      &cntl.DMXDeviceGroupSelector{ID: "asdf2"},
				Transition: &cntl.TransitionSelector{ID: "anim1"},
				Animation:  &cntl.AnimationSelector{ID: "anim2"},
			},
			expectedErr: ErrDeviceParamsDevicesInvalid,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:     &cntl.DMXDeviceSelector{ID: "asdf"},
				Transition: &cntl.TransitionSelector{ID: "anim1"},
				Animation:  &cntl.AnimationSelector{ID: "anim2"},
			},
			expectedErr: ErrDeviceParamsValuesInvalid,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:     &cntl.DMXDeviceSelector{ID: "asdf"},
				Transition: &cntl.TransitionSelector{ID: "anim1"},
				Params: []cntl.DMXParams{
					{Blue: fixtures.Value255},
				},
			},
			expectedErr: ErrDeviceParamsValuesInvalid,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:    &cntl.DMXDeviceSelector{ID: "asdf"},
				Animation: &cntl.AnimationSelector{ID: "anim1"},
				Params: []cntl.DMXParams{
					{Blue: fixtures.Value255},},
			},
			expectedErr: ErrDeviceParamsValuesInvalid,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device: &cntl.DMXDeviceSelector{ID: "asdf"},
				Params: []cntl.DMXParams{
					{Blue: fixtures.Value255},
				},
			},
			expectedErr: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:    &cntl.DMXDeviceSelector{ID: "asdf"},
				Animation: &cntl.AnimationSelector{ID: "anim1"},
			},
			expectedErr: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:     &cntl.DMXDeviceSelector{ID: "asdf"},
				Transition: &cntl.TransitionSelector{ID: "anim1"},
			},
			expectedErr: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Group: &cntl.DMXDeviceGroupSelector{ID: "asdf"},
				Params: []cntl.DMXParams{
					{Blue: fixtures.Value255},
				},
			},
			expectedErr: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Group:     &cntl.DMXDeviceGroupSelector{ID: "asdf"},
				Animation: &cntl.AnimationSelector{ID: "anim1"},
			},
			expectedErr: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Group:      &cntl.DMXDeviceGroupSelector{ID: "asdf"},
				Transition: &cntl.TransitionSelector{ID: "anim1"},
			},
			expectedErr: nil,
		},
	}

	for index, testCase := range testCases {
		err := checkDeviceParams(testCase.dp)
		if err != testCase.expectedErr {
			t.Errorf("index %d: expected to get error %v, but got %v", index, testCase.expectedErr, err)
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
				Params: []cntl.DMXParams{
					{
						Red:   fixtures.Value255,
						Green: fixtures.Value255,
						Blue:  fixtures.Value255,
						LED:   0,
					},
				},
			},
			c: []cntl.DMXCommands{
				{
					{Universe: 2, Channel: 26, Value: *fixtures.Value255},
					{Universe: 2, Channel: 27, Value: *fixtures.Value255},
					{Universe: 2, Channel: 28, Value: *fixtures.Value255},
				},
			},
			err: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Group: &cntl.DMXDeviceGroupSelector{ID: "475b71a0-0b16-11e7-9406-e3f678e8b788"},
				Params: []cntl.DMXParams{
					{
						Red:   fixtures.Value255,
						Green: fixtures.Value255,
						Blue:  fixtures.Value255,
						LED:   0,
					},
				},
			},
			c: []cntl.DMXCommands{
				{
					{Universe: 2, Channel: 10, Value: *fixtures.Value255},
					{Universe: 2, Channel: 11, Value: *fixtures.Value255},
					{Universe: 2, Channel: 12, Value: *fixtures.Value255},
					{Universe: 2, Channel: 14, Value: *fixtures.Value255},
					{Universe: 2, Channel: 15, Value: *fixtures.Value255},
					{Universe: 2, Channel: 16, Value: *fixtures.Value255},
				},
			},
			err: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Group: &cntl.DMXDeviceGroupSelector{ID: "cb58bc10-0b16-11e7-b45a-7bee591b0adb"},
				Params: []cntl.DMXParams{
					{Mode: fixtures.Value200},
				},
			},
			c: []cntl.DMXCommands{
				{
					{Universe: 1, Channel: 222, Value: *fixtures.Value200},
				},
			},
			err: nil,
		},
		{
			dp: &cntl.DMXDeviceParams{
				Device:    &cntl.DMXDeviceSelector{ID: "35cae00a-0b17-11e7-8bca-bbf30c56f20e"},
				Animation: &cntl.AnimationSelector{ID: "a51f7b2a-0e7b-11e7-bfc8-57da167865d7"},
			},
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
			err: nil,
		},
	}

	for i, e := range exp {
		c, err := RenderDeviceParams(ds, e.dp)
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

func TestRenderParams(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		ds    []*cntl.DMXDevice
		p     cntl.DMXParams
		count int
		cmds  cntl.DMXCommands
	}{
		{
			ds: []*cntl.DMXDevice{
				ds.DMXDevices["4a545466-0b17-11e7-9c61-d3c0693099ab"],
			},
			p: cntl.DMXParams{
				Red: fixtures.Value255,
			},
			count: 1,
			cmds: cntl.DMXCommands{
				{Universe: 2, Channel: 14, Value: cntl.DMXValue{Value: 255}},
			},
		},
		{
			ds: []*cntl.DMXDevice{
				ds.DMXDevices["4a545466-0b17-11e7-9c61-d3c0693099ab"],
				ds.DMXDevices["5e0335e0-0b17-11e7-ad6c-63a7138d926c"],
			},
			p: cntl.DMXParams{
				Red:   fixtures.Value255,
				Green: fixtures.Value255,
				Blue:  fixtures.Value255,
			},
			count: 6,
			cmds: cntl.DMXCommands{
				{Universe: 2, Channel: 14, Value: cntl.DMXValue{Value: 255}},
				{Universe: 2, Channel: 15, Value: cntl.DMXValue{Value: 255}},
				{Universe: 2, Channel: 16, Value: cntl.DMXValue{Value: 255}},
				{Universe: 2, Channel: 26, Value: cntl.DMXValue{Value: 255}},
				{Universe: 2, Channel: 27, Value: cntl.DMXValue{Value: 255}},
				{Universe: 2, Channel: 28, Value: cntl.DMXValue{Value: 255}},
			},
		},
	}

	for _, e := range exp {
		c, err := RenderParams(ds, e.ds, e.p)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if len(c) != e.count {
			t.Errorf("Expected to get %d commands, got %d.", e.count, len(c))
		}

		fmt.Printf("%+v\n", c)

		for i, cmd := range e.cmds {
			if !c.Contains(cmd) {
				t.Errorf("Expected command list to contain %v at index %v, but doesn't.", cmd, i)
			}
		}
	}
}
