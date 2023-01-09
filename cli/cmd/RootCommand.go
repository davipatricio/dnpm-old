package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnpm",
	Short: "Main command to run dnpm utilities",
	Run: func(cmd *cobra.Command, args []string) {
		// Run the help command if no arguments are passed
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
