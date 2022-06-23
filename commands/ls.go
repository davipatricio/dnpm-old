package commands

import (
	"dnpm/messages"
	"dnpm/utils"
	"flag"
	"os"
	"strings"

	"github.com/davipatricio/colors/colors"
	"golang.org/x/exp/maps"
)

func RunLsCmd() {
	lsCmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	showEmojis := lsCmd.Bool("emoji", false, "Whether to show emojis on the output.")
	showAll := lsCmd.Bool("all", false, "Whether to show all on the output.")

	lsCmd.Parse(os.Args[2:])
	messages.LsCmd(getPackages(*showEmojis, *showAll))
}

func getPackages(showEmojis, showAll bool) string {
	if showEmojis {
		return getPackagesEmojis(showAll)
	} else {
		return getPackagesRaw(showAll)
	}
}

func getPackagesEmojis(showAll bool) string {
	if isEmpty() {
		return "📂 Empty"
	}

	var title string

	pacJSON := utils.GetPackageJSON()

	if pacJSON.Name != "" {
		title = pacJSON.Name + colors.Magenta(" ➜  ") + utils.GetExecDir()
	} else {
		title = utils.GetExecDir()
	}

	tree := utils.MakeTree(title)
	dependencies := map[string]string{}

	utils.Assign(pacJSON.Dependencies, dependencies)
	utils.Assign(pacJSON.DevDependencies, dependencies)
	utils.Assign(pacJSON.OptionalDependencies, dependencies)
	utils.Assign(pacJSON.PeerDependencies, dependencies)

	for pac, version := range dependencies {
		version = utils.RemovePkgVersionRange(version)

		if !verifyPackageExists(pac) {
			tree.Add("📦 " + colors.Red(pac+"@"+version))
		} else {
			branch := tree.Add("📦 " + colors.Green(pac+"@"+version))

			if showAll {
				getPackagesOfPackageAndAddToBranch(true, pac, branch, []string{})
			}
		}
	}

	return tree.Text()
}

func getPackagesRaw(showAll bool) string {
	if isEmpty() {
		return "Empty"
	}

	var title string
	pacJSON := utils.GetPackageJSON()

	if pacJSON.Name != "" {
		title = pacJSON.Name + colors.Magenta("@ ") + utils.GetExecDir()
	} else {
		title = utils.GetExecDir()
	}

	tree := utils.MakeTree(title)
	dependencies := map[string]string{}

	utils.Assign(pacJSON.Dependencies, dependencies)
	utils.Assign(pacJSON.DevDependencies, dependencies)
	utils.Assign(pacJSON.OptionalDependencies, dependencies)
	utils.Assign(pacJSON.PeerDependencies, dependencies)

	for pac, version := range dependencies {
		version = utils.RemovePkgVersionRange(version)

		if !verifyPackageExists(pac) {
			if isOptionalDependency(pac, "default") {
				tree.Add(colors.Yellow(pac + "@" + version))
			} else {
				tree.Add(colors.Red(pac + "@" + version))
			}
		} else {
			branch := tree.Add(colors.Green(pac + "@" + version))

			if showAll {
				getPackagesOfPackageAndAddToBranch(false, pac, branch, []string{})
			}
		}
	}

	return tree.Text()
}

func getPackagesOfPackageAndAddToBranch(showEmojis bool, pac string, branch utils.Branch, already []string) {
	pkgJSON, found := utils.GetPackageJSONOfPackage(pac)
	dependenciesPack := map[string]string{}

	if found {
		utils.Assign(pkgJSON.Dependencies, dependenciesPack)
		utils.Assign(pkgJSON.DevDependencies, dependenciesPack)
		utils.Assign(pkgJSON.PeerDependencies, dependenciesPack)
		utils.Assign(pkgJSON.OptionalDependencies, dependenciesPack)

		for pkgName, version := range dependenciesPack {
			if utils.Contains(already, pkgName) {
				continue
			}

			version = utils.RemovePkgVersionRange(version)
			if showEmojis {
				branch.Add("📦 " + colors.Green(pkgName + "@" + version))
			} else {
				branch.Add(colors.Green(pkgName + "@" + version))
			}
		}
	}
}

func isEmpty() bool {
	return len(utils.GetPackageJSON().Dependencies) == 0
}

func verifyPackageExists(packageName string) bool {
	_, err := os.Stat("node_modules/" + packageName)

	return !os.IsNotExist(err)
}

func isOptionalDependency(pkgName string, dir string) bool {
	if dir == "default" {
		return utils.Contains(maps.Keys(utils.GetPackageJSON().OptionalDependencies), pkgName)
	} else {
		pkgJSON, _ := utils.GetPackageJSONOfPackage(dir)
		return utils.Contains(maps.Keys(pkgJSON.OptionalDependencies), pkgName)
	}
}
