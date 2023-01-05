// Store folder structure:
// dnpm/store/[package name]/[version]/data/package/[package files]
// dnpm/store/[package name]/[version]/data/temp.tgz
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
	return os.MkdirAll(GetDefaultStorePath(), 0755)
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
	return os.MkdirAll(GetDefaultStorePath()+packageName+"/"+packageVersion+"/data/package/", 0755)
}

func DeletePackageStore(packageName string, packageVersion string) error {
	return os.RemoveAll(GetDefaultStorePath() + packageName + "/" + packageVersion)
}

func WriteTempFile(data []byte, packageName string, packageVersion string) (path string, err error) {
	CreatePackageStore(packageName, packageVersion)

	path = GetDefaultStorePath() + packageName + "/" + packageVersion + "/data/temp.tgz"
	return path, os.WriteFile(path, data, 0755)
}

func DeleteTempFile(packageName string, packageVersion string) error {
	return os.Remove(GetDefaultStorePath() + packageName + "/" + packageVersion + "/data/temp.tgz")
}
