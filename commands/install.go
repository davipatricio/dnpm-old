package commands

import (
	"dnpm/messages"
	"dnpm/rest"
	"dnpm/utils"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
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
		installSpecificPackages(packagesArgs, true)
		return false
	}

	// Tell the user that we couldn't find a package.json and recommend the use of "dnpm init"
	messages.NoPkgJSONFoundInstallCmd(*showEmojis)

	return false
}

func installPackagesPresentOnPackageJSON(path string) {
	// TODO: install all packages from package.json
}

func installSpecificPackages(packages []string, manual bool) {
	var wg sync.WaitGroup
	startTime := time.Now().UnixMilli()

	wg.Add(1)
	go func() {
		for _, pkg := range packages {

			name := utils.GetPkgName(pkg)
			version := utils.GetPkgVersionOrTag(pkg)
			d, _ := rest.GetPkg(name)
			if d["error"] != nil {
				fmt.Println("Pacote desconhecido.")
				wg.Done()
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
				_, err := os.Stat(utils.GetStoreDir() + "/" + name + "/" + version)
				createFolderToPkg(name, version)

				versionData := d["versions"].(map[string]interface{})[version].(map[string]interface{})
				downloadUrl := versionData["dist"].(map[string]interface{})["tarball"].(string)

				go func() {
					if utils.PkgAlreadyCached(name, version) && err != nil {
						fmt.Println("Baixando pacote", name, "@", version, "...")
						rest.DownloadPkgTgz(downloadUrl, utils.GetTempDir()+"/"+name+"-"+version+".tgz")
						fmt.Println("Descompactando pacote", name, "@", version, "...")
						utils.DecompressTgz(utils.GetTempDir()+"/"+name+"-"+version+".tgz", utils.GetStoreDir()+"/"+name+"/"+version)
						os.Remove(utils.GetTempDir() + "/" + name + "-" + version + ".tgz")
					} else {
						fmt.Println("Pacote", name, "@", version, "já está no cache e não precisou ser baixada novamente.")
					}
					wg.Done()
				}()

				deps, ok := versionData["dependencies"].(map[string]interface{})
				if ok {
					for depName, depVer := range deps {
						fmt.Println("Verificando se a dependência", depName, "@", depVer.(string), "já foi baixada alguma vez...")
						if !utils.PkgAlreadyCached(depName, utils.RemovePkgVersionRange(depVer.(string))) {
							fmt.Println("Dependencia", depName, "nunca foi baixada anteriormente. Efetuando download (Yarn Registry)...")
							wg.Add(1)
							go func() {
								installSpecificPackages([]string{depName + "@" + depVer.(string)}, false)
								wg.Done()
							}()
						}
					}
				}
			} else {
				fmt.Println("Versão desconhecida.")
				wg.Done()
			}
		}
	}()

	wg.Wait()
	endTime := time.Now().UnixMilli()
	if manual {
		fmt.Println("\n\n - Tempo de instalação de todas as dependencias etc (sem cache):", endTime-startTime, "ms")
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
