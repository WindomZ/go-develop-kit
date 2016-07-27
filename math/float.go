package math

import (
	"github.com/shopspring/decimal"
	"math"
)

func FloatPrecision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}

func FloatRound(val float64, prec int) float64 {
	pow := math.Pow10(prec)
	digit := pow * val
	if _, div := math.Modf(digit); div >= .5 {
		return math.Ceil(digit) / pow
	}
	return math.Floor(digit)
}

func FloatFixed(value float64, places int) float64 {
	f, _ := decimal.NewFromFloat(value).Round(int32(places)).Float64()
	return f
}

func FloatFixedToInt(value float64, places int) int64 {
	return decimal.NewFromFloat(value).Mul(decimal.New(1, int32(places))).Round(0).IntPart()
}
