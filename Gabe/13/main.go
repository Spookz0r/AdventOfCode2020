package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/Gabe/lib"
)
var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")

// Code from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
var one = big.NewInt(1)
func crt(a, n []*big.Int) (*big.Int, error) {
    p := new(big.Int).Set(n[0])
    for _, n1 := range n[1:] {
        p.Mul(p, n1)
	}
	var x, q, s, z big.Int
    for i, n1 := range n {
        q.Div(p, n1)
        z.GCD(nil, &s, n1, &q)
        if z.Cmp(one) != 0 {
            return nil, fmt.Errorf("%d not coprime", n1)
        }
        x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
    }
    return x.Mod(&x, p), nil
}

func programPartTwo(input string) string{
	busses := strings.Split(replacer.Replace(input),",")
	busIndexes := []*big.Int{}
	busNrs := []*big.Int{}
	for index := 0; index < len(busses); index++{
		if busses[index] != "x"{
			busNrs = append(busNrs, big.NewInt(int64(lib.ConvertToInt(busses[index]))))
			busIndexes = append(busIndexes, big.NewInt(int64(-index)))
		}
	}

	a,_ := crt(busIndexes,busNrs)
	return a.String()
}

func programPartOne(input []string)int{
	earliestTimestamp := lib.ConvertToInt(input[0])
	busses := strings.Split(replacer.Replace(input[1]),",")
	minWaitTime := 999999
	busID := 0
	for _, bus := range busses{
		if bus != "x"{
			busNr := lib.ConvertToInt(bus)
			waitTime := busNr - (earliestTimestamp % busNr)
			if waitTime < minWaitTime{
				minWaitTime = waitTime
				busID = busNr
			}
		}
	}
	return busID * minWaitTime
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 13: Shuttle Search ----")
	testInput := lib.ReadFileAndPutInSlice(path+"\\13\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path+"\\13\\input.txt")

	result := programPartOne(testInput)
	println("Result Test One:",result)

	result = programPartOne(input)
	println("Result Part One:",result)

	result2 := programPartTwo("17,x,13,19")
	println("Result Test Part Two:",result2, "should be",3417)
	result2 = programPartTwo("67,7,59,61")
	println("Result Test Part Two:",result2, "should be",754018)
	result2 = programPartTwo("67,x,7,59,61")
	println("Result Test Part Two:",result2,"should be",779210)
	result2 = programPartTwo("67,7,x,59,61")
	println("Result Test Part Two:",result2, "should be",1261476)
	result2 = programPartTwo("1789,37,47,1889")
	println("Result Test Part Two:",result2, "should be",1202161486)

	result2 = programPartTwo(input[1])
	println("Result Part Two:",result2)
}