package commands

import (
	"dnpm/messages"
	"dnpm/utils"
	"flag"
	"fmt"
	"os"
)

func RunVersionCmd() {
	// Argument parsing

	// os.Args[1] will always be "v", "version" or "ver" (see dnpm.go)
	versionCmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	simplifyOutput := versionCmd.Bool("simplify", false, "Whether show only the version, without emojis, colors or words.")
	simplifyOutputChar := versionCmd.Bool("s", false, "Whether show only the version, without emojis, colors or words.")
	showEmojis := versionCmd.Bool("emoji", false, "Whether to show emojis on the output.")
	versionCmd.Parse(os.Args[2:])

	// Command code
	if *simplifyOutput || *simplifyOutputChar {
		fmt.Println(utils.VERSION)
	} else {
		messages.VersionCmd(*showEmojis)
	}
}
