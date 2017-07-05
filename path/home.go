package path

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
)

// HomeDir returns the path to the user's home directory (if they have one).
func HomeDir() (string, error) {
	if u, err := user.Current(); err == nil {
		return u.HomeDir, nil
	} else if strings.ToLower(runtime.GOOS) == "windows" {
		// windows system
		return homeDirWindows()
	}
	// Unix-like system, so just assume Unix
	return homeDirUnix()
}

func homeDirUnix() (home string, err error) {
	// First prefer the HOME environmental variable
	if home = os.Getenv("HOME"); len(home) != 0 {
		return
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err = cmd.Run(); err != nil {
		return
	}

	home = strings.TrimSpace(stdout.String())
	if len(home) == 0 {
		err = errors.New("The blank is output when reading the home directory")
	}
	return
}

func homeDirWindows() (home string, err error) {
	homeDrive := os.Getenv("HOMEDRIVE")
	homePath := os.Getenv("HOMEPATH")
	if len(homeDrive) == 0 || len(homePath) == 0 {
		home = os.Getenv("USERPROFILE")
	} else {
		home = path.Join(homeDrive, homePath)
	}
	if len(home) == 0 {
		err = errors.New("Environmental HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	return
}
