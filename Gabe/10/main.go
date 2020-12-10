package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/Gabe/lib"
)

func getValues(input []string) []int{
	//Convert whole list to ints...
	values := []int{}
	for _, i := range input{
		values = append(values, lib.ConvertToInt(i))
	}
	// Add start and end values
	values = append(values,0)
	sort.Ints(values)
	values = append(values,values[len(values)-1]+3)
	// fmt.Println(values)
	return values
}

func programPartTwo(input []string) int{
	
	values := getValues(input)
	// Array to hold how many combinations there are to reach each value
	// Each element will corresponds to the matching value in values
	possibleCombinationsPerIndex := []int{1}
	
	for index := 1; index < len(values); index++{
		possibleCombinations := 0
		// Go from start of list to current index
		for index2 := 0; index2 < index; index2++{
			// Check if it's possible to reach the next value from values[index2],
			// if so, add how many ways there are to reach that value to answer by using
			// possibleCombinationsPerIndex[index2].
			if (values[index2] +3) >= values[index]{
				// values[index] can be reached from values[index2]. Add how many ways there are
				// to reach values[index2] to ans.
				possibleCombinations += possibleCombinationsPerIndex[index2]
			}
		}

		// ans holds how many ways there are to reach this value.
		possibleCombinationsPerIndex = append(possibleCombinationsPerIndex, possibleCombinations)
	}
	// for i := range values{
	// 	fmt.Println("There are", possibleCombinationsPerIndex[i],"ways to get to",values[i])
	// }
	// result is the last value in solutions
	return possibleCombinationsPerIndex[len(possibleCombinationsPerIndex)-1]
}

func programPartOne(input []string) int{
	values := getValues(input)
	oneDiff, threeDiff, prevValue := 0, 0, 0
	for _, value := range values{
		if value - prevValue == 1{
			oneDiff++
		} else if value - prevValue == 3{
			threeDiff++
		}
		prevValue = value
	}
	return oneDiff * threeDiff
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 9: Encoding Error ----")
	testInput := lib.ReadFileAndPutInSlice(path+"\\10\\test_input.txt")
	testInput2 := lib.ReadFileAndPutInSlice(path+"\\10\\test_input2.txt")
	input := lib.ReadFileAndPutInSlice(path+"\\10\\input.txt")

	result := programPartOne(testInput)
	println("Result Test One:",result)

	result = programPartOne(testInput2)
	println("Result Test Two:",result)

	result = programPartOne(input)
	println("Result Part One:",result)

	result = programPartTwo(testInput)
	println("Result Test One Part Two:",result)
	
	result = programPartTwo(testInput2)
	println("Result Test Two Part Two:",result)

	result = programPartTwo(input)
	println("Result Part Two:",result)
}
