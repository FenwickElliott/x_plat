package xplat

import (
	"errors"
	"os"
	"path"
	"runtime"
)

// Appdir returns the approriate application storage directory based of runtime os
func Appdir() (string, error) {
	var appdir string
	switch runtime.GOOS {
	case "darwin":
		appdir = path.Join(os.Getenv("HOME"), "Library", "Application Support")
	case "linux":
		appdir = path.Join(os.Getenv("HOME"), ".config")
	case "windows":
		appdir = os.Getenv("APPDATA")
	default:
		return "", errors.New("Unrecogniszed os: " + runtime.GOOS)
	}
	return appdir, nil
}
