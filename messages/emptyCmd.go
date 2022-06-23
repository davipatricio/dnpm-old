package messages

import (
	"dnpm/utils"
	"fmt"

	"github.com/davipatricio/colors/colors"
	"github.com/davipatricio/colors/styles"
)

func EmptyCmd(notSpecified bool) {
	if utils.ShowEmojis() {
		emptyCmdEmojis(notSpecified)
	} else {
		emptyCmdRaw(notSpecified)
	}
}

func emptyCmdRaw(notSpecified bool) {
	if notSpecified {
		fmt.Println(colors.Red("\nYou must specify a command."))
	}
	fmt.Printf("dnpm <command>\n\nUsage:\n\n")

	fmt.Println(styles.Bold("dnpm install"), "         install all the dependecies in your project")
	fmt.Println(styles.Bold("dnpm install <pkg>"), "   add the <pkg> dependecy to your project")
	fmt.Println(styles.Bold("dnpm version"), "         shows the version of dnpm")
	fmt.Println(styles.Bold("dnpm ls"), "              shows all the versions of packages that are installed")
}

func emptyCmdEmojis(notSpecified bool) {
	if notSpecified {
		fmt.Println(colors.Red("\n❌ You must specify a command."))
	}
	fmt.Printf("dnpm <command>\n\nUsage:\n\n")

	fmt.Println(colors.Green("➕"), styles.Bold("dnpm install"), "         install all the dependecies in your project")
	fmt.Println(colors.Green("➕"), styles.Bold("dnpm install <pkg>"), "   add the <pkg> dependecy to your project")
	fmt.Println(colors.Cyan(" ℹ️"), styles.Bold("dnpm version"), "         shows the version of dnpm")
	fmt.Println(colors.Cyan(" ℹ️"), styles.Bold("dnpm ls"), "              shows all the versions of packages that are installed")
}
