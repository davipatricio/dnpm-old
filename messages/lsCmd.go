package messages

import (
	"fmt"

	"github.com/davipatricio/colors/colors"
	"github.com/davipatricio/colors/styles"
)

func LsCmd(tree string) {
	fmt.Println(tree, "\n" + colors.Red("● ") + styles.Bold("Unmet Dependency\n")  + colors.Yellow("● ") + styles.Bold("Unmet Optional Dependency\n") + colors.Green("● ") + styles.Bold("Met Dependency\n"))
}