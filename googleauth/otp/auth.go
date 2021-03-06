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
	Type         string
	SecretLength int
	mux_auth     *sync.RWMutex
	OTPAuth      map[string]OTP
	mux_open     *sync.RWMutex
	OTPOpen      map[string]bool
	mux_active   *sync.RWMutex
	OTPActive    map[string]bool
}

func NewAuthenticator(_type string, secretLength int) (*Authenticator, error) {
	if !ValidType(_type) {
		return nil, ErrType
	}
	return &Authenticator{
		Type:         _type,
		SecretLength: secretLength,
		mux_auth:     new(sync.RWMutex),
		OTPAuth:      make(map[string]OTP),
		mux_open:     new(sync.RWMutex),
		OTPOpen:      make(map[string]bool),
		mux_active:   new(sync.RWMutex),
		OTPActive:    make(map[string]bool),
	}, nil
}

func (a *Authenticator) AddSecret(id, secret string, update bool) (OTP, error) {
	if len(id) == 0 {
		return nil, ErrID
	}
	a.mux_auth.Lock()
	defer a.mux_auth.Unlock()
	if update {
	} else if v, ok := a.OTPAuth[id]; ok {
		return v, ErrExist
	}
	switch a.Type {
	case TypeTOTP:
		a.OTPAuth[id] = NewTOTP(3, a.SecretLength).SetSecret(secret)
	case TypeHOTP:
		a.OTPAuth[id] = NewHOTP(5, a.SecretLength).SetSecret(secret)
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
	a.Inactive(id)
}

func (a *Authenticator) IsOpen(id string) bool {
	a.mux_open.RLock()
	defer a.mux_open.RUnlock()
	if v, ok := a.OTPOpen[id]; ok {
		return v
	}
	return true
}

func (a *Authenticator) Active(id string) {
	a.mux_active.Lock()
	defer a.mux_active.Unlock()
	if _, ok := a.OTPActive[id]; ok {
		delete(a.OTPActive, id)
	}
}

func (a *Authenticator) Inactive(id string) {
	a.mux_active.Lock()
	defer a.mux_active.Unlock()
	a.OTPActive[id] = false
}

func (a *Authenticator) IsActive(id string) bool {
	a.mux_active.RLock()
	defer a.mux_active.RUnlock()
	if v, ok := a.OTPActive[id]; ok {
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
	if !a.IsOpen(id) {
		return false, ErrNotOpen
	} else if v, ok := a.OTPAuth[id]; ok {
		if ok, err := v.Verify(password); !ok || err != nil {
			return false, err
		} else if !a.IsActive(id) {
			a.Active(id)
		}
		return true, nil
	}
	return false, ErrNotExist
}
