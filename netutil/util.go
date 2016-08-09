package netutil

import (
	"net"
	"strings"
)

func GetHost(addr string) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(addr)); err == nil {
		return ip
	}
	return ""
}

func GetIP(addr string) string {
	return GetHost(addr)
}

func GetPort(addr string) string {
	if _, port, err := net.SplitHostPort(strings.TrimSpace(addr)); err == nil {
		return port
	}
	return ""
}
