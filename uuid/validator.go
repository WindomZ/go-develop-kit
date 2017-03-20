package uuid

import (
	"regexp"

	"github.com/satori/go.uuid"
)

var EMPTY, _ = uuid.FromString("00000000-0000-0000-0000-000000000000")

// IsEmpty returns true if is empty UUID.
func IsEmpty(id string) bool {
	u, err := uuid.FromString(id)
	if err != nil {
		return false
	}
	return uuid.Equal(EMPTY, u)
}

// Valid returns true if valid UUID and not equal empty.
func Valid(id string) (result bool) {
	if IsEmpty(id) {
	} else if len(id) == 36 {
		result, _ = regexp.MatchString(`\w{8}(-\w{4}){3}-\w{12}`, id)
	} else if len(id) == 32 {
		result = ValidNoDash(id)
	}
	return
}

// ValidNoDash returns true if valid UUID that no dash and not equal empty.
func ValidNoDash(id string) (result bool) {
	result, _ = regexp.MatchString(`\w{8}(\w{4}){3}\w{12}`, id)
	return
}
