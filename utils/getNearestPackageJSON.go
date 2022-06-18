package utils

import (
	"fmt"
	"os"
)

func GetNearestPackageJSON() (path string, found bool, err error) {
	// Get the current directory
	dir, err := os.Getwd()
	loopDir := dir
	if err != nil {
		return "", false, err
	}

	attempts := 0
	// Get the nearest package.json
	for {
		attempts++

		// When the loop starts, this will lookup in the current directory
		if _, err := os.Stat(loopDir + "/package.json"); err == nil {
			return loopDir + "/package.json", true, nil
		}
		// If the file is not found, we will try to look in the parent directory (the loop will continue)
		loopDir = loopDir + "/.."

		// We should not look more than 15 directories up
		if attempts > 15 {
			return dir, false, fmt.Errorf("could not find package.json after 15 attempts")
		}
	}
}
