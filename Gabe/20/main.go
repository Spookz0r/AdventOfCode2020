package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)

var replacer = strings.NewReplacer("\n", "", "\r", "")

type tile struct {
	id     string
	matrix [][]string
	data   []string
	top    string
	bottom string
	left   string
	right  string

	inUse       bool
	version     [][][]string
	usedVersion int

	matchingIds        []string
	matchingdirections []int
	posX               int
	posY               int
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var objects = map[int]tile{}
var cornerIds = map[int]int{}

func copyMatrix(matrix [][]string) [][]string {
	newMatrix := make([][]string, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]string, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}

func rotateMatrix90ACW(mat [][]string) [][]string {
	newMatrix := copyMatrix((mat))
	for index, row := range mat {
		i := 0
		for index2 := len(row) - 1; index2 >= 0; index2-- {
			newMatrix[i][index] = row[index2]
			i++
		}
	}
	return newMatrix
}

func flipMatrix(mat [][]string) [][]string {
	newMatrix := copyMatrix((mat))
	for index, row := range mat {
		i := 0
		for index2 := len(row) - 1; index2 >= 0; index2-- {
			newMatrix[index][i] = row[index2]
			i++
		}
	}
	return newMatrix
}

func printMatrix(mat [][]string) {
	for _, v := range mat {
		fmt.Println(v)
	}
	fmt.Println(" ")
}

func getEdgesFromMatrix(matrix [][]string) []string {
	edges := []string{}
	//top
	top := strings.Join(matrix[0], "")
	bottom := strings.Join(matrix[len(matrix)-1], "")
	left := ""
	right := ""
	for _, row := range matrix {
		left += row[0]
		right += row[len(row)-1]
	}
	edges = append(edges, top)
	edges = append(edges, bottom)
	edges = append(edges, left)
	edges = append(edges, right)
	return edges
}

func program(input []string, part int) int {
	// Create all tile objects
	tmpObject := tile{}
	for index := 0; index < len(input); index++ {
		line := input[index]
		if strings.Contains(line, "Tile") == true {
			tmp := strings.Split(line, " ")
			id := tmp[1][:len(tmp[1])-2]
			idNr, _ := strconv.Atoi(id)
			tmpObject.id = id
			data := []string{}
			// fmt.Println("Id:", id)
			// done := false
			for index < len(input)-1 {
				index++
				line2 := input[index]
				if len(line2) < 4 { // New line
					break
				}
				line2 = replacer.Replace(line2)
				data = append(data, line2)
			}
			tmpObject.data = data
			tmpObject.inUse = false
			matrix := [][]string{}
			for _, row := range tmpObject.data {
				tmpArray := []string{}
				for _, i := range row {
					tmpArray = append(tmpArray, string(i))
				}
				matrix = append(matrix, tmpArray)
			}
			tmpObject.matrix = matrix
			// Create all 8 versions of the tile
			m := copyMatrix(matrix)
			for index := 0; index < 8; index++ {
				tmpObject.version = append(tmpObject.version, m)
				// Rotate for next round
				m = rotateMatrix90ACW(m)
				if index == 3 {
					// if 4 rotations done, flip the matrix and rotate again.
					m = flipMatrix(m)
				}
			}
			objects[idNr] = tmpObject
			tmpObject = tile{}
		}
	}
	// Add top, bottom, left and right in each object
	for key, obj := range objects {
		tmpObject = obj
		tmpObject.top = obj.data[0]
		tmpObject.bottom = obj.data[len(obj.data)-1]
		left := ""
		right := ""
		for _, a := range obj.data {
			left += string(a[0])
			right += string(a[len(a)-1])
		}
		tmpObject.left = left
		tmpObject.right = right

		objects[key] = tmpObject
	}

	// find corner pieces
	// if top/left, top/right, bottom/left or bottom/right are unique => corner piece
	for _, obj := range objects {
		matches := 0
		for _, obj2 := range objects {
			if obj.id == obj2.id {
				continue
			}
			counter := 0
			cmpList := []string{obj.top, obj.bottom, obj.left, obj.right}
			// Check if there are any matches.
			for _, cmp := range cmpList {
				cmp2 := reverse(cmp)
				if cmp == obj2.top || cmp == obj2.right || cmp == obj2.bottom || cmp == obj2.left ||
					cmp2 == obj2.top || cmp2 == obj2.right || cmp2 == obj2.bottom || cmp2 == obj2.left {
					counter++
				}
			}
			if counter >= 1 {
				matches += counter
			}

		}
		// 2 matches equals side piece, 3 matches equals corner
		if matches == 2 {
			id, _ := strconv.Atoi(obj.id)
			cornerIds[id] = id
		}
	}
	result := 1
	for _, id := range cornerIds {
		result *= id
	}
	return result
}

var dirs = map[int]string{0: "TOP", 1: "BOTTOM", 2: "LEFT", 3: "RIGHT"}

func checkNeighbors(tileToCheck tile, tileMatrix [][]string, row int, col int) ([][]string, int, int) {
	tileEdges := getEdgesFromMatrix(tileToCheck.version[tileToCheck.usedVersion])
out:
	for _, obj := range objects {
		if _, ok := tilesInUse[obj.id]; ok {
			continue
		}
		// fmt.Println("Checking against Tile:", obj.id)
		for index, version := range obj.version {
			tilesInUse[obj.id] = 1
			obj.usedVersion = index
			edges := getEdgesFromMatrix(version)

			for index1, cornEdge := range tileEdges {
				// If match top/bottom, left/right, bottom/top, right/left update matchid and direction.
				if (index1 == 0 && cornEdge == edges[1]) ||
					(index1 == 1 && cornEdge == edges[0]) ||
					(index1 == 2 && cornEdge == edges[3]) ||
					(index1 == 3 && cornEdge == edges[2]) {
					tileToCheck.matchingIds = append(tileToCheck.matchingIds, obj.id)
					tileToCheck.matchingdirections = append(tileToCheck.matchingdirections, index1)
					id, _ := strconv.Atoi(tileToCheck.id)
					objects[id] = tileToCheck
					id, _ = strconv.Atoi(obj.id)
					objects[id] = obj
					break out
				}
			}
			// Delete from tilesInUse if not found
			delete(tilesInUse, obj.id)
		}
	}
	return tileMatrix, row, col
}

func createPic() int {
	tileMatrix := make([][]string, 30)
	for i := 0; i < len(tileMatrix); i++ {
		tileMatrix[i] = make([]string, len(tileMatrix))
	}
	// Start with a corner tile
	corner := tile{}
	for _, a := range cornerIds {
		corner = objects[a]
	}
	corner.usedVersion = 0
	corner.posX = 15
	corner.posY = 15
	id, _ := strconv.Atoi(corner.id)
	tilesInUse[corner.id] = 1
	tileMatrix[corner.posY][corner.posX] = corner.id
	row := 0
	col := 0
	done := false
	objects[id] = corner
	// Check all neighbors, Done when all tiles are used.
	for done == false {
		for key := range tilesInUse {
			keyNr, _ := strconv.Atoi(key)
			// Check neigbor of active tile (keyNr)
			tileMatrix, row, col = checkNeighbors(objects[keyNr], tileMatrix, row, col)
		}
		if len(tilesInUse) == len(objects) {
			break
		}
	}

	// Build matrix from matchingId and matching direction. Start anywhere is a matrix larger than final matrix :/
	corner = objects[id]
	done = false
	placed := map[string]int{}
	placed[corner.id] = 1

	for done == false {
		for _, obj := range objects {
			if _, ok := placed[obj.id]; ok {
				for index, matchedId := range obj.matchingIds {
					id, _ = strconv.Atoi(matchedId)
					matchedObj := objects[id]
					if matchedObj.posX != 0 && matchedObj.posY != 0 {
						continue
					}
					dir := obj.matchingdirections[index]
					if dirs[dir] == "TOP" {
						tileMatrix[obj.posY+1][obj.posX] = matchedId
						matchedObj.posY = obj.posY + 1
						matchedObj.posX = obj.posX
					} else if dirs[dir] == "BOTTOM" {
						tileMatrix[obj.posY-1][obj.posX] = matchedId
						matchedObj.posY = obj.posY - 1
						matchedObj.posX = obj.posX
					} else if dirs[dir] == "LEFT" {
						tileMatrix[obj.posY][obj.posX-1] = matchedId
						matchedObj.posY = obj.posY
						matchedObj.posX = obj.posX - 1
					} else if dirs[dir] == "RIGHT" {
						tileMatrix[obj.posY][obj.posX+1] = matchedId
						matchedObj.posY = obj.posY
						matchedObj.posX = obj.posX + 1
					}
					objects[id] = matchedObj
					placed[matchedObj.id] = 1
				}
			}
		}
		if len(placed) == len(objects) {
			break
		}
	}

	// Remove border of each tile
	for _, a := range tileMatrix {
		if len(a[15]) != 0 {
			for _, x := range a {
				if len(x) != 0 {
					id, _ := strconv.Atoi(x)
					tmpObj := objects[id]
					matrix := tmpObj.version[tmpObj.usedVersion]
					newMatrix := [][]string{}
					for row := 1; row < len(matrix)-1; row++ {
						tmpRow := []string{}
						for col := 1; col < len(matrix[0])-1; col++ {
							tmpRow = append(tmpRow, matrix[row][col])
						}
						newMatrix = append(newMatrix, tmpRow)
					}
					newMatrix = rotateMatrix90ACW(newMatrix)
					newMatrix = rotateMatrix90ACW(newMatrix)
					newMatrix = flipMatrix(newMatrix)
					tmpObj.matrix = newMatrix
					objects[id] = tmpObj
				}
			}
		}
	}
	// Create one big matrix:
	matr := [][]string{}
	done = false
	// BUILD IT UP FROM LEFT TO RIGHT and LINE BY LINE
	for _, rowList := range tileMatrix {
		// I have to add all rows in EACH ROWLIST to the new Matrix
		index := 0
		done = false
		// loop until every line is taken
		for done == false {
			newListRow := [][]string{}
			add := false
			for _, colID := range rowList {
				id, _ := strconv.Atoi(colID)
				// CHeck if ID is an object, otherwise empty element
				if _, ok := objects[id]; ok {
					// get object to get index row of matrix
					row := objects[id].matrix[index]
					newListRow = append(newListRow, row)
					add = true
				}
			}
			if add == true {
				newRow := []string{}
				for _, a := range newListRow {
					newRow = append(newRow, a...)
				}
				matr = append(matr, newRow)

			}
			index++
			if index > 7 { // Hard coded size of tile :D
				break
			}
		}
		index = 0
	}

	// Try every version of full matrix (rotated/flipped) until result != 0
	for index := 0; index < 8; index++ {
		result := searchForSeaMonster(matr) // 0
		if result != 0 {
			return result
		}
		// Rotate for next round
		matr = rotateMatrix90ACW(matr)
		if index == 3 {
			// if 4 rotations done, flip the matrix and rotate again.
			matr = flipMatrix(matr)
		}
	}
	return 0
}

func findMonster(image string, linelen int) (count int) {
	count = 0
	isMonster := func(x int, y int) bool {
		for _, i := range []int{18} {
			if image[y*linelen+x+i] != '#' {
				return false
			}
		}
		for _, i := range []int{0, 5, 6, 11, 12, 17, 18, 19} {
			if image[(y+1)*linelen+x+i] != '#' {
				return false
			}
		}
		for _, i := range []int{1, 4, 7, 10, 13, 16} {
			if image[(y+2)*linelen+x+i] != '#' {
				return false
			}
		}
		return true
	}
	numLines := len(image) / linelen
	for y := 0; y < numLines-2; y++ {
		for x := 0; x <= linelen-20; x++ {
			if isMonster(x, y) {
				count++
			}
		}
	}
	return count
}

func searchForSeaMonster(matrix [][]string) int {
	oneString := ""
	lenOfString := 0
	for _, row := range matrix {
		tmp := ""
		for _, col := range row {
			tmp += col
		}
		lenOfString = len(tmp)
		oneString += tmp
	}
	monsters := findMonster(oneString, lenOfString)
	hashtags := strings.Count(oneString, "#")
	// monster is 15 #
	if monsters == 0 {
		return 0
	}
	result := hashtags - monsters*15
	return result
}

var tilesInUse = map[string]int{}

func programPartTwo() int {
	tileMatrix := make([][]tile, 30)
	for i := 0; i < len(tileMatrix); i++ {
		tileMatrix[i] = make([]tile, len(tileMatrix))
	}

	return createPic()
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 20: Jurassic Jigsaw ----")
	// testInput := lib.ReadFileAndPutInSlice(path + "\\20\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path + "\\20\\input.txt")
	//
	// result := program(testInput, 1)
	// println("Result Test One:", result)

	result := program(input, 1)
	println("Result Part One:", result)

	// result2 := programPartTwo()
	// println("Result Test Two:", result2)

	result2 := programPartTwo()
	println("Result Part Two:", result2)
}

//Result Part One: 17712468069479
//Result Part Two: 2173
