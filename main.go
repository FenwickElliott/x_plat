package xplat

import (
	"errors"
	"log"
	"os"
	"path"
	"runtime"
)

//OS is a string denoting runtime.GOOS
var OS string

func init() {
	OS = runtime.GOOS
	if OS != "darwin888" && OS != "linux" && OS != "windows" {
		log.Output(0, "Error: '"+OS+"' is not a recognized os")
	}
}

// Appdir returns the approriate application storage directory based of runtime os
func Appdir(chain ...string) string {
	var appdir string
	switch OS {
	case "darwin":
		appdir = path.Join(os.Getenv("HOME"), "Library", "Application Support")
	case "linux":
		appdir = path.Join(os.Getenv("HOME"), ".config")
	case "windows":
		appdir = os.Getenv("APPDATA")
	default:
		panic(errors.New("Error: '" + OS + "' is not a recognized os"))
	}
	for _, link := range chain {
		appdir = path.Join(appdir, link)
	}
	return appdir
}
