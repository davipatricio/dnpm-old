package messages

import (
	"fmt"

	"github.com/gookit/color"
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
	color.Red.Printf("Version not found.")
	color.Cyan.Printf("Please provide a valid version.")
}

func versionNotFoundEmojis() {
	color.Red.Printf("🤷 Version not found.")
	color.Cyan.Printf("ℹ️  Please provide a valid version.")
}

// When an operation is completed
func doneOperationRaw(ms int64) {
	color.Green.Printf("Done in %vms.\n", ms)
}

func doneOperationEmojis(ms int64) {
	color.Green.Printf("✅ Done in %vms.\n", ms)
}

// Package not found on the registry
func pkgNotFoundRaw(pkg string) {
	color.Red.Println("Package '" + pkg + "' was not found on the registry.")
}

func pkgNotFoundEmojis(pkg string) {
	color.Red.Println("🤷 Package '" + pkg + "' was not found on the registry.")
}

// No package provided
func noPkgProvidedRaw() {
	color.Red.Println("No package provided.")
	color.Cyan.Println("Please provide a package to install.")
}

func noPkgProvidedEmojis() {
	color.Red.Println("🤷 No package provided.")
	color.Cyan.Println("ℹ️ Please provide a package to install.")
}

// Installing requested packages
func installingPkgsRaw(pkgs []string) {
	str := "Installing packages:"
	for _, pkg := range pkgs {
		str += "\n  " + color.OpBold.Render(color.Green.Render(pkg)) + "\n\n"
	}
	fmt.Println(str)
}

func installingPkgsEmojis(pkgs []string) {
	str := "➕ Installing packages:"
	for _, pkg := range pkgs {
		str += "\n📦  " + color.OpBold.Render(color.Green.Render(pkg)) + "\n\n"
	}
	fmt.Println(str)
}

// No package.json found
func noPkgJsonFoundRaw() {
	color.Red.Println("No package.json was found.")
	color.Cyan.Println("Please run 'dnpm init' to create a package.json file.")
}

func noPkgJsonFoundEmojis() {
	color.Red.Println("🤷 No package.json was found.")
	color.Cyan.Println("ℹ️  Please run 'dnpm init' to create a package.json file.")
}

// Installing package.json dependencies
func foundPkgInstallRaw() {
	color.Cyan.Println("Installing packages present on 'package.json'")
}

func foundPkgInstallEmojis() {
	color.Cyan.Println("📦 Installing packages present on 'package.json'")
}
