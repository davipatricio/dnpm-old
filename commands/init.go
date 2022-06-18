package commands

import (
	"dnpm/messages"
	"dnpm/utils"
	"flag"
	"os"
)

func RunInitCmd() bool {
	// Argument parsing
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	showEmojis := initCmd.Bool("emoji", false, "Whether to show emojis on the output.")
	initCmd.Parse(os.Args[2:])

	// Command code
	workingDir, err := os.Getwd()
	if err != nil {
		// Tell the user that we couldn't get the working directory
		messages.InitErrReadingCmd(*showEmojis)
		return false
	}

	// If we find a package.json, we tell the user that a package.json already exists
	path, found, _ := utils.GetNearestPackageJSON()
	if found && path == workingDir+"/package.json" {
		messages.InitExistsCmd(*showEmojis)
		return false
	}

	// If we didn't find a package.json, we create one
	if !found {
		// Notify the user that we are creating the package.json
		messages.InitCmd(*showEmojis)
		utils.CreateEmptyPackageJSON()
		// Notify the user that the operation was successful
		messages.InitDoneCmd(*showEmojis)
		return false
	}

	return false
}
