package otp

import (
	"testing"
	"time"
)

func TestTOTP_URL(t *testing.T) {
	if true {
		return
	}
	o := NewTOTP(3, 16)
	t.Logf("url: %v", o.URL("totp", "https://github.com/WindomZ/go-develop-kit/tree/master/googleauth"))
	t.Logf("secret: %v", o.GetSecret())
	if !o.ValidSecret() {
		t.Fatal("Fail to verify secret")
	}
}

func TestTOTP_Verify(t *testing.T) {
	if true {
		return
	}
	password := "766930"
	if len(password) == 0 {
		return
	}
	o := NewTOTP(3, 6).
		SetSecret("52FF7CCEWHNLESTV")
	start := time.Now()
	for {
		ok, err := o.Verify(password)
		if err != nil {
			t.Error(err)
			break
		} else if !ok {
			t.Error("Fail to verify")
			break
		}
		time.Sleep(time.Second)
	}
	t.Logf("wait seconds: %v", time.Now().Sub(start).Seconds())
}
