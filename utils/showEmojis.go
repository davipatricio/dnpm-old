package utils

import (
	"flag"
	"fmt"
)

var emojiPtr = flag.Bool("emoji", false, "Whether to show emojis on the output.")

func ShowEmojis() bool {
	flag.Parse()
	fmt.Println(*emojiPtr)
	return *emojiPtr
}
