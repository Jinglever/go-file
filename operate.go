package jgfile

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) (nBytes int64, err error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		err = fmt.Errorf("%s is not a regular file", src)
		return
	}

	source, err := os.Open(src)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := destination.Close()
		if err == nil {
			err = cerr
		}
	}()
	nBytes, err = io.Copy(destination, source)
	return
}
