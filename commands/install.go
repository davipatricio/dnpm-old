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

	path, found, err := utils.GetNearestPackageJSON()

	if len(os.Args) <= 2 {
		installCmd.Parse(os.Args[2:])
		if found {
			messages.FoundPkgInstallCmd(*showEmojis)
			installPackagesPresentOnPackageJSON(path)
			return false
		} else {
			messages.NoPkgJSONFoundInstallCmd(*showEmojis)
			return false
		}
	}

	if found {
		installCmd.Parse(os.Args[2:])

		packagesArgs := installCmd.Args()
		if len(packagesArgs) < 1 {
			messages.NoPkgProvidedInstallCmd(*showEmojis)
			return false
		}

		messages.InstallingPkgsInstallCmd(*showEmojis, packagesArgs)
		return false
	}

	if err != nil {
		panic(err)
	}

	return false
}

func installPackagesPresentOnPackageJSON(path string) {
	// TODO: install all packages from package.json
}
