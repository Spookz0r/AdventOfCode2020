package lib

import (
	"io/ioutil"
	"strings"
)

// Check if error when reading file
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFileAndPutInSlice reads file and return data in a slice
func ReadFileAndPutInSlice(path string) []string {
	dat, err := ioutil.ReadFile(path)
	Check(err)
	s := strings.Fields(string(dat))
	return s
}
