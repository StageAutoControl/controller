package artnet

import "github.com/jsimonetti/go-artnet"

func ToNet(net uint16) (add artnet.Address) {
	add.Net = uint8(net >> 8)
	add.SubUni = uint8(net)
	return
}
