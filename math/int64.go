package math

import "github.com/shopspring/decimal"

// Max returns the larger of x or y.
func MaxInt64(x, y int64) int64 {
	return decimal.Max(decimal.New(x, 0), decimal.New(y, 0)).IntPart()
}

// Min returns the smaller of x or y.
func MinInt64(x, y int64) int64 {
	return decimal.Min(decimal.New(x, 0), decimal.New(y, 0)).IntPart()
}
