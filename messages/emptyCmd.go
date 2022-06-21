package messages

import (
	"dnpm/utils"
	"fmt"

	"github.com/gookit/color"
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
		color.Red.Println("\nYou must specify a command.")
	}
	fmt.Printf("dnpm <command>\n\nUsage:\n\n")

	fmt.Println(color.OpBold.Render("dnpm install"), "         install all the dependecies in your project")
	fmt.Println(color.OpBold.Render("dnpm install <pkg>"), "   add the <pkg> dependecy to your project")
	fmt.Println(color.OpBold.Render("dnpm version"), "         shows the version of dnpm")
}

func emptyCmdEmojis(notSpecified bool) {
	if notSpecified {
		color.Red.Println("\n❌ You must specify a command.")
	}
	fmt.Printf("dnpm <command>\n\nUsage:\n\n")

	fmt.Println("➕ ", color.OpBold.Render("dnpm install"), "         install all the dependecies in your project")
	fmt.Println("➕ ", color.OpBold.Render("dnpm install <pkg>"), "   add the <pkg> dependecy to your project")
	fmt.Println(color.Cyan.Render(" ℹ️"), color.OpBold.Render("dnpm version"), "         shows the version of dnpm")
}
