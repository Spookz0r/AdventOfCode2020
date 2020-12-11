package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Gabe/lib"
)
var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")

func updateString(s string, char string, index int) string{
	output := s[:index] + char + s[index+1:]
	return output
}

func checkSeatsInLine(layout [][]string, row int, column int) int{
	occupiedCounter := 0
	for deltaX:= -1; deltaX < 2; deltaX++{ //row direction
		for deltaY := -1; deltaY < 2; deltaY++{ // column direction
			if deltaX == 0 && deltaY == 0{
				continue
			}
			x := row + deltaX
			y := column + deltaY
			a := 0
			for a < 1{
				if (0 <= x) &&  (x < len(layout)) && (0 <= y) && (y < len(layout[0]) && (layout[x][y] == ".")){
					x += deltaX
					y += deltaY
				} else{
					break
				}
			}
			if (0 <= x) &&  (x < len(layout)) && (0 <= y) && (y < len(layout[0])){
				if layout[x][y] == "#" {
					occupiedCounter++
				}
			}
		}
	}
	return occupiedCounter
}

func checkAdjacentSeats(layout [][]string, row int, column int) int{
	occupiedCounter := 0
	for i := row-1; i < row+2; i++{
		for j := column-1; j < column+2; j++{
			if i == row && j == column{
				continue
			}
			// check if index exist
			if i >= 0 && i < len(layout) && j >= 0 && j < len(layout[0]){
				if layout[i][j] == "#"{
					occupiedCounter++
				}
			}
		}
	}
	return occupiedCounter
}

func createLayout(input []string) [][]string{
	var layout [][]string
	for _, i := range input{
		var col []string
		line := replacer.Replace(i)
		for _, j := range line{
			col = append(col,string(j))
		}
		layout = append(layout,col)
	}
	fmt.Println("Layout Size: ",len(layout),len(layout[0]))
	return layout
}

func checkHowManyOccupiedSeats(layout [][]string) int{
	occupiedSeats := 0
	for i := 0; i < len(layout); i++{
		for j := 0; j < len(layout[0]); j++{
			if layout[i][j] == "#"{
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func programPartOneAndTwo(input []string, part int, nrOfNOTokOccupied int) int{
	// occupiedSeats := 0
	// Create a matrix of all the seats
	layout := createLayout(input)
	changed := true
	for changed == true{  // Break for loop if no change has been made
		changed = false
		// Copy layout and store the new values there.
		layoutCopy := make([][]string, len(input))
		for i:= range layoutCopy{
			layoutCopy[i] = make([]string,len(layout[i]))
			copy(layoutCopy[i],layout[i])
		}
		// // Go through each element
		for x := 0; x < len(layout); x++{
			for y := 0; y < len(layout[x]); y++{
				occupiedAdjacent := 0
				if part == 1{
					occupiedAdjacent = checkAdjacentSeats(layout, x,y)
				} else if part == 2{
					occupiedAdjacent = checkSeatsInLine(layout, x,y)
				}
				if layout[x][y] == "L" && occupiedAdjacent == 0{
					layoutCopy[x][y] = "#"
					changed = true
				}	else if layout[x][y] == "#" && occupiedAdjacent >= nrOfNOTokOccupied{
					layoutCopy[x][y] = "L"
					changed = true
				}
				
			}
		}
		
		// copy new layout to layout
		for i:= range layout{
			layout[i] = make([]string,len(layoutCopy[i]))
			copy(layout[i],layoutCopy[i])
		}
	}

	return checkHowManyOccupiedSeats(layout)
}


func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 11: Seating System ----")
	testInput := lib.ReadFileAndPutInSlice(path+"\\11\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path+"\\11\\input.txt")

	result := programPartOneAndTwo(testInput,1,4)
	println("Result Test One:",result)

	result = programPartOneAndTwo(input,1,4)
	println("Result Part One:",result)

	result = programPartOneAndTwo(testInput, 2,5)
	println("Result Test One Part Two:",result)

	result = programPartOneAndTwo(input, 2,5)
	println("Result Part Two:",result)
}
