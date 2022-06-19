package utils

import "os"

// Create a symlink at the provided path
func CreateSymlink(src, dst string) error {
	return os.Symlink(src, dst)
}

// Remove the symlink at the provided path
func RemoveSymlink(dst string) error {
	return os.Remove(dst)
}

// Check if the provided path is a symlink
func IsSymlink(src string) (bool, error) {
	fi, err := os.Lstat(src)
	if err != nil {
		return false, err
	}
	return fi.Mode()&os.ModeSymlink != 0, nil
}
