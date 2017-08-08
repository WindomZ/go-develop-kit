package bytes

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "0B", Format(0))
	assert.Equal(t, "1023B", Format(1023))
	assert.Equal(t, "1.00KB", Format(1024))
	assert.Equal(t, "30.59KB", Format(31324))
	assert.Equal(t, "120.56KB", Format(123456))
	assert.Equal(t, "11.77MB", Format(12345678))
	assert.Equal(t, "1.15GB", Format(1234567890))
	assert.Equal(t, "1.12TB", Format(1234567890123))
	assert.Equal(t, "1.10PB", Format(1234567890123456))
	assert.Equal(t, "109.65PB", Format(123456789012345678))

	assert.Equal(t, "1234567890123456789B", Format(1234567890123456789))
}

func TestParse(t *testing.T) {
	assert.Equal(t, uint64(0), MustParse("0B"))
	assert.Equal(t, uint64(1023), MustParse("1023B"))
	assert.Equal(t, uint64(1024), MustParse("1KB"))
	assert.Equal(t, uint64(1024), MustParse("1.0KB"))
	assert.Equal(t, uint64(1024), MustParse("1.00KB"))
	assert.Equal(t, uint64(31324), MustParse("30.59KB"))
	assert.Equal(t, uint64(1048576), MustParse("1.00MB"))
	assert.Equal(t, uint64(1234803097), MustParse("1.15GB"))
	assert.Equal(t, uint64(6442450944), MustParse("6G"))
	assert.Equal(t, uint64(1352399302164), MustParse("1.23TB"))
	assert.Equal(t, uint64(1384856885416427), MustParse("1.23PB"))

	var err error
	_, err = Parse("0")
	assert.NotEmpty(t, err)
	_, err = Parse("1..23KB")
	assert.NotEmpty(t, err)
	_, err = Parse("1.23EB")
	assert.NotEmpty(t, err)

	defer func() {
		assert.NotEmpty(t, recover())
	}()
	MustParse("1..23EB")
}
