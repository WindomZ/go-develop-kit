package math

import "github.com/WindomZ/decimal"

func FloatRound(value float64, places int) float64 {
	f, _ := decimal.NewFromFloat(value).Round(int32(places)).Float64()
	return f
}

func FloatRoundToInt(value float64, places int) int64 {
	return decimal.NewFromFloat(value).Mul(decimal.New(1, int32(places))).Round(0).IntPart()
}

func FloatSum(x, y float64) float64 {
	f, _ := decimal.NewFromFloat(x).Add(decimal.NewFromFloat(y)).Float64()
	return f
}

func FloatSumRound(x, y float64, places int, zs ...float64) float64 {
	d := decimal.NewFromFloat(x).Add(decimal.NewFromFloat(y))
	if zs != nil && len(zs) != 0 {
		for _, z := range zs {
			d = d.Add(decimal.NewFromFloat(z))
		}
	}
	f, _ := d.Round(int32(places)).Float64()
	return f
}

func FloatSub(x, y float64) float64 {
	f, _ := decimal.NewFromFloat(x).Sub(decimal.NewFromFloat(y)).Float64()
	return f
}

func FloatSubRound(x, y float64, places int, zs ...float64) float64 {
	d := decimal.NewFromFloat(x).Sub(decimal.NewFromFloat(y))
	if zs != nil && len(zs) != 0 {
		for _, z := range zs {
			d = d.Sub(decimal.NewFromFloat(z))
		}
	}
	f, _ := d.Round(int32(places)).Float64()
	return f
}

func FloatMul(x, y float64) float64 {
	f, _ := decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y)).Float64()
	return f
}

func FloatMulRound(x, y float64, places int, zs ...float64) float64 {
	d := decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y))
	if zs != nil && len(zs) != 0 {
		for _, z := range zs {
			d = d.Mul(decimal.NewFromFloat(z))
		}
	}
	f, _ := d.Round(int32(places)).Float64()
	return f
}

func FloatDiv(x, y float64) float64 {
	f, _ := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y)).Float64()
	return f
}

func FloatDivRound(x, y float64, places int, zs ...float64) float64 {
	d := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y))
	if zs != nil && len(zs) != 0 {
		for _, z := range zs {
			d = d.Div(decimal.NewFromFloat(z))
		}
	}
	f, _ := d.Round(int32(places)).Float64()
	return f
}
