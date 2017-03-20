package uuid

import "strings"

// ToUUID returns lowercase string of id.
func ToUUID(id string) string {
	return strings.ToLower(strings.TrimSpace(id))
}

// NoDashUUID returns no dash('-') id.
func NoDashUUID(id string) string {
	return strings.Replace(ToUUID(id), "-", "", -1)
}
