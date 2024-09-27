package gate

import (
	"encoding/binary"
	"fmt"
	"net"
)

func IPv4ToInt(IPv4Addr string) (uint32, error) {
	ip := net.ParseIP(IPv4Addr)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP address: %s", IPv4Addr)
	}

	ipv4 := ip.To4()
	if ipv4 == nil {
		return 0, fmt.Errorf("not an IPv4 address: %s", IPv4Addr)
	}

	return binary.BigEndian.Uint32(ipv4), nil
}
