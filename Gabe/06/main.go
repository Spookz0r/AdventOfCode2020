package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Gabe/lib"
)



func programPartTwo(input []string) int{
	sum := 0
	for _,i := range input{
		persons := strings.Split(i, " ")
		nrOfPersons := len(persons)
		keys := make(map[string]int)
		for _, entry := range persons{
			for _,char := range entry{
				keys[string(char)]++
			}
		}
		for _, element := range keys{
			if element ==nrOfPersons{
				sum++
			}
		}
	}
	return sum
}

func programPartOne(input []string) int{
	sum := 0
	for _,i := range input{
		i = strings.Replace(i, " ","",-1)
		keys := make(map[string]bool)
		for _, entry := range i{
			char := string(entry)
			if _, value := keys[char]; !value{
				keys[char] = true
				sum++
			}
		}
	}
	return sum
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 6: Custom Customs ----")
	input := lib.ReadFileAndSplitAtEmptyLineAndPutInSlice(path+"\\06\\input.txt")

	result := programPartOne(input)
	println("Result Part One:",result)

	result = programPartTwo(input)
	println("Result Part Two:",result)

}
