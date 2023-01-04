package dependencies

import (
	"github.com/davipatricio/dnpm/api"
)

func AddDevDependency(pkg api.PackageJSON, dependency string, version string) (api.PackageJSON, error) {
	// Check if the dependency is already in the package.json
	_, ok := pkg.DevDependencies[dependency]
	if ok {
		return pkg, nil
	}

	// Add the dependency to the package.json
	pkg.DevDependencies[dependency] = version

	return pkg, nil
}

func RemoveDevDependency(pkg api.PackageJSON, dependency string) (api.PackageJSON, error) {
	// Check if the dependency is in the package.json
	_, ok := pkg.DevDependencies[dependency]
	if !ok {
		return pkg, nil
	}

	// Remove the dependency from the package.json
	delete(pkg.DevDependencies, dependency)

	return pkg, nil
}
