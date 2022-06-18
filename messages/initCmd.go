package messages

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

// Could not read directory
func couldNotReadDirRaw() {
	fmt.Println(colors.Red("Could not read directory."))
}

func couldNotReadDirEmojis() {
	fmt.Println("❌ ", colors.Red("Could not read directory."))
}

// Package.json already exists
func initCmdExistsRaw() {
	fmt.Println(colors.Red("A package.json file already exists in this directory."))
}

func initCmdExistsEmojis() {
	fmt.Println("❌ ", colors.Red("A package.json file already exists in this directory."))
}

// Created package.json successfully
func initCmdDoneRaw() {
	fmt.Println(colors.Green("Created package.json successfully!"))
}

func initCmdDoneEmojis() {
	fmt.Println("✅ ", colors.Green("Created package.json successfully!"))
}

// Creating package.json
func initCmdCreatingRaw() {
	fmt.Println(colors.Green("Creating package.json on this directory..."))
}

func initCmdCreatingEmojis() {
	fmt.Println("📝 ", colors.Green("Creating package.json on this directory..."))
}
