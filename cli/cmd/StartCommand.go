package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/davipatricio/dnpm/api/pkgjson"
	"github.com/spf13/cobra"
)

// This runs a predefined command specified in the "start" property of a package's "scripts" object
// If the "scripts" object does not define a "start" property, dnpm will run 'node index.js'
func init() {
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use:     "run [script name]",
	Short:   "Runs a predefined script specified in the package.json",
	Example: "dnpm run build",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		if arg == "" {

			return
		}

		workingDir, _ := os.Getwd()

		path, found := pkgjson.FindNearestPackageJSON("./")
		// Check if the current directory has a package.json
		if !found {
			cmd.Println("Could not find a package.json in a nearby directory")
			return
		}

		packageJSON, err := pkgjson.ParseLocalPackageJSON(path)
		if err != nil {
			cmd.Println("Could not parse package.json. Is it a valid JSON file?")
			return
		}

		// Check if the package.json has a "scripts" object
		if packageJSON.Scripts == nil {
			cmd.Println("package.json does not have a \"scripts\" object")
			return
		}

		// Check if the "scripts" object has an "arg" property
		if packageJSON.Scripts[arg] == "" {
			cmd.Println("package.json does not have a \"" + arg + "\" property in the \"scripts\" object")
			return
		}

		// Run the command specified in the "arg" property
		fmt.Printf(" > %s@%s %s %s\n", packageJSON.Name, packageJSON.Version, arg, workingDir)
		fmt.Printf(" > %s\n\n", packageJSON.Scripts[arg])

		var scriptCommand *exec.Cmd = nil
		switch runtime.GOOS {
		case "windows":
			scriptCommand = exec.Command("cmd", "/c", packageJSON.Scripts[arg])
		default:
			scriptCommand = exec.Command("sh", "-c", packageJSON.Scripts[arg])
		}

		// todo: send output to the terminal
		scriptCommand.Output()
		if err != nil {
			log.Fatal(err)
		}

		return
	},
}

func unescapeString(s string) string {
	regex1 := regexp.MustCompile(`/(^|[^\\])(\\\\)*\\$/`)
	regex2 := regexp.MustCompile(`/(^|[^\\])((\\\\)*")/g/`)

	s = regex1.ReplaceAllString(s, "$&\\")
	return regex2.ReplaceAllString(s, "$1\\$2")

}
