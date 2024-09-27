package netx

import (
	"errors"
	"net"
)

var (
	ErrFailedToObtainHardwareAddr = errors.New("failed to obtain hardware address")
)

func GetHardwareAddr() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		var addr = iface.HardwareAddr.String()
		if len(addr) != 0 {
			return addr, nil
		}
	}
	return "", ErrFailedToObtainHardwareAddr
}

func GetHardwareAddrs() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var addrs []string
	for _, iface := range interfaces {
		var addr = iface.HardwareAddr.String()
		if len(addr) != 0 {
			addrs = append(addrs, addr)
		}
	}
	return addrs, nil
}
