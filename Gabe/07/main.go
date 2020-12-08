package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)

type bagStruct struct{
	containsMyBag bool
	// contains []string
}

var bags = map[string]bagStruct{}

func programPartTwo(input []string, bag string) int{
	result := 0
	reg, _ := regexp.Compile("[^a-zA-Z ]+")
	for _, i := range input{
		data := strings.Split(i," bags contain ")
		if data[0] == bag{
			contains := strings.Split(data[1], ", ")
			for _, j := range contains{
				nr , _ := strconv.Atoi(string(j[0]))
				bag := strings.TrimSpace(reg.ReplaceAllString(j, ""))
				bag = strings.TrimSpace(strings.Split(bag, "bag")[0])
				res := programPartTwo(input, bag)
				result += nr * res	+ nr
			}

		}
	}
	return result
}

func hasMyBag(input []string, bag string) bool{
	reg, _ := regexp.Compile("[^a-zA-Z ]+")
	myBag := "shiny gold"
	result := false
	if _, ok := bags[bag]; ok{
		return bags[bag].containsMyBag
	}
	for _, i := range input{
		data := strings.Split(i," bags contain ")
		if data[0] == bag{
			tmp := bagStruct{false}
			bags[data[0]] = tmp
			contains := strings.Split(data[1], ", ")
			// Go through each contained bag
			for _, j := range contains{
				b := strings.TrimSpace(reg.ReplaceAllString(j, ""))
				b = strings.TrimSpace(strings.Split(b, "bag")[0])
				if b == myBag{
					tmp.containsMyBag = true
					bags[data[0]] = tmp
					return true
				}
				result = hasMyBag(input, b)
				// If bag found, break for loop and return true
				if result == true{
					tmp.containsMyBag = true
					bags[data[0]] = tmp
					break
				}
				
			}
			break
		}
		if data[0] == myBag{
			continue
		}
	}
	return result
} 

func programPartOne(input []string) int{
	sum := 0
	myBag := "shiny gold"
	for _, i := range input{
		data := strings.Split(i," bags contain ")
		if data[0] == myBag{
			continue
		}

		result:= hasMyBag(input,data[0])
		if result == true{
			sum++
		}
	}
	return sum
}
func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 7: Handy Haversacks ----")
	input := lib.ReadFileAndPutInSlice(path+"\\07\\input.txt")

	result := programPartOne(input)
	println("Result Part One:",result)

	result = programPartTwo(input, "shiny gold")
	println("Result Part Two:",result)
	
}
