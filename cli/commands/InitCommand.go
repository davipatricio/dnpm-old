package commands

import (
	"github.com/urfave/cli/v2"
)

func InitCommandRun(c *cli.Context) error {
	return nil
}

func GetInitCommandData() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Aliases: []string{"new"},
		Usage:   "Creates a new package.json in the current directory",
		Action:  InitCommandRun,
	}
}
