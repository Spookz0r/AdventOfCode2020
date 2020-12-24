package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)

var replacer = strings.NewReplacer("\n", "", "\r", "")

func countBlackTiles(tiles map[string]bool) int {
	blackCounter := 0
	for _, val := range tiles {
		if val {
			blackCounter++
		}
	}
	return blackCounter
}

func getNeighborCoordinates(coord string) []string {
	c := strings.Split(coord, ",")
	x, _ := strconv.Atoi(c[0])
	y, _ := strconv.Atoi(c[1])
	neighbors := []string{}
	neighbors = append(neighbors, strconv.Itoa(x+2)+","+strconv.Itoa(y))   // e
	neighbors = append(neighbors, strconv.Itoa(x+1)+","+strconv.Itoa(y-1)) // se
	neighbors = append(neighbors, strconv.Itoa(x-1)+","+strconv.Itoa(y-1)) // sw
	neighbors = append(neighbors, strconv.Itoa(x-2)+","+strconv.Itoa(y))   // w
	neighbors = append(neighbors, strconv.Itoa(x-1)+","+strconv.Itoa(y+1)) // nw
	neighbors = append(neighbors, strconv.Itoa(x+1)+","+strconv.Itoa(y+1)) // ne
	return neighbors
}

func program(input []string) {
	// True = black, false = white
	tiles := map[string]bool{}
	for _, line := range input {
		line = replacer.Replace(line)
		// Find coordinate for line, add coordinate to dict and flip
		// e, se, sw, w , nw, ne
		tmpDir := ""
		x := 0
		y := 0
		for _, c := range line {
			direction := ""
			if (c == 'e' || c == 'w') && tmpDir == "" {
				direction = string(c)
			} else if len(tmpDir) == 1 {
				direction = tmpDir + string(c)
				tmpDir = ""
			} else {
				tmpDir = string(c)
			}
			if direction != "" {
				switch direction {
				case "e":
					x += 2
					break
				case "se":
					x++
					y--
					break
				case "sw":
					x--
					y--
					break
				case "w":
					x -= 2
					break
				case "nw":
					x--
					y++
					break
				case "ne":
					x++
					y++
					break
				}
			}
		}
		coordinate := strconv.Itoa(x) + "," + strconv.Itoa(y)
		tiles[coordinate] = !tiles[coordinate] // false by default, toggle it
	}

	fmt.Println("Result Part One:", countBlackTiles(tiles))

	// Part two

	// Add Neighbors to all added coordinates if there are non
	for key := range tiles {
		neighbors := getNeighborCoordinates(key)
		for _, n := range neighbors {
			// Add neighbor as white if it does not exist in map
			if _, ok := tiles[n]; !ok {
				tiles[n] = false
			}
		}
	}

	for i := 0; i < 100; i++ {
		coordinatesToUpdate := map[string]bool{}

		for coordinate, color := range tiles {
			neighbors := getNeighborCoordinates(coordinate)
			blackNeighbors := 0
			for _, n := range neighbors {
				if _, ok := tiles[n]; ok {
					if tiles[n] {
						blackNeighbors++
					}
				} else {
					tiles[n] = false
				}
			}
			// If color is black and ZERO black around them or two and more black around => white
			if color == true && (blackNeighbors == 0 || blackNeighbors > 2) {
				coordinatesToUpdate[coordinate] = false
			} else if color == false && blackNeighbors == 2 {
				coordinatesToUpdate[coordinate] = true
			}
		}
		// Update tiles
		for key, val := range coordinatesToUpdate {
			tiles[key] = val
		}
		fmt.Println("After day", i+1, "nr of black tiles are:", countBlackTiles(tiles))
	}
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 24: Lobby Layout ----")
	// testInput := lib.ReadFileAndPutInSlice(path + "\\24\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path + "\\24\\input.txt")
	// program(testInput)
	program(input)
}
