package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)


func programPartOne(input []string) int{
	validPassports := 0
	var hasValidAttribute = regexp.MustCompile(`^(byr|iyr|eyr|hgt|hcl|ecl|pid)$`).MatchString

	for _, passport := range input{
		data := strings.Split(passport, " ")
		validCounter := 0
		for _, tmp := range data{
			attr := strings.Split(tmp,":")
			if hasValidAttribute(attr[0]){
				validCounter++
			}
		}
		if validCounter == 7{
			validPassports++
		}
	}
	return validPassports
}


func programPartTwo(input []string) int{
	validPassports := 0
	var isHEXCode =   regexp.MustCompile(`^[#][a-f0-9]{6}$`).MatchString
	var isHairColor = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`).MatchString

	for _, passport := range input{
		data := strings.Split(passport, " ")
		validCounter := 0
		for _, tmp := range data{
			attr := strings.Split(tmp,":")
			switch attr[0]{
			case "byr":
				value , _:= strconv.Atoi(attr[1])
				if value >= 1920 && value <= 2002{
					validCounter++
				}
			case "iyr":
				value , _:= strconv.Atoi(attr[1])
				if value >= 2010 && value <= 2020{
					validCounter++
				}
			case "eyr":
				value , _:= strconv.Atoi(attr[1])
				if value >= 2020 && value <= 2030{
					validCounter++
				}
			case "hgt":
				if strings.Contains(attr[1], "cm"){
					height := strings.Replace(attr[1],"cm","",-1)
					value , _:= strconv.Atoi(height)
					if value >= 150 && value <= 193{
						validCounter++
					}
				}
				if strings.Contains(attr[1], "in"){
					height := strings.Replace(attr[1],"in","",-1)
					value , _:= strconv.Atoi(height)
					if value >= 59 && value <= 76{
						validCounter++
					}
				}
			case "hcl":
				if isHEXCode(attr[1]){
					validCounter++
				}
			case "ecl":
				if isHairColor(attr[1]) {
					validCounter++
				}
			case "pid":
				if len(attr[1]) == 9{
					validCounter++
				}
			}
		}

		if validCounter == 7 {
			validPassports++
		}
	}
	return validPassports
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 4: Passport Processing ----")
	input := lib.ReadFileAndSplitAtEmptyLineAndPutInSlice(path+"\\04\\input.txt")

	result := programPartOne(input)
	println("Result Part One:",result)

	result = programPartTwo(input)
	println("Result Part Two:",result)

}
