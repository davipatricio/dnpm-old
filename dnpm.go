package main

import (
	"dnpm/commands"
	"dnpm/messages"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		messages.EmptyCmd(true)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "v", "ver", "version":
		commands.RunVersionCmd()
	case "add", "i", "install":
		commands.RunInstallCmd()
	case "init":
		commands.RunInitCmd()
	case "help", "h", "?":
		messages.EmptyCmd(false)
	default:
		messages.EmptyCmd(true)
		os.Exit(1)
	}
}
