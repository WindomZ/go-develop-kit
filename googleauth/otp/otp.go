package otp

type OTP interface {
	SetSecret(secret string) OTP
	GetSecret() string
	ValidSecret() bool
	URL(user, issuer string) string
	Verify(password string) (bool, error)
}
