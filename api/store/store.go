// Store folder structure:
// dnpm/store/[package name]/[version]/data/files/[package files]
// dnpm/store/[package name]/[version]/data/data.json
package store

import (
	"os"
	"os/user"
	"runtime"
)

func DefaultStoreExists() bool {
	_, err := os.Stat(GetDefaultStorePath())
	return err == nil
}

func CreateDefaultStore() error {
	err := os.MkdirAll(GetDefaultStorePath(), 0755)
	if err != nil {
		return err
	}

	return nil
}


func GetDefaultStorePath() string {
	path := ""
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "windows":
		path = user.HomeDir + "/AppData/Local/dnpm/store/"
	case "darwin":
		path = user.HomeDir + "/Library/dnpm/store/"
	case "linux":
		path = user.HomeDir + "/.local/share/pnpm/store/"
	default:
		path = user.HomeDir + "/.dnpm/store/"
	}

	return path
}

func CreatePackageStore(packageName string, packageVersion string) error {
	err := os.MkdirAll(GetDefaultStorePath()+packageName+"/"+packageVersion+"/data/files/", 0755)
	if err != nil {
		return err
	}

	return nil
}
