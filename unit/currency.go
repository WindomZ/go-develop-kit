package unit

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"strings"
	"sync"
)

var (
	mux               *sync.RWMutex     = new(sync.RWMutex)
	currencyMapping   map[string]string = make(map[string]string)
	currencyUnMapping map[string]string = make(map[string]string)
)

func SetCurrencyMapping(s, mapping string) {
	if len(s) != 0 && len(mapping) != 0 && s != mapping {
		mux.Lock()
		currencyMapping[s] = mapping
		currencyUnMapping[mapping] = s
		mux.Unlock()
	}
}

func CurrencyMapping(s string) string {
	if len(s) != 0 {
		mux.RLock()
		mapping, ok := currencyMapping[s]
		mux.RUnlock()
		if ok {
			return mapping
		}
	}
	return s
}

func CurrencyUnMapping(mapping string) string {
	if len(mapping) != 0 {
		mux.RLock()
		s, ok := currencyUnMapping[mapping]
		mux.RUnlock()
		if ok {
			return s
		}
	}
	return mapping
}

type Currency string

func NewCurrency(s string) Currency {
	return Currency(CurrencyMapping(s))
}

func (c *Currency) MarshalJSON() ([]byte, error) {
	if c == nil {
		return nil, errors.New("MarshalJSON on nil pointer")
	}
	var b bytes.Buffer
	b.WriteByte('"')
	b.WriteString(CurrencyUnMapping(c.String()))
	b.WriteByte('"')
	return b.Bytes(), nil
}

func (c *Currency) UnmarshalJSON(data []byte) error {
	if c == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	}
	c.SetString(strings.Replace(string(data), `"`, ``, -1))
	return nil
}

func (c Currency) Value() (driver.Value, error) {
	return c.String(), nil
}

func (c *Currency) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch o := src.(type) {
	case string:
		c.SetString(o)
	case []byte:
		c.SetString(string(o))
	default:
		return errors.New("Incompatible type for Currency")
	}
	return nil
}

func (c Currency) String() string {
	return string(c)
}

func (c *Currency) SetString(s string) *Currency {
	*c = NewCurrency(s)
	return c
}

func (c *Currency) Equal(s *Currency) bool {
	return c.String() == s.String()
}
