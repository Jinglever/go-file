package jgfile

import (
	"encoding/json"
	"testing"

	"github.com/bmizerany/assert"
)

func TestReadCSVFile(t *testing.T) {
	file := "_test_data/test.csv"
	data, err := ReadCSVFile(file)
	if err != nil {
		t.Fatal(err)
	}
	str, _ := json.Marshal(data)
	assert.Equal(t,
		`[["","","","","","",""],["1","2","","","5","6","7"],["","","3","4","","",""]]`,
		string(str))
}
