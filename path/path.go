package path

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// IsExist returns boolean indicating whether the path of file or directory already exists.
// Returns an unknown error if not match some syscall errors.
func IsExist(_path string) (ok bool, err error) {
	_, err = os.Stat(_path)
	if ok = err == nil; !ok && os.IsNotExist(err) {
		err = nil
	}
	return
}

// MustExist returns a boolean, it true if the path of file or directory already exists.
func MustExist(_path string) bool {
	ok, _ := IsExist(_path)
	return ok
}

// Ensure returns an error if the path of file or directory(dir is true) isn't existed or catches other error.
func Ensure(_path string, dir bool) error {
	if ok, _ := IsExist(_path); !ok {
		if dir {
			if err := os.MkdirAll(_path, os.ModePerm); err != nil {
				return err
			}
		} else if err := Ensure(filepath.Dir(_path), true); err != nil {
			return err
		} else if err := CreateFile(_path); err != nil {
			return err
		}
	}
	return nil
}

// ExecPath returns the path of the command-line program.
func ExecPath() string {
	if len(os.Args) != 0 {
		return os.Args[0]
	}
	return ""
}

// ExecDir returns the directory path of the command-line program.
func ExecDir() string {
	file, _ := exec.LookPath(ExecPath())
	filePath, _ := filepath.Abs(file)
	index := strings.LastIndex(filePath, string(os.PathSeparator))
	ret := filePath[:index]
	return ret
}
