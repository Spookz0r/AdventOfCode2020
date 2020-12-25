package main

import (
	"fmt"
)

func getSubjectNumberAndLoopSize(cardVal int, subjectNumber int) int {
	loopSize := 0
	a := 1
	for a != cardVal {
		loopSize++
		a = (a * subjectNumber) % 20201227
	}
	return loopSize
}

func program(card int, door int) {
	// Use the public keys to determine cards public keys loop size
	loopCard := getSubjectNumberAndLoopSize(card, 7)
	// Now use the door public key as subject number with loop size of card to get encryption key
	encryptionKey := 1
	for i := 0; i < loopCard; i++ {
		encryptionKey = (encryptionKey * door) % 20201227
	}
	fmt.Println("Encryption key is:", encryptionKey)
}

func main() {
	fmt.Println("---- Advent Of Code 2020 ----")
	fmt.Println("---- Day 25: Combo Breaker ----")
	// program(5764801, 17807724)
	program(8335663, 8614349)
}
