package main

import (
	"fmt"
	"os"

	"github.com/davipatricio/dnpm/cli/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "dnpm",
		Version:   "0.0.1",
		Usage:     "Main command to run dnpm utilities",
		ArgsUsage: "[args and options]",
		Authors: []*cli.Author{
			{
				Name:  "Davi Patricio",
				Email: "davipatricio2006@gmail.com",
			},
		},
		// add help
		Commands: []*cli.Command{
			commands.GetHelpCommandData(),
			commands.GetInitCommandData(),
		},
        CommandNotFound: func(cCtx *cli.Context, command string) {
			fmt.Printf("The command %s was not found.\nSee 'dnpm help' for more information.\n", command)
        },
		EnableBashCompletion: true,
		Suggest: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
