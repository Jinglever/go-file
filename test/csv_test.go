package jgfile_test

import (
	"encoding/json"
	"testing"

	jgfile "github.com/Jinglever/go-file"
	"github.com/bmizerany/assert"
)

func TestReadCSVFile(t *testing.T) {
	file := "_data/test.csv"
	data, err := jgfile.ReadCSVFile(file)
	if err != nil {
		t.Fatal(err)
	}
	str, _ := json.Marshal(data)
	assert.Equal(t,
		`[["","","","","","",""],["1","2","","","5","6","7"],["","","3","4","","",""]]`,
		string(str))
}
