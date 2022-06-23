package utils

import (
	"dnpm/structs"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

func GetPackageJSON() structs.PackageJSONFormat {
	dir, _, err := GetNearestPackageJSON()

	jsonFile, _ := ioutil.ReadFile(dir)

	if err != nil {
		panic(err)
	}

	var pkgJSON structs.PackageJSONFormat
	json.Unmarshal(jsonFile, &pkgJSON)
	return pkgJSON
}

func GetPackageJSONOfPackage(pkgName string) (structs.PackageJSONFormat, bool) {
	pkgDir := path.Join(GetExecDir(), "/node_modules", pkgName)

	var pkgJSON structs.PackageJSONFormat
	if _, err := os.Stat(pkgDir); os.IsNotExist(err) {
		return pkgJSON, false
	}

	jsonFile, _ := ioutil.ReadFile(pkgDir + "/package.json")
	json.Unmarshal(jsonFile, &pkgJSON)

	return pkgJSON, true
}
