package utils

import (
	"regexp"
	"strings"
)

var RangeOperators = []string{">=", ">", "<=", "<", "~", "^"}

// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
var SemverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// Verify if the provided string is a valid semver
func IsValidSemver(version string) bool {
	return SemverRegex.MatchString(version)
}

// Get the version from the provided string
// e.g. "typescript@nightly" -> "nightly"
// e.g. "ms@>=1.0.0" -> "1.0.0"
func GetPkgVersionOrTag(pkgName string) string {
	cleanPkgName := strings.TrimPrefix(pkgName, "@")
	if !strings.Contains(cleanPkgName, "@") {
		return "latest"
	}

	// Get everything after the @ (can contain version range operators)
	version := cleanPkgName
	version = version[strings.Index(version, "@")+1:]

	cleanVersion := RemovePkgVersionRange(version)

	if cleanVersion == "" || cleanVersion == "*" {
		return "latest"
	}

	if IsValidSemver(cleanVersion) {
		return cleanVersion
	}

	// This will probably is a tag (e.g "dev")
	if !strings.Contains(cleanVersion, ".") {
		return cleanVersion
	}

	return ""
}

// Get the package name from the provided string
// e.g. "typescript@nightly" -> "typescript"
// e.g @pkg/name@version -> "@pkg/name"
func GetPkgName(pkgName string) string {
	if strings.HasPrefix(pkgName, "@") {
		slice := strings.Split(pkgName, "@")
		if len(slice) >= 3 {
			if slice[2] != "" {
				return "@" + slice[1]
			}
		}
		return pkgName
	}

	if strings.Contains(pkgName, "@") {
		return pkgName[:strings.Index(pkgName, "@")]
	}

	return pkgName

}

// Remove the version range operator from the version
// e.g. ">=1.0.0" -> "1.0.0"
func RemovePkgVersionRange(fullVersion string) string {
	finalStr := fullVersion
	for _, operator := range RangeOperators {
		finalStr = strings.ReplaceAll(finalStr, operator, "")
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

// "@types\node\18.0.0\node" -> "@types\node\18.0.0"
func RemoveLastSubstring(original string, substring string) string {
	// Check if the substring appears two times in different places
	if strings.Count(original, substring) >= 2 {
		return strings.TrimSuffix(original, substring)
	}
	return original
}

// Get the package name from the provided string e.g. "@myorg/mypkg@nightly" -> "mypkg@nightly"
func RemoveOrgName(original string) string {
	if strings.HasPrefix(original, "@") {
		return original[strings.Index(original, "/")+1:]
	}
	return original
}

// Get the package name from the provided string e.g. "@myorg/mypkg@nightly" -> "@mypkg"
func GetOrgName(original string) string {
	if strings.HasPrefix(original, "@") {
		return original[:strings.Index(original, "/")]
	}
	return ""
}
