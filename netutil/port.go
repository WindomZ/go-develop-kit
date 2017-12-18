package netutil

import "net"

// FreePort returns an address of TCP free port from the kernel.
func FreePort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return
	}
	defer l.Close()

	addr, ok := l.Addr().(*net.TCPAddr)
	if ok {
		port = addr.Port
	}
	return
}
