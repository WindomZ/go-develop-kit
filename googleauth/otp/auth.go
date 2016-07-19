package otp

const (
	TypeTOTP int16 = 1
	TypeHOTP       = 2
)

func ValidType(t int16) bool {
	return (t == TypeTOTP || t == TypeHOTP)
}

type Authenticator struct {
	Type    int16
	OTPAuth map[string]OTP
}

func NewAuthenticator(_type int16) (*Authenticator, error) {
	if !ValidType(_type) {
		return nil, ErrType
	}
	return &Authenticator{
		Type:    _type,
		OTPAuth: make(map[string]OTP),
	}, nil
}

func (a *Authenticator) AddSecret(id, secret string) (OTP, error) {
	if len(id) == 0 {
		return nil, ErrID
	} else if v, ok := a.OTPAuth[id]; ok {
		return v, ErrExist
	}
	switch a.Type {
	case TypeTOTP:
		a.OTPAuth[id] = NewTOTP(3).SetSecret(secret)
	case TypeHOTP:
		a.OTPAuth[id] = NewHOTP(5).SetSecret(secret)
	default:
		return nil, ErrType
	}
	return a.OTPAuth[id], nil
}

func (a *Authenticator) GetSecret(id string) string {
	if v, ok := a.OTPAuth[id]; ok {
		return v.GetSecret()
	}
	return ""
}

func (a *Authenticator) ValidSecret(id string) bool {
	if v, ok := a.OTPAuth[id]; ok {
		return v.ValidSecret()
	}
	return false
}

func (a *Authenticator) URL(id, user, issuer string) string {
	if v, ok := a.OTPAuth[id]; ok {
		return v.URL(user, issuer)
	}
	return ""
}

func (a *Authenticator) Verify(id, password string) (bool, error) {
	if v, ok := a.OTPAuth[id]; ok {
		return v.Verify(password)
	}
	return false, ErrNotExist
}
