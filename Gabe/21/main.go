package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Gabe/lib"
)

var replacer = strings.NewReplacer("\n", "", "\r", "", ",", "", ")", "")

type foodType struct {
	name       string
	mayContain map[string]int
}

var foodMap = map[string]foodType{}
var foodOccurence = map[string]int{}

func program(input []string) {

	for _, line := range input {
		data := strings.Split(line, " (contains ")
		foods := strings.Split(data[0], " ")
		allergens := strings.Split(replacer.Replace(data[1]), " ")
		for _, food := range foods {
			foodItem := foodType{}
			foodItem.name = food
			contains := map[string]int{}
			if _, ok := foodMap[food]; ok {
				foodItem = foodMap[food]
				contains = foodMap[food].mayContain
			}
			for _, allergen := range allergens {
				contains[allergen]++
			}
			foodItem.mayContain = contains
			foodMap[food] = foodItem
			foodOccurence[food]++
		}
	}
	// fmt.Println(foodOccurence)

	// Check if any food as the allergen multiple times
	allergen := map[string]map[string]int{}
	for _, val := range foodMap {
		for key2, val2 := range val.mayContain {
			// If allergen not yet added, add it with info from this
			if _, ok := allergen[key2]; !ok {
				tmpMap := map[string]int{val.name: val2}
				allergen[key2] = tmpMap
			} else {
				// Allergen exist so update map
				tmpMap := allergen[key2]
				tmpMap[val.name] = val2
				allergen[key2] = tmpMap
			}
		}
	}

	// Find max for each allergen
	for _, aller := range allergen {
		max := 0
		for _, val := range aller {
			if val > max {
				max = val
			}
		}
		// check which items that does not have max and remove them
		for key2, val := range aller {
			if val < max {
				delete(aller, key2)
			}
		}
	}
	// Add all foods to list, to then be removed
	var foodWithoutAllergen = map[string]int{}
	for key := range foodMap {
		foodWithoutAllergen[key] = 1
	}

	// check if length of an allergen is 1. If so remove that food from all other allergens
	done := false
	pairedFood := map[string]int{}

	foodToDelete := ""
	for done == false {
		counter := 0
		for a, aller := range allergen {
			if _, ok := pairedFood[a]; ok {
				continue
			}
			if len(aller) == 1 {
				for key := range aller { // Take the ONLY value in the map... so pretty
					foodToDelete = key
					delete(foodWithoutAllergen, foodToDelete)
					pairedFood[foodToDelete] = 1
				}
				counter++
			} else {
				// Remove from the rest
				if _, ok := aller[foodToDelete]; ok {
					delete(aller, foodToDelete)
				}
			}
		}
		if counter == len(allergen) {
			break
		}
	}
	// COunt how many times the foods without allergen exist
	result := 0
	for key, val := range foodOccurence {
		for key2 := range foodWithoutAllergen {
			if key == key2 {
				result += val
			}
		}
	}

	// Part 2. Create a list of allergens and sort alphabetically
	listOfAllergens := []string{}
	for key := range allergen {
		listOfAllergens = append(listOfAllergens, key)
	}
	sort.Strings(listOfAllergens)
	// Use sorted list of allergens to create the list of ingredients
	listOfIngredients := []string{}
	for _, key := range listOfAllergens {
		tmpMap := allergen[key]
		ingredient := ""
		for key := range tmpMap { // Take the ONLY value in the map... so pretty
			ingredient = key
		}
		listOfIngredients = append(listOfIngredients, ingredient)
	}
	result2 := strings.Join(listOfIngredients, ",")

	fmt.Println("Result part One:", result)
	fmt.Println("Result part Two:", result2)
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 21: Allergen Assessment ----")
	// testInput := lib.ReadFileAndPutInSlice(path + "\\21\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path + "\\21\\input.txt")
	// program(testInput)
	program(input)
}

// 20899048083289 too high
