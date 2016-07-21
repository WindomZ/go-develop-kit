package freecache

import "errors"

var (
	ErrNoKey           error = decorateErrorString("The key is empty")
	ErrNoValue               = decorateErrorString("The value is nil")
	ErrNotSupportValue       = decorateErrorString("The value type is not supoort")
)

func decorateErrorString(str string) error {
	return errors.New("freecache: " + str)
}

func decorateError(err error) error {
	return decorateErrorString(err.Error())
}
