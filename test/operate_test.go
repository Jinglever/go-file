package jgfile_test

import (
	"fmt"
	"os"
	"testing"

	jgfile "github.com/Jinglever/go-file"
	"github.com/bmizerany/assert"
)

func TestCopyFile(t *testing.T) {
	src := "./_data/test.csv"
	dst := "./_data/test_copy.csv"
	defer os.Remove(dst)

	fileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("获取文件信息失败：", err)
		return
	}
	fileSize := fileInfo.Size()

	nBytes, err := jgfile.CopyFile(src, dst)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, jgfile.IsFile(dst), true)
	assert.Equal(t, fileSize, nBytes)
}
