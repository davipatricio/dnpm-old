package store

import (
	"encoding/json"
	"os"
)

type CachedPackageData struct {
	SuccessfullDownload bool   `json:"successfull_download"`
	Shasum              string `json:"shasum"`
}

func GetCachedPackageData(packageName string, version string) (data CachedPackageData, err error) {
	file, err := os.Open(GetDefaultStorePath() + packageName + "/" + version + "/data/data.json")
	if err != nil {
		return 
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return
	}

	return
}

func SetCachedPackageData(packageName string, version string, data CachedPackageData) (err error) {
	os.MkdirAll(GetDefaultStorePath()+packageName+"/"+version+"/data/files/", 0755)
	file, err := os.Create(GetDefaultStorePath() + packageName + "/" + version + "/data/data.json")
	if err != nil {
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return
	}

	return
}