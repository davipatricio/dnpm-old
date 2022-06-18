package utils

import (
	"os"
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

	createEmptyStoreFolder(path)
	return path
}

func PkgAlreadyCached(name, version string) bool {
	storeDir := GetStoreDir()
	_, err := os.Stat(storeDir + "/" + name + "/" + version)
	return err == nil
}

func createEmptyStoreFolder(dir string) {
	// Verify if the store folder exists
	_, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
