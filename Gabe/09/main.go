package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)

var minInt = int(^uint(0) >> 1)
var maxInt = -minInt -1
var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")

func programPartTwo(input []string, wantedValue int) int{
	sum := 0
	indexToRemove := 0
	for index := 0; index < len(input); index++{
		val, _ := strconv.Atoi(replacer.Replace(input[index]))
		sum += val
		// If sum is larger than wanted value, remove first value and redo last index
		if sum > wantedValue{
			tmp, _ := strconv.Atoi(replacer.Replace(input[indexToRemove]))
			sum -= (tmp + val)
			indexToRemove++
			index--
			continue
		}

		if sum == wantedValue{
			// Find min and max of values used
			min := minInt
			max := maxInt
			for i := indexToRemove; i <= index; i++{
				value, _ := strconv.Atoi(replacer.Replace(input[i]))
				if value < min{
					min = value
				}
				if value > max{
					max = value
				}
			}
			return min + max
		}
	}
	return 0
}

func programPartOne(input []string, preamble int) int{
	/*
	Check if any sum of two values within a subset of variables is NOT equal
	to the value after the subset
	*/
	for index := range input{
		subdata := input[index:(index +preamble)]
		tmp := replacer.Replace(input[index + preamble])
		valueToCheck,_ :=strconv.Atoi(tmp)
		ok := false
		out:
		for index2, j := range subdata{
			tmp := replacer.Replace(j)
			val1 ,_ := strconv.Atoi(tmp)
			for k := index2; k < len(subdata); k++{
				tmp = replacer.Replace(subdata[k])
				val2, _ := strconv.Atoi(tmp)
				if val1 == val2{
					continue
				}
				if (val1+val2) == valueToCheck{
					ok = true
					break out
				}
			}
		}
		if ok == false{
			return valueToCheck
		}
	}
	return 0
}


func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 9: Encoding Error ----")
	input := lib.ReadFileAndPutInSlice(path+"\\09\\input.txt")

	result := programPartOne(input, 25)
	println("Result Part One:",result)

	result = programPartTwo(input,result)
	println("Result Part Two:",result)
}
