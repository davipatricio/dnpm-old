package rest

import (
	"io"
	"net/http"
	"os"
)

// Download the file from the given url and save it to the given path.
func DownloadPkgTgz(url, pathToSave string) (error) {
	// Create the file
	out, err := os.Create(pathToSave)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	return nil
}
