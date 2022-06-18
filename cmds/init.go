package cmds

import (
	"dnpm/msgs"
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
		msgs.InitErrReadingCmd(*showEmojis)
		return false
	}

	path, found, err := utils.GetNearestPackageJSON()
	if found && path == workingDir+"/package.json" {
		msgs.InitExistsCmd(*showEmojis)
		return false
	}

	if path != "" && !found && err != nil {
		msgs.InitCmd(*showEmojis)
		utils.CreateEmptyPackageJSON(path)
		msgs.InitDoneCmd(*showEmojis)
		return false
	}

	msgs.InitErrReadingCmd(*showEmojis)
	return false
}
