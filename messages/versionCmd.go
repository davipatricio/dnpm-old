package messages

import (
	"dnpm/utils"
	"fmt"

	"github.com/davipatricio/colors/colors"
)

func VersionCmd(showEmojis bool) {
	if showEmojis {
		versionCmdEmojis()
	} else {
		versionCmdRaw()
	}
}

func versionCmdRaw() {
	fmt.Println(colors.Green("Version: ") + utils.VERSION)
}

func versionCmdEmojis() {
	fmt.Println(colors.Green("📦 Version: ") + utils.VERSION)
}
