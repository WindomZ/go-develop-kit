package netutil

import (
	"net"
	"net/http"
	"strings"
)

// GetHost returns ip from addr that real client IP like '127.0.0.1:80'
func GetHost(addr string) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(addr)); err == nil {
		return ip
	}
	return ""
}

// GetIP same as GetHost
func GetIP(addr string) string {
	return GetHost(addr)
}

// GetHost returns port from addr that real client IP like '127.0.0.1:80'
func GetPort(addr string) string {
	if _, port, err := net.SplitHostPort(strings.TrimSpace(addr)); err == nil {
		return port
	}
	return ""
}

// ClientIP implements a best effort algorithm to return the real client IP, it parses
// X-Real-IP and X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
func ClientIP(r *http.Request) string {
	if r == nil {
		return ""
	}
	var clientIP string
	if values, _ := r.Header["X-Real-Ip"]; len(values) > 0 {
		clientIP = strings.TrimSpace(values[0])
	}
	if clientIP != "" {
		return clientIP
	}
	if values, _ := r.Header["X-Forwarded-For"]; len(values) > 0 {
		clientIP = strings.TrimSpace(values[0])
	}
	if index := strings.IndexByte(clientIP, ','); index >= 0 {
		clientIP = clientIP[0:index]
	}
	if clientIP = strings.TrimSpace(clientIP); clientIP != "" {
		return clientIP
	}
	return GetIP(r.RemoteAddr)
}
