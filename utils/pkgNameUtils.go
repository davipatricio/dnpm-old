package utils

import (
	"regexp"
	"strings"
)

var RangeOperators = []string{">=", ">", "<=", "<", "~", "^"}

// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
var SemverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

func IsValidSemver(version string) bool {
	return SemverRegex.MatchString(version)
}

// Get the version from the provided string
// e.g. "typescript@nightly" -> "nightly"
// e.g. "ms@>=1.0.0" -> "1.0.0"
func GetPkgVersionOrTag(pkgName string) string {
	if !strings.Contains(pkgName, "@") {
		return "latest"
	}

	cleanPkgName := RemovePkgVersionRange(pkgName)

	// Get everything after the @ (can contain version range operators)
	version := cleanPkgName[strings.Index(cleanPkgName, "@")+1:]
	if IsValidSemver(version) {
		return version
	}

	// This will probably be a tag (e.g "dev")
	if !strings.Contains(version, ".") {
		return version
	}

	return ""
}

// Get the package name from the provided string
// e.g. "typescript@nightly" -> "typescript"
func GetPkgName(pkgName string) string {
	if !strings.Contains(pkgName, "@") {
		return pkgName
	}

	// Get everything before the @
	return pkgName[:strings.Index(pkgName, "@")]
}

// Remove the version range operator from the version
// e.g. ">=1.0.0" -> "1.0.0"
func RemovePkgVersionRange(fullVersion string) string {
	finalStr := fullVersion
	for _, operator := range RangeOperators {
		finalStr = strings.Replace(fullVersion, operator, "", -1)
		break
	}

	return finalStr
}

// Get the version range operator from the version.
// e.g. ">=1.0.0" -> ">="
func GetPkgVersionRange(fullVersion string) string {
	for _, operator := range RangeOperators {
		if strings.HasPrefix(fullVersion, operator) {
			return operator
		}
	}

	return ""
}
