package artnet

import (
	"fmt"
	"net"
	"strings"
)

const (
	// addressRange specifies the network CIDR an artnet network should have
	addressRange = "2.0.0.0/8"
)

// FindArtNetIP finds the matching interface with an IP address inside of the addressRange
func FindArtNetIP() (net.IP, error) {
	var ip net.IP

	_, cidrnet, _ := net.ParseCIDR(addressRange)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip, fmt.Errorf("error getting ips: %s", err)
	}

	for _, addr := range addrs {
		ip = addr.(*net.IPNet).IP

		if strings.Contains(ip.String(), ":") {
			continue
		}

		if cidrnet.Contains(ip) {
			break
		}
	}

	return ip, nil
}
