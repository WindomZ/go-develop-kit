package otp

import (
	"testing"
	"time"
)

func TestHOTP_URL(t *testing.T) {
	if true {
		return
	}
	o := NewHOTP(5, 16)
	t.Logf("url: %v", o.URL("hotp", "https://github.com/WindomZ/go-develop-kit/tree/master/googleauth"))
	t.Logf("secret: %v", o.GetSecret())
}

func TestHOTP_Verify(t *testing.T) {
	if true {
		return
	}
	password := "973335"
	if len(password) == 0 {
		return
	}
	// TODO: fix it!
	o := NewHOTP(5, 16).
		SetSecret("GNWEONCNHE2XMQ3MMVBVSNCLPJ4GE4CYGR2UUYTULJRTOUTJIJ3TKZ2JNNUWUODIGFAUK5RUJZVHSZ2BG5DWGRDONNIXK3DUJBHVMM2UOBBVI5KZNQ4FOSLYMY3TA23SMNRTGTJZNJLDAZSOIMZXGMRQJ5EEU2DR")
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
