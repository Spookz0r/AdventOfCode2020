package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)
var replacer = strings.NewReplacer("\n", "", "\r", "", " ", "")

// North = 0
// East = 90
// South = 180
// West = 270
func cosDeg(angleDegrees int) int{
	angleInRadians := (float64(angleDegrees) * math.Pi)/180
	return int(math.Cos(angleInRadians))
}
func sinDeg(angleDegrees int) int{
	angleInRadians := (float64(angleDegrees) * math.Pi)/180
	return int(math.Sin(angleInRadians))
}

func programPartTwo(input []string)int{
	x := 0
	y := 0
	waypointX := 10
	waypointY := 1
	for _, i := range input{
		txt := replacer.Replace(i)
		command := string(txt[0])
		value, _ := strconv.Atoi(string(txt[1:]))
		switch command {
		case "N":
			waypointY += value
			break
		case "S":
			waypointY -= value
			break
		case "E":
			waypointX += value
			break
		case "W":
			waypointX -= value
			break
		case "L":
			x1 := waypointX
			y1 := waypointY
			waypointX = x1 * cosDeg(value) - y1 * sinDeg(value)
			waypointY = x1 * sinDeg(value) + y1 * cosDeg(value)
			break
		case "R":
			x1 := waypointX
			y1 := waypointY
			waypointX = x1 * cosDeg(-value) - y1 * sinDeg(-value)
			waypointY = x1 * sinDeg(-value) + y1 * cosDeg(-value)
			break
		case "F":
			x += waypointX * value
			y += waypointY * value
		}
	}
	fmt.Println("End position: ", x, y)

	xAbs := int(math.Abs(float64(x)))
	yAbs := int(math.Abs(float64(y)))
	
	return xAbs + yAbs
}

func programPartOne(input []string)int{
	rotation := 90
	x := 0
	y := 0
	
	for _, i := range input{
		txt := replacer.Replace(i)
		command := string(txt[0])
		value, _ := strconv.Atoi(string(txt[1:]))
		switch command {
		case "N":
			y += value
			break
		case "S":
			y -= value
			break
		case "E":
			x += value
			break
		case "W":
			x -= value
			break
		case "L":
			rotation -= value
			break
		case "R":
			rotation += value
			break
		case "F":
			switch rotation{
			case 0:
				y += value
				break
			case 90:
				x += value
				break
			case 180:
				y -= value
				break
			case 270:
				x -= value
				break
			}
			break
		}
		if rotation < 0{
			rotation = 360 + rotation
		} else if rotation >= 360{
			rotation = rotation - 360
		}
	}
	xAbs := int(math.Abs(float64(x)))
	yAbs := int(math.Abs(float64(y)))
	
	return xAbs + yAbs
}


func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 12: Rain Risk ----")
	testInput := lib.ReadFileAndPutInSlice(path+"\\12\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path+"\\12\\input.txt")

	result := programPartOne(testInput)
	println("Result Test One:",result)

	result = programPartOne(input)
	println("Result Part One:",result)

	result = programPartTwo(testInput)
	println("Result Test One Part Two:",result)

	result = programPartTwo(input)
	println("Result Part Two:",result)
}
