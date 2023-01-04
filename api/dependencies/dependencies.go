package api

import (
	"github.com/davipatricio/dnpm/api"
)

func AddDependency(pkg api.PackageJSON, dependency string, version string) (api.PackageJSON, error) {
	// Check if the dependency is already in the package.json
	_, ok := pkg.Dependencies[dependency]
	if ok {
		return pkg, nil
	}

	// Add the dependency to the package.json
	pkg.Dependencies[dependency] = version

	return pkg, nil
}

func RemoveDependency(pkg api.PackageJSON, dependency string) (api.PackageJSON, error) {
	// Check if the dependency is in the package.json
	_, ok := pkg.Dependencies[dependency]
	if !ok {
		return pkg, nil
	}

	// Remove the dependency from the package.json
	delete(pkg.Dependencies, dependency)

	return pkg, nil
}
