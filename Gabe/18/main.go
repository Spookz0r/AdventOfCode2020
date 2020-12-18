package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)
var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")

func getValuesAndOperators(expression string)([]int, []string){
	values := []int{}
	operators := []string{}
	tmpval := ""
	for index := 0; index < len(expression); index++{
		char := expression[index]
		switch char{
		case '*':
			v, _ := strconv.Atoi(tmpval)
			values = append(values, v)
			tmpval = ""
			operators = append(operators,string(char))
			break
		case '+':
			v, _ := strconv.Atoi(tmpval)
			values = append(values, v)
			tmpval = ""
			operators = append(operators,string(char))
			break
		default:
			tmpval = tmpval + string(char)
		}
	}
	v, _ := strconv.Atoi(tmpval)
	values = append(values, v)
	return values, operators
}

func calcExpression(values []int, operators []string)int{
	sum := values[0]
	// No priority, just go left to right
	for index := 0; index < len(operators); index++{
		val2 := values[index+1]
		operator := operators[index]
		if operator == "+"{
			sum += val2
		}
		if operator == "*"{
			sum *= val2
		}
	}
	return sum
}

func calcExpression2(values []int, operators []string)int{
	// Prioritize + before *. Find first +
	for index := 0; index < len(operators); index++{
		if operators[index] == "+"{
			tmpSum := values[index] + values[index+1]
			// Remove operator and replace values with sum
			values = append(values[:index], values[index+1:]...)
			values[index] = tmpSum
			operators = append(operators[:index], operators[index+1:]...)
			//restart evaluation with new list of values and operators
			index = 0
		}
	}
	// Now that all + has been prioritized and calculated, use func from part 1 for result
	return calcExpression(values, operators)
}

func program(input []string, part int)int{
	sum := 0
	for _, line := range input{
		line = replacer.Replace(line)
		var stack []int
		// First go through each parenthesis expression and replace with resulting value
		for index := 0; index < len(line); index++{
			char := line[index]
			if char == '('{
				stack = append(stack, index)
			}
			if char == ')'{
				startIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				values, operators := getValuesAndOperators(line[startIndex+1:index])
				var val int
				if part == 1{
					val = calcExpression(values, operators)
				} else if part == 2{
					val = calcExpression2(values, operators)
				}
				// Replace calculated expression with the result of the expression and re-loop
				newExpression := line[:startIndex] + strconv.Itoa(val) + line[index+1:]
				line = newExpression
				index = 0
			}
		}
		// Now all parenthesis are gone. Get values and operators and calculate
		values, operators := getValuesAndOperators(line)
		
		if part == 1{
			sum += calcExpression(values, operators)
		} else if part == 2{
			sum += calcExpression2(values, operators)
		}
	}
	return sum
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 18: Operation Order ----")
	testInput := lib.ReadFileAndPutInSlice(path+"\\18\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path+"\\18\\input.txt")

	result := program(testInput,1)
	println("Result Test One:",result)

	result = program(input,1)
	println("Result Part One:",result)

	result2 := program(testInput,2)
	println("Result Test Two:",result2)

	result2 = program(input,2)
	println("Result Part Two:",result2)
}
