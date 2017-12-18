package netutil

import (
	"regexp"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestIP(t *testing.T) {
	assert.NotEqual(t, EmptyIP, IP())
	assert.Regexp(t, regexp.MustCompile(`\d{2,3}\.\d{2,3}\.\d{1,3}\.\d{1,3}`), IP())
}

func TestIPTail(t *testing.T) {
	assert.NotEqual(t, "0", IPTail())
	assert.Regexp(t, regexp.MustCompile(`\d{1,3}`), IPTail())
}
