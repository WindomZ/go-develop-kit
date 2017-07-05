package path

import (
	"bytes"
	"fmt"
	"os"
)

// CreateFile creates the path file with mode 0666 (before umask), truncating it if it already exists.
// If there is an error, it will be of type *PathError.
func CreateFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return f.Close()
}

// AppendToFile appends the contents of s to the path file, each s ends with new line.
func AppendToFile(path string, s ...string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	sb := bytes.Buffer{}
	for _, str := range s {
		sb.WriteString(str)
		sb.WriteByte('\n')
	}
	_, err = f.Write(sb.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// OverwriteFile overwrite the path file with the contents of s, each s ends with a new line.
func OverwriteFile(path string, s ...string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, str := range s {
		if _, err = f.WriteString(fmt.Sprintln(str)); err != nil {
			return err
		}
	}
	return nil
}

// ReadFile reads the contents from the path file.
func ReadFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	sb := bytes.Buffer{}
	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		sb.Write(buf[:n])
	}
	return sb.String(), nil
}

// RemoveFile if dir then removes path and any children it contains,
// otherwise then removes the path file or directory.
func RemoveFile(path string, dir bool) error {
	if dir {
		return os.RemoveAll(path)
	}
	return os.Remove(path)
}
