package file

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestIsDir(t *testing.T) {
	assert.Equal(t, true, IsDir("."))
	assert.Equal(t, false, IsDir("go.mod"))
	assert.Equal(t, false, IsDir("not_exists.go"))
}

func TestExists(t *testing.T) {
	assert.Equal(t, true, Exists("./check.go"))
	assert.Equal(t, false, Exists("./not_exists.go"))
}

func TestIsFile(t *testing.T) {
	assert.Equal(t, false, IsFile("."))
	assert.Equal(t, true, IsFile("go.mod"))
	assert.Equal(t, false, IsFile("not_exists.go"))
}
