package utils

import (
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

	return path
}
