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

	initCommand.Flags().StringP("version", "v", "0.0.1", "Package version, should be a valid semver version")
	initCommand.Flags().StringP("license", "l", "", "The license of the package")
	initCommand.Flags().BoolP("overwrite", "", false, "Overwrite existing package.json")
}

var initCommand = &cobra.Command{
	Use:     "init [package name] [options]",
	Short:   "Creates a new package.json in the current directory",
	Example: "dnpm init my-awesome-package",
	Aliases: []string{"new"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Flag parsing
		version, _ := cmd.Flags().GetString("version")
		overwrite, _ := cmd.Flags().GetBool("overwrite")
		license, _ := cmd.Flags().GetString("license")

		// If the overwrite flag is not set, check if package.json already exists
		if !overwrite {
			// Check if package.json already exists on current directory
			if _, err := os.Stat("package.json"); err == nil {
				cmd.Println("package.json already exists on the current directory")
				return
			}
		}

		packageName := args[0]
		cwd, _ := os.Getwd()

		if packageName == "" {
			packageName = filepath.Base(cwd)
		}

		packageJSON := api.PackageJSON{}
		packageJSON.Name = packageName
		packageJSON.Version = version
		packageJSON.Description = "An awesome project created with dnpm!"
		packageJSON.Main = "index.js"
		packageJSON.Scripts = map[string]string{
			"test": "echo \"Error: no test specified\" && exit 1",
		}
		packageJSON.Type = "commonjs"
		packageJSON.License = license

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
