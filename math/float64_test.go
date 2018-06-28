package math

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestFloatRound(t *testing.T) {
	assert.Equal(t, FloatRound(2.012345, 5), 2.01235)
	assert.Equal(t, FloatRound(4.012345, 5), 4.01235)
}

func TestFloatRoundToInt(t *testing.T) {
	assert.Equal(t, FloatRoundToInt(2.012345, 5), int64(201235))
	assert.Equal(t, FloatRoundToInt(4.012345, 5), int64(401235))
}

func TestFloatSumRound(t *testing.T) {
	const F1 = 1.012345
	const F2 = 2.012345
	const F3 = 3.012345
	const F4 = 4.012345
	const F5 = 5.012345
	assert.Equal(t, FloatSum(F1, F2), 3.02469)
	assert.Equal(t, FloatSumRound(F1, F2, 5, F3, F4, F5), 15.06173)
}

func TestFloatSubRound(t *testing.T) {
	const F1 = 1.012345
	const F2 = 2.012345
	const F3 = 3.012345
	const F4 = 4.012345
	const F5 = 5.012345
	assert.Equal(t, FloatSub(F1, F2), float64(-1))
	assert.Equal(t, FloatSubRound(F1, F2, 5, F3, F4, F5), -13.03704)
}

func TestFloatMulRound(t *testing.T) {
	const F1 = 1.012345
	const F2 = 2.012345
	const F3 = 3.012345
	const F4 = 4.012345
	const F5 = 5.012345
	assert.Equal(t, FloatMul(F1, F2), 2.037187399025)
	assert.Equal(t, FloatMulRound(F1, F2, 5, F3, F4, F5), 123.41698)
}

func TestFloatDivRound(t *testing.T) {
	const F1 = 1.012345
	const F2 = 2.012345
	const F3 = 3.012345
	const F4 = 4.012345
	const F5 = 5.012345
	assert.Equal(t, FloatDiv(F1, F2), 0.5030673169859045)
	assert.Equal(t, FloatDivRound(F1, F2, 5, F3, F4, F5), 0.0083)
}
