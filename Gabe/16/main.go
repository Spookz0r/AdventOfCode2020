package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)
var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")

func program(input []string){
	rules := map[string][]int{}
	ticket := []int{}
	inputType := "rules"
	nearbyTickets := [][]int{}
	validNumbers := map[int]int{}
	invalidNumbersSum := 0

	for index, line := range input{
		// fmt.Println(index,line)
		if line == "your ticket:\r"{
			inputType = "my ticket"
			continue
		} else if line == "nearby tickets:\r"{
			inputType = "nearby tickets"
			continue
		}

		if line != "\r"{
			line = strings.Replace(line, "\r","",-1)
			switch inputType {
			case "rules":
				data := strings.Split(line,": ")
				values := strings.Split(strings.Replace(data[1],"\r","",-1)," or ")
				// Add each value of the range to the dict
				ruleType := data[0]
				for _, val := range values{
					val = replacer.Replace(val)
					tmp := strings.Split(val,"-")
					
					startRange, _ := strconv.Atoi(tmp[0])
					endRange, _ := strconv.Atoi(tmp[1])
					for j := startRange; j <= endRange; j++{
						rules[ruleType] = append(rules[ruleType], j)
						validNumbers[j] = j
					}
				}
				break
			case "my ticket":
				tmp := replacer.Replace(input[index])
				a := strings.Split(tmp,",")
				// add values to ticket
				for _, v := range a{
					val, _ := strconv.Atoi(v)
					ticket = append(ticket, val)
				}
				break
			case "nearby tickets":
				tmp := replacer.Replace(input[index])
				a := strings.Split(tmp,",")
				// add values to ticket
				tmp2 := []int{}
				validCounter := 0
				for _, v := range a{
					val, _ := strconv.Atoi(v)
					tmp2 = append(tmp2, val)
					if _, ok := validNumbers[val]; !ok {
						invalidNumbersSum += val
					} else{
						validCounter++
					}
				}
				if validCounter == len(a){
					nearbyTickets = append(nearbyTickets, tmp2)
				}

				break
			}
		}
	}

	matchingIndexRule := map[string][]int{}

	for key, rule := range rules{
		// Go through each nearby tickets i index and check if all exist in a class
		counter := 0
		for i := 0; i < len(nearbyTickets[0]); i++{
			counter = 0
			for j := 0; j < len(nearbyTickets); j++{
				val := nearbyTickets[j][i]
				// Check if value exist in rule
				for _, k := range rule{
					if k == val {
						counter++
						break
					}
				}
				if counter == len(nearbyTickets){
					matchingIndexRule[key] = append(matchingIndexRule[key], i)
				}
			}
		}
	}
	// Loop until each rule only has one index related to it
	done := false
	for done == false{
		nrWithLen1 := 0
		for key, val := range matchingIndexRule{
			if len(val) == 1{
				nrWithLen1++
				// Remove this val from every entry
				for key2, val2 := range matchingIndexRule{
					if key2 != key{
						for index := 0; index < len(val2); index++{
							if val2[index] == val[0]{
								matchingIndexRule[key2] = lib.RemoveIndexFromList(val2, index)
								break
							}
						}
					}
				}
			}
		}
		if nrWithLen1 == len(matchingIndexRule){
			done = true
		}
	}

	// Find Index for each Departure type and multiply ticket values
	result := 1
	for key, val := range matchingIndexRule{
		if strings.Contains(key, "departure"){
			result *= ticket[val[0]]
		}
	}
	fmt.Println("Result Part 1:",invalidNumbersSum)
	fmt.Println("Result Part 2:",result)
}


func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 16: Ticket Translation ----")
	input := lib.ReadFileAndPutInSlice(path+"\\16\\input.txt")

	program(input)
}
