package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/davipatricio/dnpm/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:     "init [package name]",
	Short:   "Creates a new package.json in the current directory",
	Aliases: []string{"new"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if package.json already exists on current directory
		if _, err := os.Stat("package.json"); err == nil {
			cmd.Println("package.json already exists on current directory")
			return
		}

		packageName := args[0]
		cwd, _ := os.Getwd()

		if packageName != "" {
			packageName = filepath.Base(cwd)
		}

		packageJSON := api.PackageJSON{}
		packageJSON.Name = packageName
		packageJSON.Version = "0.0.1"
		packageJSON.Description = "An awesome project created with dnpm!"
		packageJSON.Main = "index.js"
		packageJSON.Scripts = map[string]string{
			"test": "echo \"Error: no test specified\" && exit 1",
		}
		packageJSON.Type = "module"

		// Beautify package.json and disable escaping
		bytes := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bytes)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.SetIndent("", "  ")
		jsonEncoder.Encode(packageJSON)

		json.Unmarshal(bytes.Bytes(), &packageJSON)

		os.WriteFile("package.json", bytes.Bytes(), 0644)

		return
	},
}
