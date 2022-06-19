package utils

import (
	"os/user"
	"runtime"
)

func GetStoreDir() string {
	path := ""
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "windows":
		path = user.HomeDir + "/AppData/Local/dnpm/store"
	case "darwin":
		path = user.HomeDir + "/Library/dnpm/store"
	case "linux":
		path = user.HomeDir + "/.local/share/pnpm/store"
	default:
		path = user.HomeDir + "/.dnpm/store"
	}

	return path
}
