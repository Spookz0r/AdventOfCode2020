package main

import (
	"fmt"
)

func programPartOne(input []int, wantedNumber int)int{
	spokenNumbers := map[int][]int{}
	spokenNumber := 0
	for counter := 0; counter < wantedNumber; counter++{
		if counter < len(input){
			spokenNumbers[input[counter]] = append(spokenNumbers[input[counter]], counter)
			spokenNumber = input[counter]
		} else{
			if _, ok := spokenNumbers[spokenNumber]; !ok {
				spokenNumbers[spokenNumber] = append(spokenNumbers[input[counter]], counter)
				spokenNumber = 0
				continue
			}else if len(spokenNumbers[spokenNumber]) < 2 {
				spokenNumber = 0
			} else{
				nrs := len(spokenNumbers[spokenNumber])
				spokenNumber = spokenNumbers[spokenNumber][nrs-1]- spokenNumbers[spokenNumber][nrs-2]
			}
			spokenNumbers[spokenNumber] = append(spokenNumbers[spokenNumber], counter)
		}
	}
	return spokenNumber
}

func main() {
	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 15: Rambunctious Recitation ----")
	data := []int{0,3,6}
	result := programPartOne(data, 2020)
	println("Result Test One:",result)
	
	data = []int{2,1,10,11,0,6}
	result = programPartOne(data, 2020)
	println("Result Part One:",result)

	result = programPartOne(data, 30000000)
	println("Result Part Two:",result)
}