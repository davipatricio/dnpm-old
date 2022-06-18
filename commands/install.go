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

	"github.com/davipatricio/colors/colors"
)

func RunInstallCmd() bool {
	// Argument parsing

	// os.Args[1] will always be "add", "install" or "i" (see dnpm.go)
	installCmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	showEmojis := installCmd.Bool("emoji", false, "Whether to show emojis on the output.")
	showDebug := installCmd.Bool("debug", false, "Whether to show additional information on the output.")

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
		installSpecificPackages(packagesArgs, true, *showEmojis, *showDebug)
		return false
	}

	// Tell the user that we couldn't find a package.json and recommend the use of "dnpm init"
	messages.NoPkgJSONFoundInstallCmd(*showEmojis)

	return false
}

// TODO: install all packages from package.json
func installPackagesPresentOnPackageJSON(path string) {
	// If this function is called, means that we found a package.json and no packages/arguments were provided
}

func installSpecificPackages(packages []string, manual, showEmojis, showDebug bool) {
	var wg sync.WaitGroup
	startTime := time.Now().UnixMilli()

	wg.Add(1)
	// Spawn a goroutine to download the packages to download packages faster
	go func() {
		for _, pkg := range packages {
			// Get the package name from the provided string e.g. "typescript@nightly" -> "typescript"
			name := utils.GetPkgName(pkg)
			// Get the version from the provided string e.g. "typescript@nightly" -> "nightly"
			version := utils.GetPkgVersionOrTag(pkg)
			// Make a request to the Yarn registry requesting the package info
			d, _ := rest.GetPkg(name)
			if d["error"] != nil {
				messages.PkgNotFoundInstallCmd(showEmojis, name)
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

			// If we have the list of versions, we should check if the provided version is valid
			if d["versions"].(map[string]interface{})[version] != nil {
				// Verify if the package is already cached
				_, err := os.Stat(utils.GetStoreDir() + "/" + name + "/" + version)
				// If the folder doesn't exist, we should create it
				if err != nil {
					createFolderToPkg(name, version)
				}

				versionData := d["versions"].(map[string]interface{})[version].(map[string]interface{})
				downloadUrl := versionData["dist"].(map[string]interface{})["tarball"].(string)

				go func() {
					// If we could create the folder and if got an error means that the package is not cached
					if utils.PkgAlreadyCached(name, version) && err != nil {
						installDebug("Downloading package "+name+" ("+version+")", showDebug)
						// Download the tgz to the temp folder
						rest.DownloadPkgTgz(downloadUrl, utils.GetTempDir()+"/"+name+"-"+version+".tgz")
						installDebug("Extracting package "+name+" ("+version+")", showDebug)
						// Extract the tgz to the store folder
						utils.DecompressTgz(utils.GetTempDir()+"/"+name+"-"+version+".tgz", utils.GetStoreDir()+"/"+name+"/"+version)
						// Remove the temp tgz
						os.Remove(utils.GetTempDir() + "/" + name + "-" + version + ".tgz")
					} else {
						installDebug("Package "+name+" ("+version+") is already cached", showDebug)
					}
					wg.Done()
				}()

				// Check if there are dependencies
				deps, ok := versionData["dependencies"].(map[string]interface{})
				if ok {
					// Loop through each dependencies
					for depName, depVer := range deps {
						installDebug("Verifying if dependency "+depName+" ("+depVer.(string)+") is cached", showDebug)
						// If the dependency is not cached, we should download it
						if !utils.PkgAlreadyCached(depName, utils.RemovePkgVersionRange(depVer.(string))) {
							installDebug("Dependency "+depName+" is not cached.\nDownloading dependency "+depName+" ("+depVer.(string)+")", showDebug)
							wg.Add(1)
							go func() {
								// Call this function again to download the dependency
								// So we don't have duplicated code
								installSpecificPackages([]string{depName + "@" + depVer.(string)}, false, showEmojis, showDebug)
								wg.Done()
							}()
						}
					}
				}
			} else {
				// If for some reason, we cant get the version list (server down?), warn the user
				messages.VersionNotFoundInstallCmd(showEmojis)
				wg.Done()
			}
		}
	}()

	wg.Wait()
	endTime := time.Now().UnixMilli()

	// We should check this so we don't spam the output
	// Saying which packages were downloaded
	if manual {
		messages.DoneInstallCmd(showEmojis, endTime-startTime)
	}
}

func createFolderToPkg(pkg, version string) {
	// Get the folder we store cached packages
	dir := utils.GetStoreDir()
	// Verify if the package is already cached
	depCached := utils.PkgAlreadyCached(pkg, version)
	if !depCached {
		// If the package is not cached, we should create the folder
		err := os.MkdirAll(dir+"/"+pkg+"/"+version, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func installDebug(info string, showDebug bool) {
	if showDebug {
		fmt.Println(colors.Gray(info))
	}
}
