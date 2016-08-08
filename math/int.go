package math

// Max returns the larger of x or y.
func MaxInt(x, y int) int {
	return int(MaxInt64(int64(x), int64(y)))
}

// Min returns the smaller of x or y.
func MinInt(x, y int) int {
	return int(MinInt64(int64(x), int64(y)))
}
