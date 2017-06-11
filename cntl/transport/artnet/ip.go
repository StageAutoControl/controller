package artnet

import (
	"fmt"
	"net"
	"strings"
)

const (
	// ArtNetCIDR specifies the network CIDR an artnet network should have
	ArtNetCIDR = "2.0.0.0/8"
)

// FindArtNetIP finds the matching interface with an IP address inside of the ArtNetCIDR
func FindArtNetIP() (net.IP, error) {
	var ip net.IP

	_, cidrnet, _ := net.ParseCIDR(ArtNetCIDR)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip, fmt.Errorf("error getting ips: %s\n", err)
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
