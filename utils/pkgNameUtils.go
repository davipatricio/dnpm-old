package utils

import "strings"

func GetPkgVersion(pkgName string) string {
	if !strings.Contains(pkgName, "@") {
		return "latest"
	}

	return pkgName[strings.Index(pkgName, "@")+1:]
}

func GetPkgName(pkgName string) string {
	if !strings.Contains(pkgName, "@") {
		return pkgName
	}

	return pkgName[:strings.Index(pkgName, "@")]
}
