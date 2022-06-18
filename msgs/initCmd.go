package msgs

import (
	"fmt"

	"github.com/davipatricio/colors/colors"
)

func InitCmd(showEmojis bool) {
	if showEmojis {
		initCmdCreatingEmojis()
	} else {
		initCmdCreatingRaw()
	}
}

func InitExistsCmd(showEmojis bool) {
	if showEmojis {
		initCmdExistsEmojis()
		initCmdExistsRaw()
	} else {
		initCmdExistsRaw()
	}
}

func InitDoneCmd(showEmojis bool) {
	if showEmojis {
		initCmdDoneEmojis()
	} else {
		initCmdDoneRaw()
	}
}

func InitErrReadingCmd(showEmojis bool) {
	if showEmojis {
		couldNotReadDirEmojis()
	} else {
		couldNotReadDirRaw()
	}
}

func couldNotReadDirRaw() {
	fmt.Println(colors.Red("Could not read directory."))
}

func couldNotReadDirEmojis() {
	fmt.Println("❌ ", colors.Red("Could not read directory."))
}

func initCmdExistsRaw() {
	fmt.Println(colors.Red("A package.json file already exists in this directory."))
}

func initCmdExistsEmojis() {
	fmt.Println("❌ ", colors.Red("A package.json file already exists in this directory."))
}

func initCmdDoneRaw() {
	fmt.Println(colors.Green("Created package.json successfully!"))
}

func initCmdDoneEmojis() {
	fmt.Println("✅ ", colors.Green("Created package.json successfully!"))
}

func initCmdCreatingRaw() {
	fmt.Println(colors.Green("Creating package.json on this directory..."))
}

func initCmdCreatingEmojis() {
	fmt.Println("📝 ", colors.Green("Creating package.json on this directory..."))
}
