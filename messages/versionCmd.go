package messages

import (
	"dnpm/utils"

	"github.com/gookit/color"
)

func VersionCmd(showEmojis bool) {
	if showEmojis {
		versionCmdEmojis()
	} else {
		versionCmdRaw()
	}
}

func versionCmdRaw() {
	color.Green.Println("Version: " + utils.VERSION)
}

func versionCmdEmojis() {
	color.Green.Println("📦 Version: " + utils.VERSION)
}
