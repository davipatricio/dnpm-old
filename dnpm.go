package main

import (
	"dnpm/commands"
	"dnpm/messages"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		msgs.EmptyCmd()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "v", "ver", "version":
		cmds.RunVersionCmd()
	case "add", "i", "install":
		cmds.RunInstallCmd()
	case "init":
		cmds.RunInitCmd()
	default:
		msgs.EmptyCmd()
		os.Exit(1)
	}
}
