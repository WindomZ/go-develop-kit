package uuid

import (
	"regexp"

	"github.com/google/uuid"
)

var EMPTY, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")

// IsEmpty returns true if is empty UUID.
func IsEmpty(id string) bool {
	u, err := uuid.Parse(id)
	if err != nil {
		return id == "00000000000000000000000000000000"
	}
	return u.String() == EMPTY.String()
}

// Valid returns true if valid UUID.
func Valid(id string) (result bool) {
	if len(id) == 36 {
		result, _ = regexp.MatchString(`\w{8}(-\w{4}){3}-\w{12}`, id)
	} else if len(id) == 32 {
		result = ValidNoDash(id)
	}
	return
}

// ValidNoEmpty returns true if valid UUID and not equal empty.
func ValidNoEmpty(id string) bool {
	return !IsEmpty(id) && Valid(id)
}

// ValidNoDash returns true if valid UUID that no dash and not equal empty.
func ValidNoDash(id string) (result bool) {
	result, _ = regexp.MatchString(`\w{8}(\w{4}){3}\w{12}`, id)
	return
}
