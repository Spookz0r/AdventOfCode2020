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
	s := strings.Split(string(dat),"\n")
	return s
}

// ReadFileAndSplitAtEmptyLineAndPutInSlice Read file, split when there's an empty line
// and remove \n \r from line
func ReadFileAndSplitAtEmptyLineAndPutInSlice(path string) []string {
	dat, err := ioutil.ReadFile(path)
	Check(err)
	s := strings.Split(string(dat),"\n\r")
	for index, value := range s{
		value = strings.ReplaceAll(value,"\n"," ") 
		value = strings.ReplaceAll(value, "\r", " ")
		value = strings.ReplaceAll(value, "  ", " ")
		s[index] = strings.TrimSpace(value)
	}
	return s
}

// RemoveIndexFromList  Remove index from list
func RemoveIndexFromList(s []int, i int) []int{
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
