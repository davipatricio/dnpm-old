package store

import (
	"encoding/json"
	"os"
)

type CachedPackageData struct {
	SuccessfullDownload bool   `json:"successfull_download"`
	SuccessfullExtract  bool   `json:"successfull_extract"`
}

// GetCachedPackageData returns the cached package data for the given package and version
func GetCachedPackageData(packageName string, version string) (data CachedPackageData, err error) {
	file, err := os.Open(GetDefaultStorePath() + packageName + "/" + version + "/data/data.json")
	if err != nil {
		return
	}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return
	}

	return data, file.Close()
}

// SetCachedPackageData sets the cached package data for the given package and version
func SetCachedPackageData(packageName string, version string, data CachedPackageData) (err error) {
	CreatePackageStore(packageName, version)
	file, err := os.Create(GetDefaultStorePath() + packageName + "/" + version + "/data/data.json")
	if err != nil {
		return
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return
	}

	return file.Close()
}
