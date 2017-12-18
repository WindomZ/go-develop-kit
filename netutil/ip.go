package netutil

import (
	"net"
	"strings"
	"sync"
)

const EmptyIP = "0.0.0.0"

var (
	once        sync.Once
	localIP     string
	localIPTail string
)

// ip returns the current IP address.
func ip() (string, error) {
	if localIP != "" {
		return localIP, nil
	}
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return EmptyIP, err
	}
	for _, address := range addresses {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				localIP = ipNet.IP.String()
				return localIP, nil
			}
		}
	}
	return EmptyIP, nil
}

// IP returns the current IP address.
func IP() string {
	if localIP != "" {
		return localIP
	}
	once.Do(func() {
		localIP, _ = ip()
		localIPTail = localIP[strings.LastIndexByte(localIP, '.'):]
	})
	return localIP
}

// IPTail returns the end of the current IP address.
// like: 4 in 192.168.1.4
func IPTail() string {
	return localIPTail
}
