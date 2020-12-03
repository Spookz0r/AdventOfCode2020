package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Gabe/lib"
)

var wg sync.WaitGroup
var treeCounterGlob int

func programPartOne(input []string, right int, down int,) {
	maxHeight := len(input)
	width := len(input[0]) - 1
	treeCounter := 0

	// Continue until bottom is reached
	sideIndex := right
	for index := down; index < maxHeight; index += down{
		if input[index][sideIndex % width] == '#'{
			treeCounter++
		}
		sideIndex += right
	}
	treeCounterGlob *= treeCounter
	defer wg.Done()
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 3: Toboggan Trajectory ----")
	input := lib.ReadFileAndPutInSlice(path+"\\03\\input.txt")

	wg.Add(1)
	treeCounterGlob = 1
	go programPartOne(input,3,1)
	wg.Wait()
	println("Result Part One:",treeCounterGlob)

	// Part two
	treeCounterGlob = 1
	testPoints := [][2]int{
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}
	// Probably takes longer time with go subroutines but I wanted to try them :D 
	for index := 0; index < len(testPoints); index++ {
		wg.Add(1)
		go programPartOne(input, testPoints[index][0], testPoints[index][1])
	}
	wg.Wait()

	println("Result Part Two:",treeCounterGlob)
}
