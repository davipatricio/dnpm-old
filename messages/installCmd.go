package messages

import (
	"fmt"

	"github.com/davipatricio/colors/colors"
	"github.com/davipatricio/colors/styles"
)

func FoundPkgInstallCmd(showEmojis bool) {
	if showEmojis {
		foundPkgInstallEmojis()
	} else {
		foundPkgInstallRaw()
	}
}

func NoPkgJSONFoundInstallCmd(showEmojis bool) {
	if showEmojis {
		noPkgJsonFoundEmojis()
	} else {
		noPkgJsonFoundRaw()
	}
}

func NoPkgProvidedInstallCmd(showEmojis bool) {
	if showEmojis {
		noPkgProvidedEmojis()
	} else {
		noPkgProvidedRaw()
	}
}

func InstallingPkgsInstallCmd(showEmojis bool, pkgs []string) {
	if showEmojis {
		installingPkgsEmojis(pkgs)
	} else {
		installingPkgsRaw(pkgs)
	}
}

func PkgNotFoundInstallCmd(showEmojis bool, pkg string) {
	if showEmojis {
		pkgNotFoundEmojis(pkg)
	} else {
		pkgNotFoundRaw(pkg)
	}
}

func DoneInstallCmd(showEmojis bool, ms int64) {
	if showEmojis {
		doneOperationEmojis(ms)
	} else {
		doneOperationRaw(ms)
	}
}

func VersionNotFoundInstallCmd(showEmojis bool) {
	if showEmojis {
		versionNotFoundEmojis()
	} else {
		versionNotFoundRaw()
	}
}

/* */

// Version not found
func versionNotFoundRaw() {
	fmt.Println(colors.Red("Version not found."))
	fmt.Println(colors.Cyan("Please provide a valid version."))
}

func versionNotFoundEmojis() {
	fmt.Println("🤷 " + colors.Red("Version not found."))
	fmt.Println("ℹ️ " + colors.Cyan("Please provide a valid version."))
}

// When an operation is completed
func doneOperationRaw(ms int64) {
	fmt.Printf(colors.Green("Done in %vms.\n"), ms)
}

func doneOperationEmojis(ms int64) {
	fmt.Printf("✅ "+colors.Green("Done in %vms.\n"), ms)
}

// Package not found on the registry
func pkgNotFoundRaw(pkg string) {
	fmt.Println(colors.Red("Package '" + pkg + "' was not found on the registry."))
}

func pkgNotFoundEmojis(pkg string) {
	fmt.Println("🤷 " + colors.Red("Package '"+pkg+"' was not found on the registry."))
}

// No package provided
func noPkgProvidedRaw() {
	fmt.Println(colors.Red("No package provided."))
	fmt.Println(colors.Cyan("Please provide a package to install."))
}

func noPkgProvidedEmojis() {
	fmt.Println("🤷 " + colors.Red("No package provided."))
	fmt.Println("ℹ️ " + colors.Cyan("Please provide a package to install."))
}

// Installing requested packages
func installingPkgsRaw(pkgs []string) {
	str := "Installing packages:"
	for _, pkg := range pkgs {
		str += "\n  " + styles.Bold(colors.Green(pkg)) + "\n\n"
	}
	fmt.Println(str)
}

func installingPkgsEmojis(pkgs []string) {
	str := "➕ Installing packages:"
	for _, pkg := range pkgs {
		str += "\n📦  " + styles.Bold(colors.Green(pkg)) + "\n\n"
	}
	fmt.Println(str)
}

// No package.json found
func noPkgJsonFoundRaw() {
	fmt.Println(colors.Red("No package.json was found."))
	fmt.Println(colors.Cyan("Please run 'dnpm init' to create a package.json file."))
}

func noPkgJsonFoundEmojis() {
	fmt.Println("🤷 " + colors.Red("No package.json was found."))
	fmt.Println("ℹ️ " + colors.Cyan("Please run 'dnpm init' to create a package.json file."))
}

// Installing package.json dependencies
func foundPkgInstallRaw() {
	fmt.Println(colors.Cyan("Installing packages present on 'package.json'"))
}

func foundPkgInstallEmojis() {
	fmt.Println("📦 " + colors.Cyan("Installing packages present on 'package.json'"))
}
