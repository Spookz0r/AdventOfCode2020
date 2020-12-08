package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)

func programPartTwo(input []string) int{
	// BRUTE FORCE 
	var replacer = strings.NewReplacer("jmp", "nop", "nop", "jmp")
	sum := 0
	size := len(input)
	for index, i := range input{
		endIndex := 0
		orgString := i
		input[index] = replacer.Replace(i)
		sum, endIndex = programPartOne(input)
		// Restore to original string
		input[index] = orgString
		if endIndex == size{
			break
		}
	}
	return sum
}

func programPartOne(input []string) (int, int){
	var replacer = strings.NewReplacer("+", "", "\r", "", "\n", "")
	sum := 0
	indexes := map[int]bool{}
	index := 0
	size := len(input)
	for index < size{
		// Check if index already checked
		if _, ok := indexes[index]; ok{
			break
		}
		data := strings.Split(replacer.Replace(input[index])," ")
		val, _ := strconv.Atoi(data[1])
		// add index to map
		indexes[index] = true

		switch data[0] {
		case "acc":
			sum += val
			fallthrough
		case "nop":
			index++
			break
		case "jmp":
			index += val
			break
		}
	}
	return sum, index
}
func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 8: Handheld Halting ----")
	input := lib.ReadFileAndPutInSlice(path+"\\08\\input.txt")

	result, _ := programPartOne(input)
	println("Result Part One:",result)

	result = programPartTwo(input)
	println("Result Part Two:",result)
	
}
