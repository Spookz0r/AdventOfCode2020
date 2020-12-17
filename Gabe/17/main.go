package main

import (
	"fmt"
	"os"

	"github.com/Gabe/lib"
)

type coordinate struct{
	x int
	y int
	z int
	w int
}

func checkNeighbors(spaceCoordinates map[string]coordinate, thisKey string, coord coordinate, part int) (map[string]coordinate, int){
	tmpMap := map[string]coordinate{}
	activeNeighbors := 0
	for deltaW := -1 ; deltaW < 2; deltaW++{
		if part == 1{
			deltaW = 0
		}
		for deltaZ := -1 ; deltaZ < 2; deltaZ++{
			for deltaY := -1; deltaY < 2; deltaY++{
				for deltaX := -1; deltaX < 2; deltaX++{
					if deltaZ == 0 && deltaY == 0 && deltaX == 0 && deltaW == 0{
						continue
					}
					tX := coord.x + deltaX
					tY := coord.y + deltaY
					tZ := coord.z + deltaZ
					tW := coord.w + deltaW
					
					key := fmt.Sprintf("%d,%d,%d,%d", tX,tY,tZ, tW)
					// Key does not exist, check surrounding, if they add up to 3, add to map for next cycle
					if _, ok := spaceCoordinates[key]; !ok {
						count := 0
						for deltaW2 := -1 ; deltaW2 < 2; deltaW2++{
							for deltaZ2 := -1 ; deltaZ2 < 2; deltaZ2++{
								for deltaY2 := -1; deltaY2 < 2; deltaY2++{
									for deltaX2 := -1; deltaX2 < 2; deltaX2++{
										if deltaZ2 == 0 && deltaY2 == 0 && deltaX2 == 0 && deltaW2 == 0{
											continue
										}
										ttx := tX + deltaX2
										tty := tY + deltaY2
										ttz := tZ + deltaZ2
										ttw := tW + deltaW2
										key2 := fmt.Sprintf("%d,%d,%d,%d", ttx,tty,ttz, ttw)
										if _, ok := spaceCoordinates[key2]; ok {
											count++
										}
									}
								}
							}
						}
						if count == 3{
							tmpCoord := coordinate{
								tX,tY,tZ,tW,
							}
							tmpMap[key] = tmpCoord 
						}
					} else{
						activeNeighbors++
					}
				}
			}
		}
		if part == 1{
			break
		}
	}
	return tmpMap, activeNeighbors
}

func program(input []string, part int)int{
	// Create a map of all coordinates to easy check value of coordinate
	spaceCoordinates := map[string]coordinate{}
	
	for y, line := range input{
		z := 0
		w := 0
		for x := 0; x < len(line); x++{
			if line[x] == '#'{
				coordKey := fmt.Sprintf("%d,%d,%d,%d", x,y,z,w)
				coord := coordinate{
					x,y,z,w,
				}
				spaceCoordinates[coordKey] = coord
			}
		}
	}

	for cycle := 0; cycle < 6; cycle++{
		tmpMap := map[string]coordinate{}
		// Copy org map to work in tmp map
		tmpMap = spaceCoordinates
		for key, value := range spaceCoordinates{
			tmpCoord := coordinate{
				value.x, value.y, value.z, value.w,
			}
			tmpMap[key] = tmpCoord
		}

		for key, coord := range spaceCoordinates{
			updatedMap, activeNeighbors := checkNeighbors(spaceCoordinates, key, coord, part)
			for k,v := range updatedMap{
				if _, ok := tmpMap[k]; !ok {
					tmpMap[k] = v
				}
			}
			if activeNeighbors != 2 && activeNeighbors != 3{
				delete(tmpMap,key)
			}
		}
		// Copy the tmp map back to spaceMap
		spaceCoordinates = tmpMap
	}
	return len(spaceCoordinates)
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 17: Conway Cubes ----")
	testInput := lib.ReadFileAndPutInSlice(path+"\\17\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path+"\\17\\input.txt")

	result := program(testInput,1)
	println("Result Test One:",result)

	result = program(input,1)
	println("Result Part One:",result)

	result2 := program(testInput,2)
	println("Result Test Two:",result2)

	result2 = program(input,2)
	println("Result Part Two:",result2)
}
