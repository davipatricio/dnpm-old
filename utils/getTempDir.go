package utils

import (
	"os"
	"os/user"
	"runtime"
)

func GetTempDir() string {
	path := ""
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "windows":
		path = user.HomeDir + "/AppData/Local/dnpm/temp"
	case "darwin":
		path = user.HomeDir + "/Library/dnpm/temp"
	case "linux":
		path = user.HomeDir + "/.local/share/pnpm/temp"
	default:
		path = user.HomeDir + "/.dnpm/temp"
	}

	createEmptyTempFolder(path)
	return path
}

func createEmptyTempFolder(dir string) {
	// Verify if the store folder exists
	_, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
