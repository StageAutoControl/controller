package dmx

import (
	"errors"
	"github.com/StageAutoControl/controller/pkg/cntl"
	"testing"

	"github.com/StageAutoControl/controller/pkg/internal/fixtures"
)

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
			t.Errorf("Expected to get %v for ID %q, got %v", e.has, e.d.ID, ok)
		}
	}
}

func TestGetDeviceChannel(t *testing.T) {
	ds := fixtures.DataStore()
	exp := []struct {
		d   *cntl.DMXDevice
		c   cntl.DMXChannel
		led uint16
		res cntl.DMXChannel
		err error
	}{
		{
			d:   ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			c:   ChannelRed,
			led: 3,
			res: 234,
			err: nil,
		},
		{
			d:   ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			c:   ChannelGreen,
			led: 0,
			res: 223,
			err: nil,
		},
		{
			d:   ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			c:   ChannelBlue,
			led: 2,
			res: 232,
			err: nil,
		},
		{
			d:   ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			c:   ChannelStrobe,
			led: 0,
			res: 224,
			err: nil,
		},
		{
			d:   ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			c:   ChannelDimmer,
			led: 16,
			res: 0,
			err: errors.New("given device has insufficient biggest index of LEDs 15 to handle the given LED index 16"),
		},
		{
			d:   ds.DMXDevices["35cae00a-0b17-11e7-8bca-bbf30c56f20e"],
			c:   ChannelDimmer,
			led: 15,
			res: 223,
			err: nil,
		},
		{
			d:   ds.DMXDevices["s429fc37c-0b17-11e7-8b94-c3b6519355d3"],
			c:   ChannelMode,
			led: 0,
			res: 0,
			err: ErrDeviceHasDisabledModeChannel,
		},
	}

	for i, e := range exp {
		res, err := getDeviceChannel(ds, e.d, e.c, e.led)
		if e.err != nil && (err == nil || err.Error() != e.err.Error()) {
			t.Fatalf("Expected to get error %v, got %v", e.err, err)
		}

		if res != e.res {
			t.Errorf("Expected to get res %d, got %d at index %d", e.res, res, i)
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
