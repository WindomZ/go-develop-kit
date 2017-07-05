package path

import (
	"os"
	"path"
)

// IsExist returns a boolean indicating whether the path of file or directory already exists.
// Returns an unknown error if not match some syscall errors.
func IsExist(_path string) (ok bool, err error) {
	_, err = os.Stat(_path)
	if ok = err == nil; !ok && os.IsNotExist(err) {
		err = nil
	}
	return
}

// Ensure returns an error if the path of file or directory(dir is true) isn't existed or catches other error.
func Ensure(_path string, dir bool) error {
	if ok, _ := IsExist(_path); !ok {
		if dir {
			if err := os.MkdirAll(_path, os.ModePerm); err != nil {
				return err
			}
		} else if err := Ensure(path.Dir(_path), true); err != nil {
			return err
		} else {
			f, err := os.Create(_path)
			if err != nil {
				return err
			}
			f.Close()
		}
	}
	return nil
}
