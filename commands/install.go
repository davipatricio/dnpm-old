package commands

import (
	"dnpm/messages"
	"dnpm/rest"
	"dnpm/utils"
	"flag"
	"fmt"
	"os"
)

func RunInstallCmd() bool {
	// Argument parsing

	// os.Args[1] will always be "add", "install" or "i" (see dnpm.go)
	installCmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	showEmojis := installCmd.Bool("emoji", false, "Whether to show emojis on the output.")

	// Command code
	path, found, _ := utils.GetNearestPackageJSON()

	// Check if the only argument/subcommand is "install"
	if len(os.Args) == 2 {
		installCmd.Parse(os.Args[2:])
		// If we found a package.json, we should run installPackagesPresentOnPackageJSON()
		// and install all dependencies and devDependecies.
		if found {
			messages.FoundPkgInstallCmd(*showEmojis)
			installPackagesPresentOnPackageJSON(path)
			return false
		} else {
			// Tell the user that we couldn't find a package.json and recommend the use of "dnpm init"
			messages.NoPkgJSONFoundInstallCmd(*showEmojis)
			return false
		}
	}

	// Check if we found a package.json and there are packages/arguments
	if found {
		installCmd.Parse(os.Args[2:])
		packagesArgs := installCmd.Args()
		if len(packagesArgs) < 1 {
			// If the user provide only an argument (e.g "dnpm install --emoji")
			// Tell the user no packages were provided to install
			messages.NoPkgProvidedInstallCmd(*showEmojis)
			return false
		}

		// Notify the user that we are installing the requested packages
		messages.InstallingPkgsInstallCmd(*showEmojis, packagesArgs)
		installSpecificPackages(packagesArgs)
		return false
	}

	// Tell the user that we couldn't find a package.json and recommend the use of "dnpm init"
	messages.NoPkgJSONFoundInstallCmd(*showEmojis)

	return false
}

func installPackagesPresentOnPackageJSON(path string) {
	// TODO: install all packages from package.json
}

func installSpecificPackages(packages []string) {
	for _, pkg := range packages {
		name := utils.GetPkgName(pkg)
		version := utils.GetPkgVersionOrTag(pkg)
		d, _ := rest.GetPkg(name)
		if d["error"] != nil {
			fmt.Println("Pacote desconhecido.")
			continue
		}

		// If there the package has a tag that is the same as the provided version,
		// we should install that tag instead of the version.
		if d["dist-tags"] != nil {
			distTags := d["dist-tags"].(map[string]interface{})
			// If no version was provided, use the latest version
			if version == "" {
				// Get the property latest from d.dist-tags
				version = distTags["latest"].(string)
			} else if distTags[version] != nil {
				// Get the version of the tag
				version = distTags[version].(string)
			}
		}

		if d["versions"].(map[string]interface{})[version] != nil {
			createFolderToPkg(name, version)

			versionData := d["versions"].(map[string]interface{})[version].(map[string]interface{})
			downloadUrl := versionData["dist"].(map[string]interface{})["tarball"].(string)

			deps, ok := versionData["dependencies"].(map[string]interface{})
			if ok {
				for depName, depVer := range deps {
					fmt.Println("Verificando se a dependência", depName, "@", depVer.(string), "já foi baixada alguma vez...")

					if !utils.PkgAlreadyCached(depName, depVer.(string)) {
						fmt.Println("Dependencia", depName, "nunca foi baixada anteriormente. Efetuando download (Yarn Registry)...")
						installSpecificPackages([]string{depName + "@" + depVer.(string)})
					}
				}
			}

			fmt.Println("URL: ", downloadUrl)
		} else {
			fmt.Println("Versão desconhecida.")
		}
	}
}

func createFolderToPkg(pkg, version string) {
	dir := utils.GetStoreDir()
	depCached := utils.PkgAlreadyCached(pkg, version)
	if !depCached {
		err := os.MkdirAll(dir+"/"+pkg+"/"+version, 0755)
		if err != nil {
			panic(err)
		}
	}
}
