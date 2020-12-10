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
	solutions := []int{1}
	
	for index := 1; index < len(values); index++{
		ans := 0
		// Go from start of list to current index
		for j := 0; j < index; j++{
			// Check if it's possible to reach the next value from values[j],
			// if so, add how many ways there are to reach that value to answer by using solutions[j].
			if (values[j] +3) >= values[index]{
				// values[index] can be reached from values[j]. Add how many ways there are
				// to reach values[j] to ans.
				ans += solutions[j]
			}
		}

		// ans holds how many ways there are to reach this value.
		solutions = append(solutions, ans)
	}
	// for i := range values{
	// 	fmt.Println("There are", solutions[i],"ways to get to",values[i])
	// }
	// result is the last value in solutions
	return solutions[len(solutions)-1]
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
