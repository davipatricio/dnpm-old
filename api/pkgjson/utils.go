package pkgjson

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

	return pkg, file.Close()
}

// Tries to find the nearest package.json (up to 20 directories)
//
//	path, wasFound := FindNearestPackageJSON("./")
func FindNearestPackageJSON(dir string) (path string, found bool) {
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

