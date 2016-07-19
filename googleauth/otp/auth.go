package otp

import "sync"

const (
	TypeTOTP string = "totp"
	TypeHOTP        = "hotp"
)

func ValidType(t string) bool {
	return (t == TypeTOTP || t == TypeHOTP)
}

type Authenticator struct {
	Type     string
	mux_auth *sync.RWMutex
	OTPAuth  map[string]OTP
	mux_open *sync.RWMutex
	OTPOpen  map[string]bool
}

func NewAuthenticator(_type string) (*Authenticator, error) {
	if !ValidType(_type) {
		return nil, ErrType
	}
	return &Authenticator{
		Type:     _type,
		mux_auth: new(sync.RWMutex),
		OTPAuth:  make(map[string]OTP),
		mux_open: new(sync.RWMutex),
		OTPOpen:  make(map[string]bool),
	}, nil
}

func (a *Authenticator) AddSecret(id, secret string) (OTP, error) {
	if len(id) == 0 {
		return nil, ErrID
	}
	a.mux_auth.Lock()
	defer a.mux_auth.Unlock()
	if v, ok := a.OTPAuth[id]; ok {
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

func (a *Authenticator) Open(id string) {
	a.mux_open.Lock()
	defer a.mux_open.Unlock()
	if _, ok := a.OTPOpen[id]; ok {
		delete(a.OTPOpen, id)
	}
}

func (a *Authenticator) Close(id string) {
	a.mux_open.Lock()
	defer a.mux_open.Unlock()
	a.OTPOpen[id] = false
}

func (a *Authenticator) IsOpen(id string) bool {
	a.mux_open.RLock()
	defer a.mux_open.RUnlock()
	if v, ok := a.OTPOpen[id]; ok {
		return v
	}
	return true
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
