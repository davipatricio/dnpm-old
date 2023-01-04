package semver

import (
	"strconv"
	"strings"
)

// Checks if a version is compatible with a semver range
//
//	satisfies := VersionSatisfiesRange("1.2.3", ">=1.2.0")
//	satisfies := VersionSatisfiesRange("1.2.3", "<=1.2.0")
//	satisfies := VersionSatisfiesRange("1.2.3", "1.2.0")
//	satisfies := VersionSatisfiesRange("1.2.3", "^1.2.0")
//	satisfies := VersionSatisfiesRange("1.2.3", "x.x.x")
//	satisfies := VersionSatisfiesRange("1.2.3", "*")
//	satisfies := VersionSatisfiesRange("1.2.0", "1.0.0 || 1.2.0")
func VersionSatisfiesRange(version string, semver string) bool {
	if semver == "*" {
		return true
	}

	// >=
	if strings.HasPrefix(semver, ">=") {
		return versionGTE(version, strings.TrimPrefix(semver, ">="))
	}

	// <=
	if strings.HasPrefix(semver, "<=") {
		return versionLTE(version, strings.TrimPrefix(semver, "<="))
	}

	// >
	if strings.HasPrefix(semver, ">") {
		return versionGT(version, strings.TrimPrefix(semver, ">"))
	}

	// <
	if strings.HasPrefix(semver, "<") {
		return versionLT(version, strings.TrimPrefix(semver, "<"))
	}

	// =
	if strings.HasPrefix(semver, "=") {
		return versionEQ(version, strings.TrimPrefix(semver, "="))
	}

	// ~
	if strings.HasPrefix(semver, "~") {
		return versionTilde(version, strings.TrimPrefix(semver, "~"))
	}

	// ^
	if strings.HasPrefix(semver, "^") {
		return versionCaret(version, strings.TrimPrefix(semver, "^"))
	}

	// ||
	if strings.Contains(semver, "||") {
		return versionOr(version, semver)
	}

	// x.x.x
	if strings.Contains(semver, "x") {
		return versionX(version, semver)
	}

	// 1.2.3
	return versionEQ(version, semver)
}

// Checks if a version is greater than or equal to a semver range
//
//	satisfies := versionGTE("1.2.3", "1.2.0")
func versionGTE(version string, semver string) bool {
	return versionGT(version, semver) || versionEQ(version, semver)
}

// Checks if a version is less than or equal to a semver range
//
//	satisfies := versionLTE("1.2.3", "1.2.0")
func versionLTE(version string, semver string) bool {
	return versionLT(version, semver) || versionEQ(version, semver)
}

// Checks if a version is greater than a semver range
//
//	satisfies := versionGT("1.2.3", "1.2.0")
//	satisfies := versionGT("2.0.0", "1.2.0")
func versionGT(version string, semver string) bool {
	vs := strings.Split(version, ".")
	ss := strings.Split(semver, ".")

	for i := 0; i < len(ss); i++ {
		v, _ := parseInt(vs[i])
		s, _ := parseInt(ss[i])

		if v > s {
			return true
		}

		if v < s {
			return false
		}

	}

	return false
}

// Checks if a version is less than a semver range
//
//	satisfies := versionLT("1.2.3", "1.2.0")
func versionLT(version string, semver string) bool {
	vs := strings.Split(version, ".")
	ss := strings.Split(semver, ".")

	for i := 0; i < len(ss); i++ {
		v, _ := parseInt(vs[i])
		s, _ := parseInt(ss[i])

		if v < s {
			return true
		}

		if v > s {
			return false
		}
	}

	return false
}

// Checks if a version is equal to a semver range
//
//	satisfies := versionEQ("1.2.3", "1.2.0")
func versionEQ(version string, semver string) bool {
	vs := strings.Split(version, ".")
	ss := strings.Split(semver, ".")

	for i := 0; i < len(ss); i++ {
		v, _ := parseInt(vs[i])
		s, _ := parseInt(ss[i])

		if v != s {
			return false
		}
	}

	return true
}

// Checks if a version is equal to a semver range
//
//	satisfies := versionTilde("1.2.0", "1.2.0")
//	satisfies := versionTilde("1.2.1", "1.2.0")
func versionTilde(version string, semver string) bool {
	vs := strings.Split(version, ".")
	ss := strings.Split(semver, ".")

	for i := 0; i < len(ss); i++ {
		v, _ := parseInt(vs[i])
		s, _ := parseInt(ss[i])

		if v != s {
			return false
		}

		if i == 1 {
			return true
		}

	}

	return true
}

// Checks if a version is equal to a semver range
//
//	satisfies := versionCaret("1.2.0", "1.2.0")
//	satisfies := versionCaret("1.2.3", "1.2.0")
func versionCaret(version string, semver string) bool {
	vs := strings.Split(version, ".")
	ss := strings.Split(semver, ".")

	for i := 0; i < len(ss); i++ {
		v, _ := parseInt(vs[i])
		s, _ := parseInt(ss[i])

		if v != s {
			return false
		}

		if i != 0 {
			return true
		}

	}

	return true
}

// Checks if a version is equal to a semver range
//
//	satisfies := versionOr("1.2.3", "1.1.0 || 1.2.3")
func versionOr(version string, semver string) bool {
	ors := strings.Split(semver, "||")

	for _, or := range ors {
		if VersionSatisfiesRange(version, strings.TrimSpace(or)) {
			return true
		}
	}

	return false
}

// Checks if a version is equal to a semver range
//
//	satisfies := versionX("1.2.3", "1.2.0")
func versionX(version string, semver string) bool {
	vs := strings.Split(version, ".")
	ss := strings.Split(semver, ".")

	for i := 0; i < len(ss); i++ {
		v, _ := parseInt(vs[i])
		s, _ := parseInt(ss[i])

		if s == -1 {
			return true
		}

		if v != s {
			return false
		}
	}

	return true
}

// Checks if a version is equal to a semver range
//
//	satisfies := versionX("1.2.3", "1.2.0")
func versionXRange(version string, semver string) bool {
	vs := strings.Split(version, ".")
	ss := strings.Split(semver, ".")

	for i := 0; i < len(ss); i++ {
		v, _ := parseInt(vs[i])
		s, _ := parseInt(ss[i])

		if s == -1 {
			return true
		}

		if v != s {
			return false
		}
	}

	return true
}

// parseInt
func parseInt(s string) (int, error) {
	if s == "x" {
		return -1, nil
	}

	return strconv.Atoi(s)
}
