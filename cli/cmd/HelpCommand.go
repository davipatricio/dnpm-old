package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.SetHelpCommand(helpCommand)
}

var helpCommand = &cobra.Command{
	Use:     "help [command name]",
	Short:   "Shows the list of available commands or help for a specific command",
	Aliases: []string{"h"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmdArg := args[0]
			for _, command := range rootCmd.Commands() {
				if command.Name() == cmdArg {
					command.Help()
					return
				}
			}
		} else {
			fmt.Println("Available commands:")
			for _, command := range rootCmd.Commands() {
				// <command name>, <aliases> tab <command description>
				fmt.Printf("   %s, %s\t%s\n", command.Name(), command.Aliases, command.Short)
			}
		}
	},
}
