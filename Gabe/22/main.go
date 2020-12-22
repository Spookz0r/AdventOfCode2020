package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gabe/lib"
)

var replacer = strings.NewReplacer("\n", "", "\r", "")

func getPlayers(input []string) ([]int, []int) {
	playerOne := []int{}
	playerTwo := []int{}
	player := 0
	for _, line := range input {
		line = replacer.Replace(line)
		if line == "Player 1:" {
			player = 1
			continue
		} else if line == "Player 2:" {
			player = 2
			continue
		} else if len(line) == 0 {
			continue
		}
		if player == 1 {
			nr, _ := strconv.Atoi(line)
			playerOne = append(playerOne, nr)
		} else if player == 2 {
			nr, _ := strconv.Atoi(line)
			playerTwo = append(playerTwo, nr)
		}
	}
	return playerOne, playerTwo
}
func calculateResult(playerOne []int, playerTwo []int, win int) int {
	winner := playerTwo
	if len(playerOne) > 0 || win == 0 {
		winner = playerOne
	}
	result := 0
	for index := len(winner) - 1; index >= 0; index-- {
		result += winner[index] * (len(winner) - index)
	}
	return result
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func updatePlayers(winner []int, loser []int) ([]int, []int) {
	a := winner[0]
	b := loser[0]
	winner = winner[1:]
	loser = loser[1:]
	winner = append(winner, a)
	winner = append(winner, b)
	return winner, loser
}

func recursiveCombat(playerOne []int, playerTwo []int) ([]int, []int, int) {
	playedRoundsOne := map[string]int{}
	playedRoundsTwo := map[string]int{}
	done := false
	winner := 0

	for done == false {
		if len(playerOne) == 0 || len(playerTwo) == 0 {
			break
		}
		// END GAME IF ROUND HAS BEEN PLAYED BEFORE
		if _, ok := playedRoundsOne[arrayToString(playerOne, "")]; ok {
			if _, ok := playedRoundsTwo[arrayToString(playerTwo, "")]; ok {
				playerOne, playerTwo = updatePlayers(playerOne, playerTwo)
				break
			}
		}
		// ADD PLAYED ROUNDS TO MAP
		playedRoundsOne[arrayToString(playerOne, "")] = 1
		playedRoundsTwo[arrayToString(playerTwo, "")] = 1

		// Check if card number is the same as the amount of cards in both decks - the card drawn
		if (playerOne[0] < len(playerOne)) && (playerTwo[0] < len(playerTwo)) {
			// Make a deep copy of original deck and use that
			p1 := make([]int, len(playerOne))
			p2 := make([]int, len(playerTwo))
			copy(p1, playerOne)
			copy(p2, playerTwo)
			_, _, win := recursiveCombat(p1[1:p1[0]+1], p2[1:p2[0]+1])
			if win == 1 {
				playerOne, playerTwo = updatePlayers(playerOne, playerTwo)
			} else {
				playerTwo, playerOne = updatePlayers(playerTwo, playerOne)
			}
		} else {
			// Compare first element in both lists
			if playerOne[0] > playerTwo[0] {
				playerOne, playerTwo = updatePlayers(playerOne, playerTwo)
			} else {
				playerTwo, playerOne = updatePlayers(playerTwo, playerOne)
			}
		}
	}

	if len(playerOne) > 0 {
		winner = 1
	} else {
		winner = 2
	}
	return playerOne, playerTwo, winner

}

func programPartTwo(input []string) {
	playerOne, playerTwo := getPlayers(input)
	playerOne, playerTwo, winner := recursiveCombat(playerOne, playerTwo)
	result := calculateResult(playerOne, playerTwo, winner)
	fmt.Println("Result Part Two", result)
}

func program(input []string) {
	playerOne, playerTwo := getPlayers(input)
	done := false
	for done == false {
		if len(playerOne) == 0 || len(playerTwo) == 0 {
			break
		}
		// Compare first element in both lists
		if playerOne[0] > playerTwo[0] {
			playerOne, playerTwo = updatePlayers(playerOne, playerTwo)
		} else {
			playerTwo, playerOne = updatePlayers(playerTwo, playerOne)
		}
	}
	result := calculateResult(playerOne, playerTwo, 6)
	fmt.Println("Result Part One", result)
}

func main() {
	path, err := os.Getwd()
	lib.Check(err)

	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 22: Crab Combat ----")
	// testInput := lib.ReadFileAndPutInSlice(path + "\\22\\test_input.txt")
	input := lib.ReadFileAndPutInSlice(path + "\\22\\input.txt")
	// program(testInput)
	// programPartTwo(testInput)
	program(input)
	programPartTwo(input)
}
