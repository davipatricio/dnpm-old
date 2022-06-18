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
		messages.InitErrReadingCmd(*showEmojis)
		return false
	}

	path, found, err := utils.GetNearestPackageJSON()
	if found && path == workingDir+"/package.json" {
		messages.InitExistsCmd(*showEmojis)
		return false
	}

	if path != "" && !found && err != nil {
		messages.InitCmd(*showEmojis)
		utils.CreateEmptyPackageJSON()
		messages.InitDoneCmd(*showEmojis)
		return false
	}

	messages.InitErrReadingCmd(*showEmojis)
	return false
}
