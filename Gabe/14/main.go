package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)
var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")
var replacer2 = strings.NewReplacer("mem[", "", "]", "")

func createAllMasks(n int, arr [35]int, i int){
	if i == n{
		combos = append(combos, arr)
		return
	}
	arr[i] = 0
	createAllMasks(n,arr,i+1)
	arr[i] = 1
	createAllMasks(n,arr,i+1)

}
var combos = [][35]int{}

func programPartTwo(input []string)int64{
	sum := int64(0)

	memory := map[string]int64{}
	for index := 0; index < len(input); index++{
		combos = [][35]int{}
		data := input[index]
		mask := replacer.Replace(strings.Split(data," = ")[1])
		mask1Int, _ := strconv.ParseInt(strings.Replace(mask,"X","0",-1),2,64)
		
		// Find index of all Xs
		floatIndexes := []int{}
		for i, char := range mask{
			if char == 'X'{
				floatIndexes = append(floatIndexes, i)
			}
		}
		// Put all combinations in combos (global)
		tmpMask1 := [35]int{}
		createAllMasks(len(floatIndexes),tmpMask1,0)
		allMasksList := []map[int]int{}
		//Match combination with index.
		for _, combo := range combos{
			tmp := map[int]int{}
			for i := 0; i < len(floatIndexes); i++{
				tmp[floatIndexes[i]] = combo[i]
			}
			allMasksList = append(allMasksList, tmp)
		}

		for index = index+1; index < len(input); index++{
			// If next index is a mask, break and redo
			if strings.Contains(input[index],"mask"){
				index--
				break
			}

			data := strings.Split(replacer.Replace(input[index]), "=")
			value, _ := strconv.Atoi(data[1])
			memAddress, _ := strconv.Atoi(replacer2.Replace((data[0])))
			maskedOnce := int64(memAddress) | mask1Int
			for _, m := range allMasksList{
				memAddressTmp := maskedOnce
				for key, val := range m{
					if val == 1{ // Set bit
						memAddressTmp |= (1 << (35-key))
					}
					if val == 0{ // Clear bit
						memAddressTmp &= ^(1 << (35-key))
					}
					memory[strconv.Itoa(int(memAddressTmp))] = int64(value)
				}
			}
		}
	}
	for _, val := range memory{
		sum += val
	}

	return sum
}

func programPartOne(input []string)int64{
	sum := int64(0)
	memory := map[string]int64{}
	for index := 0; index < len(input); index++{
		data := input[index]
		mask := replacer.Replace(strings.Split(data," = ")[1])
		mask1Int, _ := strconv.ParseInt(strings.Replace(mask,"X","0",-1),2,64)
		mask2Int, _ := strconv.ParseInt(strings.Replace(mask,"X","1",-1),2,64)

		for index = index+1; index < len(input); index++{
			// If next index is a mask, break and redo
			if strings.Contains(input[index],"mask"){
				index--
				break
			}
			data := strings.Split(replacer.Replace(input[index]), "=")
			value, _ := strconv.Atoi(data[1])
			memAddress := replacer2.Replace((data[0]))
			// First OR with mask, replace X with 0.
			result := int64(value) | mask1Int
			//  AND with result, replace X with 1
			result = result & mask2Int
			memory[memAddress] = result
		}
	}
	for _, val := range memory{
		sum += val
	}

	return sum
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 14: Docking Data ----")
	testInput := lib.ReadFileAndPutInSlice(path+"\\14\\test_input.txt")
	testInput2 := lib.ReadFileAndPutInSlice(path+"\\14\\test2_input.txt")
	input := lib.ReadFileAndPutInSlice(path+"\\14\\input.txt")

	result := programPartOne(testInput)
	println("Result Test One:",result)

	result = programPartOne(input)
	println("Result Part One:",result)

	result2 := programPartTwo(testInput2)
	println("Result Test Two:",result2)

	result2 = programPartTwo(input)
	println("Result Part Two:",result2)
}
