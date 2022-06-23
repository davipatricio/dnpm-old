package messages

import (
	"fmt"

	"github.com/gookit/color"
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
	fmt.Println(color.Red.Render("Could not read directory."))
}

func couldNotReadDirEmojis() {
	fmt.Println("❌ ", color.Red.Render("Could not read directory."))
}

// Package.json already exists
func initCmdExistsRaw() {
	fmt.Println(color.Red.Render("A package.json file already exists in this directory."))
}

func initCmdExistsEmojis() {
	fmt.Println("❌ ", color.Red.Render("A package.json file already exists in this directory."))
}

// Created package.json successfully
func initCmdDoneRaw() {
	fmt.Println(color.Green.Render("Created package.json successfully!"))
}

func initCmdDoneEmojis() {
	fmt.Println("✅ ", color.Green.Render("Created package.json successfully!"))
}

// Creating package.json
func initCmdCreatingRaw() {
	fmt.Println(color.Green.Render("Creating package.json on this directory..."))
}

func initCmdCreatingEmojis() {
	fmt.Println("📝 ", color.Green.Render("Creating package.json on this directory..."))
}
