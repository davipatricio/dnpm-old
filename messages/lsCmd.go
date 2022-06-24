package messages

import (
	"fmt"

	"github.com/gookit/color"
)

func LsCmd(tree string) {
	final := tree

	final += "\n" + color.Red.Render("● ") + color.Bold.Render("Unmet Dependency\n")
	final += color.Bold.Render("Unmet Optional Dependency\n") + color.Yellow.Render("● ")
	final += color.Green.Render("● ") + color.Bold.Render("Met Dependency\n")

	fmt.Println(final)
}
