package utils

import (
	"os"
)

func GetExecDir() string {
	dir, _ := os.Getwd()
	return dir
}