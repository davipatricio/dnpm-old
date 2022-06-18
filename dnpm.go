package main

import (
	"dnpm/cmds"
	"dnpm/msgs"
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
