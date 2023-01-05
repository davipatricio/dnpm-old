package commands

import "github.com/urfave/cli/v2"

func HelpCommandRun(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func GetHelpCommandData() *cli.Command {
	return &cli.Command{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Shows the application help",
		Action:  HelpCommandRun,
	}
}
