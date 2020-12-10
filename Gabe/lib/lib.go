package lib

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// MinInt min int value
var MinInt = int(^uint(0) >> 1)
// MaxInt max int value
var MaxInt = -MinInt -1


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

// ConvertToInt convert string to int, remove new line and spaces
func ConvertToInt(s string) int{
	var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")
	val, _ := strconv.Atoi(replacer.Replace(s))
	return val
}
