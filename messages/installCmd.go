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

func noPkgProvidedRaw() {
	fmt.Println(colors.Red("No package provided."))
	fmt.Println(colors.Cyan("Please provide a package to install."))
}

func noPkgProvidedEmojis() {
	fmt.Println("🤷 " + colors.Red("No package provided."))
	fmt.Println("ℹ️ " + colors.Cyan("Please provide a package to install."))
}

func installingPkgsRaw(pkgs []string) {
	str := "Installing packages:"
	for _, pkg := range pkgs {
		str += "\n  " + styles.Bold(colors.Green(pkg))
	}
	fmt.Println(str)
}

func installingPkgsEmojis(pkgs []string) {
	str := "➕ Installing packages:"
	for _, pkg := range pkgs {
		str += "\n📦  " + styles.Bold(colors.Green(pkg))
	}
	fmt.Println(str)
}

func noPkgJsonFoundRaw() {
	fmt.Println(colors.Red("No package.json was found."))
	fmt.Println(colors.Cyan("Please run 'dnpm init' to create a package.json file."))
}

func noPkgJsonFoundEmojis() {
	fmt.Println("🤷 " + colors.Red("No package.json was found."))
	fmt.Println("ℹ️ " + colors.Cyan("Please run 'dnpm init' to create a package.json file."))
}

func foundPkgInstallRaw() {
	fmt.Println(colors.Cyan("Installing packages present on 'package.json'"))
}

func foundPkgInstallEmojis() {
	fmt.Println("📦 " + colors.Cyan("Installing packages present on 'package.json'"))
}
