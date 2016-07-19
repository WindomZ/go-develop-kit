package otp

import "github.com/WindomZ/go-develop-kit/googleauth"

type TOTP struct {
	OTP googleauth.OTPConfig
}

func NewTOTP(windowSize int) *TOTP {
	if windowSize <= 0 {
		windowSize = googleauth.DefaultWindowSize
	}
	return &TOTP{
		OTP: googleauth.OTPConfig{
			WindowSize: windowSize,
		},
	}
}

func (t *TOTP) normalize() *TOTP {
	if len(t.GetSecret()) == 0 {
		t.SetSecret("")
	}
	return t
}

func (t *TOTP) SetSecret(secret string) OTP {
	if len(secret) == 0 {
		t.OTP.Secret = GenerateSecret()
	} else {
		t.OTP.Secret = secret
	}
	return t
}

func (t *TOTP) GetSecret() string {
	if t == nil {
		return ""
	}
	return t.OTP.Secret
}

func (t *TOTP) ValidSecret() bool {
	if len(t.GetSecret()) == 0 {
		t.SetSecret("")
		return true
	}
	return ValidSecret(t.GetSecret())
}

func (t *TOTP) URL(user, issuer string) string {
	if t == nil {
		return ""
	}
	return t.normalize().OTP.ProvisionURIWithIssuer(user, issuer)
}

func (t *TOTP) Verify(password string) (bool, error) {
	if t == nil {
		return false, ErrNil
	}
	return t.normalize().OTP.Authenticate(password)
}
