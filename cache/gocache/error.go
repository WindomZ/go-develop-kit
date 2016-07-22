package gocache

import "errors"

var (
	ErrNoKey       error = decorateErrorString("The key is empty")
	ErrNotExistKey error = decorateErrorString("The key is not existed")
	ErrNoValue           = decorateErrorString("The value is nil")
)

func decorateErrorString(str string) error {
	return errors.New("gocache: " + str)
}
