package commands

import (
	"dnpm/messages"
	"dnpm/utils"
	"flag"
	"os"
)

func RunInstallCmd() bool {
	// Argument parsing

	// os.Args[1] will always be "add", "install" or "i" (see dnpm.go)
	installCmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	showEmojis := installCmd.Bool("emoji", false, "Whether to show emojis on the output.")

	// Command code
	path, found, _ := utils.GetNearestPackageJSON()

	// Check if the only argument/subcommand is "install"
	if len(os.Args) == 2 {
		installCmd.Parse(os.Args[2:])
		// If we found a package.json, we should run installPackagesPresentOnPackageJSON()
		// and install all dependencies and devDependecies.
		if found {
			messages.FoundPkgInstallCmd(*showEmojis)
			installPackagesPresentOnPackageJSON(path)
			return false
		} else {
			// Tell the user that we couldn't find a package.json and recommend the use of "dnpm init"
			messages.NoPkgJSONFoundInstallCmd(*showEmojis)
			return false
		}
	}

	// Check if we found a package.json and there are packages/arguments
	if found {
		installCmd.Parse(os.Args[2:])
		packagesArgs := installCmd.Args()
		if len(packagesArgs) < 1 {
			// If the user provide only an argument (e.g "dnpm install --emoji")
			// Tell the user no packages were provided to install
			messages.NoPkgProvidedInstallCmd(*showEmojis)
			return false
		}

		// Notify the user that we are installing the requested packages
		messages.InstallingPkgsInstallCmd(*showEmojis, packagesArgs)
		return false
	}

	// Tell the user that we couldn't find a package.json and recommend the use of "dnpm init"
	messages.NoPkgJSONFoundInstallCmd(*showEmojis)

	return false
}

func installPackagesPresentOnPackageJSON(path string) {
	// TODO: install all packages from package.json
}
