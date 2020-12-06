package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/Gabe/lib"
)


func getMaxVal(min int, max int) int{
	return min + ((max - min) / 2 )
}

func getMinVal(min int, max int) int{
	return max - ((max - min) / 2 )
}

func getSeatID(boardingPass string) int{
	minRow := 0
	maxRow := 127
	minColumn := 0
	maxColumn := 7
	for index, char := range boardingPass{
		if index < 7{

			if string(char) == "F" {
				maxRow = getMaxVal(minRow, maxRow)
			}
			if string(char) == "B"{
				minRow = getMinVal(minRow, maxRow)
			}
		} else{
			if string(char) == "L" {
				maxColumn = getMaxVal(minColumn, maxColumn)
			}
			if string(char) == "R"{
				minColumn = getMinVal(minColumn, maxColumn)
			}
		}
	}
	seatID := minRow * 8 + minColumn
	return seatID
}

func programPartOne(input []string) int{
	highestSeatID := 0
	for _, boardingPass := range input{
		seatID := getSeatID(boardingPass)

		if seatID > highestSeatID{
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func programPartTwo(input []string) int{
	allSeatIDs := []int{}
	// Get all Seat IDs and add to list
	for _, boardingPass := range input{
		seatID := getSeatID(boardingPass)
		allSeatIDs = append(allSeatIDs, seatID)
	}

	// Sort list and then check when there's a gap
	sort.Ints(allSeatIDs)
	index := 0
	for index = 1; index < len(allSeatIDs)-1; index++{
		if (allSeatIDs[index-1] != allSeatIDs[index] - 1){
			break
		}
	}
	
	return allSeatIDs[index] -1
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 5: Binary Boarding ----")
	input := lib.ReadFileAndPutInSlice(path+"\\05\\input.txt")

	result := programPartOne(input)
	println("Result Part One:",result)

	result = programPartTwo(input)
	println("Result Part Two:",result)

}
