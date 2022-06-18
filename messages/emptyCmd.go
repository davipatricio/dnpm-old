package msgs

import (
	"dnpm/utils"
	"fmt"

	"github.com/davipatricio/colors/colors"
	"github.com/davipatricio/colors/styles"
)

func EmptyCmd() {
	if utils.ShowEmojis() {
		emptyCmdEmojis()
	} else {
		emptyCmdRaw()
	}
}

func emptyCmdRaw() {
	fmt.Println(colors.Red("\nYou must specify a command."))
	fmt.Printf("dnpm <command>\n\nUsage:\n\n")

	fmt.Println(styles.Bold("dnpm install"), "         install all the dependecies in your project")
	fmt.Println(styles.Bold("dnpm install <pkg>"), "   add the <pkg> dependecy to your project")
	fmt.Println(styles.Bold("dnpm version"), "         shows the version of dnpm")
}

func emptyCmdEmojis() {
	fmt.Println(colors.Red("\n❌ You must specify a command."))
	fmt.Printf("dnpm <command>\n\nUsage:\n\n")

	fmt.Println(colors.Green("➕"), styles.Bold("dnpm install"), "         install all the dependecies in your project")
	fmt.Println(colors.Green("➕"), styles.Bold("dnpm install <pkg>"), "   add the <pkg> dependecy to your project")
	fmt.Println(colors.Cyan(" ℹ️"), styles.Bold("dnpm version"), "         shows the version of dnpm")
}
