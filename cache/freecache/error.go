package freecache

import "errors"

var (
	ErrNoKey   error = decorateErrorString("The key is empty")
	ErrNoValue       = decorateErrorString("The value is nil")
)

func decorateErrorString(str string) error {
	return errors.New("freecache: " + str)
}

func decorateError(err error) error {
	if err != nil {
		return decorateErrorString(err.Error())
	}
	return nil
}
