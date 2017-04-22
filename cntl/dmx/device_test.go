package dmx

import (
	"errors"
	"testing"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
)

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
			err: errors.New("Given device has insufficient biggest index of LEDs 15 to handle the given LED index 16"),
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
			err: errDeviceHasDisabledModeChannel,
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
