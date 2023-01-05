package downloader

import (
	"fmt"
	"io"
	"net/http"

	"github.com/davipatricio/dnpm/api/integrity"
	"github.com/davipatricio/dnpm/api/store"
	"github.com/davipatricio/dnpm/api/tgz"
)

// DownloadPackage downloads a package from a url, saves it to the store and extracts it
// The first boolean value is true if the package was already downloaded and false if it was downloaded now
// To skip the integrity check, pass an empty string as the shasum
func DownloadAndSavePackage(url string, shasum, packageName, version string) (bool, error) {
	d, err := store.GetCachedPackageData(packageName, version)
	if err != nil {
		store.SetCachedPackageData(packageName, version, store.CachedPackageData{SuccessfullDownload: false})
	}

	// If we didn't download the package yet, download it
	if !d.SuccessfullDownload {
		// Download the package
		data, err := downloadPackage(url)
		if err != nil {
			return false, err
		}

		store.SetCachedPackageData(packageName, version, store.CachedPackageData{SuccessfullDownload: true})

		if shasum != "" {
			if !integrity.CheckIntegrity(data, shasum) {
				return false, fmt.Errorf("integrity check failed")
			}
		}

		// Write data to the store/temp
		path, err := store.WriteTempFile(data, packageName, version)
		if err != nil {
			return false, err
		}

		// Decompress the package (tgz)
		err = tgz.LoadTgzAndExtractTo(path, store.GetDefaultStorePath()+packageName+"/"+version+"/data/")
		if err != nil {
			return false, err
		}
	
		store.SetCachedPackageData(packageName, version, store.CachedPackageData{SuccessfullDownload: true, SuccessfullExtract: true})

		// Delete the temp file
		err = store.DeleteTempFile(packageName, version)
		return false, err
	}

	return true, nil
}

func downloadPackage(url string) (data []byte, err error) {
	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	// Create a new client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	// Read the response body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return data, resp.Body.Close()
}
