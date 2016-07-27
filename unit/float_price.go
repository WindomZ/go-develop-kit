package unit

import (
	"bytes"
	"database/sql/driver"
	"errors"
	. "github.com/WindomZ/go-develop-kit/math"
	"math"
	"strconv"
	"strings"
)

var (
	FloatPricePrecision int     = 8
	FloatPricePow       float64 = math.Pow10(FloatPricePrecision)
)

func SetFloatPricePrecision(e int) {
	FloatPricePrecision = e
	FloatPricePow = math.Pow10(FloatPricePrecision)
}

type FloatPrice float64

func NewFloatPrice(f float64, places ...int) FloatPrice {
	if f == 0 {
		return FloatPrice(0)
	} else if places != nil && len(places) != 0 {
		return FloatPrice(FloatFixed(f, places[0]))
	}
	return FloatPrice(FloatFixed(f, FloatPricePrecision))
}

func NewFloatPriceInt(i int64, places ...int) FloatPrice {
	return NewFloatPrice(float64(i)/FloatPricePow, places...)
}

func NewFloatPriceString(value string, places ...int) FloatPrice {
	p := NewFloatPrice(0)
	p.Scan(value)
	return p
}

func NewFloatPriceIntString(value string, places ...int) FloatPrice {
	if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		return NewFloatPriceInt(i)
	}
	return NewFloatPrice(0)
}

func (p *FloatPrice) MarshalJSON() ([]byte, error) {
	if p == nil {
		return nil, errors.New("MarshalJSON on nil pointer")
	}
	var b bytes.Buffer
	b.WriteByte('"')
	b.WriteString(p.StringFloat())
	b.WriteByte('"')
	return b.Bytes(), nil
}

func (p *FloatPrice) UnmarshalJSON(data []byte) error {
	if p == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	} else if f, err := strconv.ParseFloat(strings.Replace(string(data), `"`, ``, -1), 64); err != nil {
		return err
	} else {
		p.SetFloat64(f)
	}
	return nil
}

func (p FloatPrice) Value() (driver.Value, error) {
	return p.Float64(), nil
}

func (p *FloatPrice) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch o := src.(type) {
	case float32, float64:
		p.SetFloat64(o.(float64))
	case int, int8, int16, int32, int64:
		p.SetInt64(o.(int64))
	case string:
		f, err := strconv.ParseFloat(o, 64)
		if err != nil {
			return err
		}
		p.SetFloat64(f)
	case []byte:
		return p.Scan(string(o))
	default:
		return errors.New("Incompatible type for FloatPrice")
	}
	return nil
}

func (p FloatPrice) Int64() int64 {
	return FloatFixedToInt(p.Float64(), FloatPricePrecision)
}

func (p *FloatPrice) SetInt64(i int64, places ...int) *FloatPrice {
	return p.SetFloat64(float64(i)/FloatPricePow, places...)
}

func (p FloatPrice) Float64() float64 {
	return float64(p)
}

func (p FloatPrice) ReciprocalFloat64(places ...int) float64 {
	if places != nil && len(places) != 0 {
		return FloatFixed(1/p.Float64(), places[0])
	}
	return FloatFixed(1/p.Float64(), FloatPricePrecision)
}

func (p *FloatPrice) SetFloat64(f float64, places ...int) *FloatPrice {
	*p = NewFloatPrice(f, places...)
	return p
}

func (p FloatPrice) String() string {
	return p.StringFloat()
}

func (p FloatPrice) StringInt() string {
	return strconv.FormatInt(p.Int64(), 10)
}

func (p FloatPrice) StringFloat() string {
	return strconv.FormatFloat(p.Float64(), 'f', FloatPricePrecision, 64)
}

func (p *FloatPrice) IsPositive() bool {
	return p.Float64() > 0
}

func (p *FloatPrice) IsNegative() bool {
	return p.Float64() < 0
}

func (p *FloatPrice) Zero() {
	p.SetFloat64(0)
}

func (p *FloatPrice) IsZero() bool {
	return p.Float64() == 0
}

func (p *FloatPrice) Equal(f float64) bool {
	return p.Float64() == f
}

func (p *FloatPrice) LessThan(f float64) bool {
	return p.Float64() < f
}

func (p *FloatPrice) LessEqual(f float64) bool {
	return p.Float64() <= f
}

func (p *FloatPrice) GreaterThan(f float64) bool {
	return p.Float64() > f
}

func (p *FloatPrice) GreaterEqual(f float64) bool {
	return p.Float64() >= f
}

// rounded p+q and returns p
func (p *FloatPrice) Add(q FloatPrice) *FloatPrice {
	return p.SetFloat64(p.Float64() + q.Float64())
}

// rounded p-q and returns p
func (p *FloatPrice) Sub(q FloatPrice) *FloatPrice {
	return p.SetFloat64(p.Float64() - q.Float64())
}

// rounded product p*q and returns p
func (p *FloatPrice) Mul(q FloatPrice) *FloatPrice {
	return p.SetFloat64(p.Float64() * q.Float64())
}

// rounded quotient p/q and returns p
func (p *FloatPrice) Quo(q FloatPrice) *FloatPrice {
	return p.SetFloat64(p.Float64() / q.Float64())
}

// rounded p+x... and returns p
func (p *FloatPrice) Sum(x ...FloatPrice) *FloatPrice {
	for _, y := range x {
		p.Add(y)
	}
	return p
}

// rounded p-x... and returns p
func (p *FloatPrice) Diff(x ...FloatPrice) *FloatPrice {
	for _, y := range x {
		p.Sub(y)
	}
	return p
}

// rounded p+x... and returns
func (p FloatPrice) GetSum(x ...FloatPrice) FloatPrice {
	sum := p.Float64()
	for _, y := range x {
		sum += y.Float64()
	}
	return NewFloatPrice(sum)
}

// rounded p+x... and returns
func (p FloatPrice) GetDiff(x ...FloatPrice) FloatPrice {
	diff := p.Float64()
	for _, y := range x {
		diff -= y.Float64()
	}
	return NewFloatPrice(diff)
}

// returns negation
func (p FloatPrice) GetNegation() FloatPrice {
	return NewFloatPrice(-p.Float64())
}

// Cmp compares p and p and returns:
//
//   -1 if p <  p
//    0 if p == p (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
//   +1 if p >  p
//
func (p FloatPrice) Cmp(q FloatPrice) int {
	switch {
	case p.Float64() < q.Float64():
		return -1
	case p.Float64() > q.Float64():
		return +1
	}
	return 0
}
