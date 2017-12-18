package iputil

import (
	"fmt"
	"net"
	"testing"
)

func TestFreePort(t *testing.T) {
	port, err := FreePort()
	if err != nil {
		t.Fatal(err)
	}

	if l, err := net.Listen("tcp",
		fmt.Sprintf("localhost:%d", port)); err != nil {
		t.Error(err)
	} else {
		l.Close()
	}
}
