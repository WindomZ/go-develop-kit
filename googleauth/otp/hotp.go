package otp

import "github.com/WindomZ/go-develop-kit/googleauth"

type HOTP struct {
	OTP          googleauth.OTPConfig
	SecretLength int
}

func NewHOTP(windowSize, secretLength int) *HOTP {
	if windowSize <= 0 {
		windowSize = googleauth.DefaultWindowSize
	}
	if secretLength <= 0 {
		secretLength = googleauth.DefaultRandomSecretLength
	}
	return &HOTP{
		OTP: googleauth.OTPConfig{
			WindowSize:  windowSize,
			HotpCounter: 1,
		},
		SecretLength: secretLength,
	}
}

func (t *HOTP) normalize() *HOTP {
	if len(t.GetSecret()) == 0 {
		t.SetSecret("")
	}
	return t
}

func (t *HOTP) SetSecret(secret string) OTP {
	if len(secret) == 0 {
		t.OTP.Secret = GenerateSecret(t.SecretLength)
	} else {
		t.OTP.Secret = secret
	}
	return t
}

func (t *HOTP) GetSecret() string {
	if t == nil {
		return ""
	}
	return t.OTP.Secret
}

func (t *HOTP) ValidSecret() bool {
	if len(t.GetSecret()) == 0 {
		t.SetSecret("")
		return true
	}
	return ValidSecret(t.GetSecret(), t.SecretLength)
}

func (t *HOTP) URL(user, issuer string) string {
	if t == nil {
		return ""
	}
	return t.normalize().OTP.ProvisionURIWithIssuer(user, issuer)
}

func (t *HOTP) Verify(password string) (bool, error) {
	if t == nil {
		return false, ErrNil
	}
	return t.normalize().OTP.Authenticate(password)
}
