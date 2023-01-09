package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/davipatricio/dnpm/api/integrity"
	"github.com/davipatricio/dnpm/api/store"
	"github.com/davipatricio/dnpm/api/tgz"
)

const (
	ErrorPackageOrVersionNotFound = "package or version not found"
	ErrorServerError              = "server error"
)

// DownloadPackage downloads a package from a url, saves it to the store and extracts it
// The first boolean value is true if the package was already downloaded and false if it was downloaded now
// To skip the integrity check, pass an empty string as the shasum
//
//	wasCached, err := DownloadAndSavePackage("https://registry.npmjs.org/helly/-/helly-1.1.0.tgz", "", "helly", "1.1.0")
func DownloadAndSavePackage(url string, shasum, packageName, version string) (bool, error) {
	d, err := store.GetCachedPackageData(packageName, version)
	if err != nil {
		store.SetCachedPackageData(packageName, version, store.CachedPackageData{SuccessfullDownload: false})
	}

	// if we had a successful download, but not a successful extraction, try to extract it again
	if d.SuccessfullDownload && !d.SuccessfullExtract {
		basePath := store.GetDefaultStorePath() + packageName + "/" + version + "/data"
		path := basePath + "/temp.tgz"

		// Delete the package folder incase of incomplete extraction
		os.RemoveAll(basePath + "/package/")

		// Extract the package
		err = tgz.LoadTgzAndExtractTo(path, basePath+"/")

		// If we have an error, lets redownload the package
		if err == nil {
			store.SetCachedPackageData(packageName, version, store.CachedPackageData{SuccessfullDownload: true, SuccessfullExtract: true})

			// Delete the temp file
			err = store.DeleteTempFile(packageName, version)
			return true, err
		}

		// Set the successfull download to false, so we can redownload the package
		d.SuccessfullDownload = false
	}

	// If we didn't download the package yet, download it
	if !d.SuccessfullDownload {
		// Download the package
		data, err := downloadPackage(url)
		if err != nil {
			return false, err
		}

		store.SetCachedPackageData(packageName, version, store.CachedPackageData{SuccessfullDownload: true})

		// If we have a shasum, check the integrity
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

	err = resp.Body.Close()

	// Check the response status code
	switch resp.StatusCode {
	case 200:
		return data, err
	case 404:
		return data, fmt.Errorf(ErrorPackageOrVersionNotFound)
	case 500:
		return data, fmt.Errorf(ErrorServerError)
	default:
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return data, err
		}
		return data, fmt.Errorf("unknown error")
	}
}
