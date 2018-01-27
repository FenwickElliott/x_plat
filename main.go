package xplat

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

//OS is a string denoting runtime.GOOS
var OS string

func init() {
	OS = runtime.GOOS
	if OS != "darwin" && OS != "linux" && OS != "windows" {
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

// Openbrowser opens the default broswer of the host system to given url
func Openbrowser(url string) {
	switch OS {
	case "darwin":
		fmt.Println("here")
		exec.Command("open", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		panic(fmt.Errorf("Error: '" + OS + "' is not a recognized os"))
	}
}
