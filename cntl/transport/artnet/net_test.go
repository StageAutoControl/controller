package artnet

import (
	"testing"

	"github.com/jsimonetti/go-artnet"
)

func TestToNet(t *testing.T) {
	cases := []struct{
		net uint16
		artnet artnet.Address
	}{
		{net: 0, artnet: artnet.Address{Net: 0, SubUni: 0}},
		{net: 15, artnet: artnet.Address{Net: 0, SubUni: 15}},
		{net: 256, artnet: artnet.Address{Net: 1, SubUni: 0}},
		{net: 33333, artnet: artnet.Address{Net: 130, SubUni: 53}},
		{net: 65535, artnet: artnet.Address{Net: 255, SubUni: 255}},
	}

	for i, testCase := range cases {
		result := ToNet(testCase.net)

		if result.SubUni != testCase.artnet.SubUni || result.Net != testCase.artnet.Net {
			t.Errorf("Expceted case %d to return %+v, got %+v", i, testCase.artnet, result)
		}
	}
}
