package api

import (
	"encoding/json"
	"os"
)

func AddDependency(pkg PackageJSON, dependency string, version string) (PackageJSON, error) {
	// Check if the dependency is already in the package.json
	_, ok := pkg.Dependencies[dependency]
	if ok {
		return pkg, nil
	}

	// Add the dependency to the package.json
	pkg.Dependencies[dependency] = version

	return pkg, nil
}

func RemoveDependency(pkg PackageJSON, dependency string) (PackageJSON, error) {
	// Check if the dependency is in the package.json
	_, ok := pkg.Dependencies[dependency]
	if !ok {
		return pkg, nil
	}

	// Remove the dependency from the package.json
	delete(pkg.Dependencies, dependency)

	return pkg, nil
}

func AddDevDependency(pkg PackageJSON, dependency string, version string) (PackageJSON, error) {
	// Check if the dependency is already in the package.json
	_, ok := pkg.DevDependencies[dependency]
	if ok {
		return pkg, nil
	}

	// Add the dependency to the package.json
	pkg.DevDependencies[dependency] = version

	return pkg, nil
}

func RemoveDevDependency(pkg PackageJSON, dependency string) (PackageJSON, error) {
	// Check if the dependency is in the package.json
	_, ok := pkg.DevDependencies[dependency]
	if !ok {
		return pkg, nil
	}

	// Remove the dependency from the package.json
	delete(pkg.DevDependencies, dependency)

	return pkg, nil
}

func AddOptionalDependency(pkg PackageJSON, dependency string, version string) (PackageJSON, error) {
	// Check if the dependency is already in the package.json
	_, ok := pkg.OptionalDependencies[dependency]
	if ok {
		return pkg, nil
	}

	// Add the dependency to the package.json
	pkg.OptionalDependencies[dependency] = version

	return pkg, nil
}

func RemoveOptionalDependency(pkg PackageJSON, dependency string) (PackageJSON, error) {
	// Check if the dependency is in the package.json
	_, ok := pkg.OptionalDependencies[dependency]
	if !ok {
		return pkg, nil
	}

	// Remove the dependency from the package.json
	delete(pkg.OptionalDependencies, dependency)

	return pkg, nil
}

func SavePackageJSON(pkg PackageJSON, path string) error {
	// Open the file
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
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
