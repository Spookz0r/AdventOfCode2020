package main

import (
	"os"
	"fmt"
	"log"
	"github.com/Gabe/lib"
	"strings"
	"strconv"
)

// Replacer to have a uniform splitting of the input string
var replacer = strings.NewReplacer(" ", "-", ":", "")

func programPartOne(input []string) int{
	validPasswords := 0
	for _,i := range input {
		data := strings.Split(replacer.Replace(i), "-")
		occurrence := strings.Count(data[3],data[2])
		min, _ := strconv.Atoi(data[0])
		max, _ := strconv.Atoi(data[1])
		if (min <= occurrence) &&  (occurrence <= max) {
			validPasswords++
		}
	}
	return validPasswords
}

func programPartTwo(input []string) int{
	validPasswords := 0
	for _,i := range input {
		data := strings.Split(replacer.Replace(i), "-")
		pos1, _ := strconv.Atoi(data[0])
		pos2, _ := strconv.Atoi(data[1])
		if (string(data[3][pos1 - 1]) == data[2]) && (string(data[3][pos2 - 1]) != data[2]) ||
		   (string(data[3][pos1 - 1]) != data[2]) && (string(data[3][pos2 - 1]) == data[2]){
			validPasswords++
		}
	}
	return validPasswords
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	
	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 2: Password Philosophy ----")
	input := lib.ReadFileAndPutInSlice(path+"\\02\\input.txt")

	result := programPartOne(input)
	println("Result Part One:",result, "\n")

	result = programPartTwo(input)
	println("Result Part Two:",result, "\n")

}
