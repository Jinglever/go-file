package jgfile

import "os"

// check if is dir
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// check if file exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

// check if is file
func IsFile(path string) bool {
	return Exists(path) && !IsDir(path)
}
