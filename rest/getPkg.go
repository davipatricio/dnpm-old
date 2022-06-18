package rest

import (
	"encoding/json"
	"net/http"
)

const RegistryURL string = "https://registry.yarnpkg.com"

func GetPkg(pkgName string) (map[string]interface{}, error) {
	url := RegistryURL + "/" + pkgName
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
