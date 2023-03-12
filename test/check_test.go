package jgfile_test

import (
	"testing"

	jgfile "github.com/Jinglever/go-file"
	"github.com/bmizerany/assert"
)

func TestIsDir(t *testing.T) {
	assert.Equal(t, true, jgfile.IsDir("."))
	assert.Equal(t, false, jgfile.IsDir("check_test.go"))
	assert.Equal(t, false, jgfile.IsDir("not_exists.go"))
}

func TestExists(t *testing.T) {
	assert.Equal(t, true, jgfile.Exists("./check_test.go"))
	assert.Equal(t, false, jgfile.Exists("./not_exists.go"))
}

func TestIsFile(t *testing.T) {
	assert.Equal(t, false, jgfile.IsFile("."))
	assert.Equal(t, true, jgfile.IsFile("check_test.go"))
	assert.Equal(t, false, jgfile.IsFile("not_exists.go"))
}
