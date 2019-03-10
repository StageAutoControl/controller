package artnet

import "github.com/jsimonetti/go-artnet"

// ToNet transforms the given net to an artnet.Address
func ToNet(net uint16) (add artnet.Address) {
	add.Net = uint8(net >> 8)
	add.SubUni = uint8(net)
	return
}
