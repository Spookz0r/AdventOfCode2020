package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func program(input string, nrOfCups int, iterations int, part int) {
	cups := ring.New(nrOfCups)
	// Map for pointer of all values in ring to save ALOT of time
	cupsMap := make(map[int]*ring.Ring, nrOfCups)

	for _, x := range input {
		val, _ := strconv.Atoi(string(x))
		cups.Value = val
		cupsMap[cups.Value.(int)] = cups
		cups = cups.Next()
	}
	if part == 2 {
		// Now add values from length of input + 1 to length of ring
		for i := len(input) + 1; i <= nrOfCups; i++ {
			cups.Value = i
			cupsMap[cups.Value.(int)] = cups
			cups = cups.Next()
		}
	}

	for i := 0; i < iterations; i++ {
		currentCup := cups.Value.(int)
		// Remove the three following values and store them in a separate ring
		threeCups := cups.Unlink(3)
		// Now search for the wanted value which is current cup - 1
		wantedValue := currentCup
		for {
			wantedValue--
			if wantedValue < 1 {
				wantedValue = nrOfCups
			}
			// Check if wanted value is in the three removed cups
			isRemoved := false
			threeCups.Do(func(v interface{}) {
				if v.(int) == wantedValue {
					isRemoved = true
				}
			})
			// If the are not removed break while loop and use cupsMap to add cups after destination
			if isRemoved == false {
				break
			}
		}
		// Put now the three picked up cups back after destination
		cupsMap[wantedValue].Link(threeCups)
		// Go to the cup after current value
		cups = cups.Next()
	}
	cups = cupsMap[1]

	if part == 1 {
		res := ""
		cups.Do(func(v interface{}) {
			if v.(int) != 1 {
				res += strconv.Itoa(v.(int))
			}
		})
		fmt.Println("Result part one:", res)
	} else if part == 2 {
		// Take the two cups after 1 and multiply them
		fmt.Println("Result part two:", cups.Next().Value.(int), "*", cups.Next().Next().Value.(int), "=", cups.Next().Value.(int)*cups.Next().Next().Value.(int))
	}
}

func main() {
	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 23: Crab Cups ----")
	// program("389125467", 9, 100, 1)
	// programTwo("389125467", 1000000, 10000000, 2)
	program("614752839", 9, 100, 1)
	program("614752839", 1000000, 10000000, 2)
}
