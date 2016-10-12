package netutil

import "testing"

func TestGetHost(t *testing.T) {
	if ip := GetHost("127.0.0.1:1234"); ip != "127.0.0.1" {
		t.Fatal("Fail:", ip)
	}
}

func TestGetPort(t *testing.T) {
	if port := GetPort("127.0.0.1:1234"); port != "1234" {
		t.Fatal("Fail:", port)
	}
}
