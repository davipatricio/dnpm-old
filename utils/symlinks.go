package utils

import "os"

func CreateSymlink(src, dst string) error {
	return os.Symlink(src, dst)
}

func RemoveSymlink(dst string) error {
	return os.Remove(dst)
}

func IsSymlink(src string) (bool, error) {
	fi, err := os.Lstat(src)
	if err != nil {
		return false, err
	}
	return fi.Mode()&os.ModeSymlink != 0, nil
}
