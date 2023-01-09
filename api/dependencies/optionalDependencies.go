package dependencies

import (
	"github.com/davipatricio/dnpm/api"
)

func AddOptionalDependency(pkg api.PackageJSON, dependency string, version string) (api.PackageJSON, error) {
	// Check if the dependency is already in the package.json
	_, ok := pkg.OptionalDependencies[dependency]
	if ok {
		return pkg, nil
	}

	// Add the dependency to the package.json
	pkg.OptionalDependencies[dependency] = version

	return pkg, nil
}

func RemoveOptionalDependency(pkg api.PackageJSON, dependency string) (api.PackageJSON, error) {
	// Check if the dependency is in the package.json
	_, ok := pkg.OptionalDependencies[dependency]
	if !ok {
		return pkg, nil
	}

	// Remove the dependency from the package.json
	delete(pkg.OptionalDependencies, dependency)

	return pkg, nil
}
