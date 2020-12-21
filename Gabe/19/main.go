package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)

var replacer = strings.NewReplacer("\n", "", "\r", "")

var rules = map[string][]string{}
var col = map[string]string{}
var dp = map[string]bool{}

func matchList(msg string, start int, end int, rules2 string) bool {
	rules2 = strings.TrimSpace(rules2)
	if start == end && len(rules2) == 0 {
		return true
	} else if start == end {
		return false
	} else if len(rules2) == 0 {
		return false
	}

	ret := false
	r := strings.Split(rules2, " ")
	r2 := strings.Join(r[1:], " ")
	for index := start + 1; index < end+1; index++ {
		a := match(msg, start, index, string(r[0]))
		if a == false {
			continue
		}

		b := matchList(msg, index, end, r2)
		if a == true && b == true {
			return true
		}
	}
	return ret

}

func match(msg string, start int, end int, rule string) bool {
	ret := false
	// Check if key combination is already checked, if so return value
	key := strconv.Itoa(start) + "," + strconv.Itoa(end) + "," + rule
	if _, ok := dp[key]; ok {
		return dp[key]
	}

	if _, ok := col[rule]; ok {
		if msg[start:end] == col[rule] {
			ret = true
		}
	} else {
		for _, option := range rules[rule] {
			if matchList(msg, start, end, option) == true {
				ret = true
			}
		}
	}
	dp[key] = ret
	return ret
}

func program(input []string, part int) int {
	// Split input in rules and messages
	counter := 0

	for _, data := range input {
		if strings.Contains(data, ":") == true {
			// Rule
			data = replacer.Replace(data)
			r := strings.Split(data, " ")
			name := r[0][:len(r[0])-1]
			rest := ""
			if name == "8" && part == 2 {
				rest = "42 | 42 8"
			} else if name == "11" && part == 2 {
				rest = "42 31 | 42 11 31"
			} else {
				rest = strings.Join(r[1:], " ")
			}
			if strings.Contains(rest, "\"") == true {
				col[name] = rest[1 : len(rest)-1]
			} else {
				options := strings.Split(rest, " | ")
				tmpList := []string{}
				for _, opt := range options {
					tmpList = append(tmpList, opt)
				}
				rules[name] = tmpList
			}
		} else if strings.Contains(data, "a") == true {
			dp = map[string]bool{}
			data = replacer.Replace(data)
			if match(data, 0, len(data), "0") == true {
				counter++
			}
		}
	}
	return counter
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 19: Monster Messages ----")
	// testInput := lib.ReadFileAndPutInSlice(path + "\\19\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path + "\\19\\input.txt")

	// result := program(testInput)
	// println("Result Test One:", result)

	result := program(input, 1)
	println("Result Part One:", result)

	// result2 := program(testInput,2)
	// println("Result Test Two:",result2)

	result2 := program(input, 2)
	println("Result Part Two:", result2)
}
