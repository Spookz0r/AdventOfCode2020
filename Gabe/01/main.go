package main

import (
	"sort"
	"os"
	"fmt"
	"log"
	"strconv"
	"github.com/Gabe/lib"
)

func programPartOne(input []int) int{
	// Brute force
	iterationCounter := 0
	maxIndex := len(input)-1
	var result int
	out:
	for index := 0; index < maxIndex; index++ {
		value1 := input[index]
		if value1 < 2020{
			for index2 := index + 1; index2 < maxIndex; index2++{
				value2 := input[index2+1]
				sum := value1 + value2
				if sum == 2020{
					result = value1 * value2
					break out
				}
				iterationCounter++
			}
		}
	}
	println("Iterations: ", iterationCounter)
	return result
}

func programPartTwo(input []int) int{
	// Brute force
	iterationCounter := 0
	maxIndex := len(input)-1
	var result int
	out:
	for index := 0; index < maxIndex; index++ {
		value1 := input[index]
		if value1 < 2020{
			for index2 := index + 1; index2 < maxIndex-1; index2++{
				value2 := input[index2+1]
				for index3 := index2 + 1; index3 < maxIndex; index3++{
					value3 := input[index3+1]
					sum := value1 + value2 + value3
					if sum == 2020{
						result = value1 * value2 * value3
						break out
					}
					iterationCounter++
				}
			}
		}
	}
	println("Iterations: ", iterationCounter)
	return result
}

// Alternative solution, less brute force
type Slice struct{
	sort.IntSlice
	idx []int
}

// comment
func (s Slice) Swap(i, j int){
	s.IntSlice.Swap(i,j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

// comment
func NewSlice(n sort.IntSlice) *Slice {
	s := &Slice{IntSlice: n, idx: make([]int, n.Len())}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}
func programPartOneAlt(input []int) int{
	// Sort list but keep indexes, not necessary but whatever
	iterationCounter := 0
	s := NewSlice(sort.IntSlice(input))
	sort.Sort(s)
	maxIndex := len(s.IntSlice)-1
	for index1 := 0; index1 < maxIndex; index1++{
		// Fetch value from lowest to highest
		value1 := s.IntSlice[index1]
		for index2 := maxIndex-1; index2 > index1; index2--{
			// fetch value from highest to lowest
			value2 := s.IntSlice[index2]
			if (value1 +value2) < 2020{
				iterationCounter++
				break
			} else if (value1 +value2) == 2020{
				println("Iterations One Alternative: ", iterationCounter)
				return value1*value2
			}
		}
	}
	return 0
}
func programPartTwoAlt(input []int) int{
	// Sort list but keep indexes, not necessary but whatever
	iterationCounter := 0
	s := NewSlice(sort.IntSlice(input))
	sort.Sort(s)
	maxIndex := len(s.IntSlice)-1
	for index1 := 0; index1 < maxIndex; index1++{
		// Fetch value from lowest to highest
		value1 := s.IntSlice[index1]
		for index2 := maxIndex-1; index2 > index1; index2--{
			// fetch value from highest to lowest
			value2 := s.IntSlice[index2]
			for index3 := index2-1; index3 > index1; index3--{
				value3 := s.IntSlice[index3]
				if (value1 +value2 + value3) < 2020{
					iterationCounter++
					break
				} else if (value1 +value2 + value3) == 2020{
					println("Iterations Two Alternative: ", iterationCounter)
					return value1*value2*value3
				}
			}
		}
	}
	return 0
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	
	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 1: Report Repair ----")
	slc := lib.ReadFileAndPutInSlice(path+"\\01\\input.txt")

	// Convert String slice to int slice
	var data = []int{}
	for _,i := range slc {
		j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
		}
		data = append(data,j)
	}

	result := programPartOne(data)
	println("Result Part One:",result, "\n")  // took 0 nanoseconds apparently :D

	result2 := programPartTwo(data)
	println("Result Part Two:",result2,"\n") // took 6001200 nanoseconds... lame
	
	result1adv := programPartOneAlt(data)
	println("Result Part One Alternative:",result1adv, "\n") // took 0 nanoseconds apparently :D

	result2adv := programPartTwoAlt(data)
	println("Result Part Two Alternative:",result2adv,"\n") // took 0 nanoseconds apparently :D

}
