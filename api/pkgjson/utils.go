package api

import (
	"encoding/json"
	"io"
	"os"

	"github.com/davipatricio/dnpm/api"
)

// Reads a package.json, transforms it into a PackageJSON
//
//	pkg, err := ReadPackageJSON("./package.json")
//	fmt.Println(pkg.Name)
func ParseLocalPackageJSON(path string) (pkg api.PackageJSON, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()

	// Read the file and store its data
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}

	// Parse the JSON data using JSON marshal
	err = json.Unmarshal(data, &pkg)
	if err != nil {
		return
	}

	return
}

// Tries to find the nearest package.json (up to 20 directories)
//
//	path, wasFound := FindNearestPackageJSON("./")
func FindNearestPackageJSON(initialPath string) (path string, found bool) {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	tries := 0

	for {
		tries++
		if tries >= 20 {
			return
		}

		// Check if there is a package.json in the current directory
		_, err := os.Stat(dir + "/package.json")
		if err == nil {
			// If there is, return the path
			return dir + "/package.json", true
		}

		// If there is not, go up one directory
		dir = dir + "/.."
	}
}

// Saves a package.json to a file
func SavePackageJSON(pkg api.PackageJSON, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// Write the package.json to the file
	err = json.NewEncoder(file).Encode(pkg)
	if err != nil {
		return err
	}

	return nil
}

