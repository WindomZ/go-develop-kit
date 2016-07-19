package otp

import (
	"testing"
	"time"
)

func TestTOTP_URL(t *testing.T) {
	if true {
		return
	}
	o := NewTOTP(3)
	t.Logf("url: %v", o.URL("totp", "https://github.com/WindomZ/go-develop-kit/tree/master/googleauth"))
	t.Logf("secret: %v", o.GetSecret())
}

func TestTOTP_Verify(t *testing.T) {
	if true {
		return
	}
	password := "109478"
	if len(password) == 0 {
		return
	}
	o := NewTOTP(3).
		SetSecret("OZGGQ5DZGZMXQM2FIZUFK3D2NB2HAVDCLE2XQUSKKVIDGZKBGBME4YKSGVDGOWRRKBXTA43QIVUTA2DNLBGTA5TVINRWGY2VNFLWMSTUOJWESUCTKZVEG2CLNNQXGTRTGJWHOR2JGRBXSNCJPB4TGSTZMJCHS4JY")
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
