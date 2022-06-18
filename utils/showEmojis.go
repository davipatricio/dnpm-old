package utils

import (
	"flag"
)

var emojiPtr = flag.Bool("emoji", false, "Whether to show emojis on the output.")

func ShowEmojis() bool {
	flag.Parse()
	return *emojiPtr
}
