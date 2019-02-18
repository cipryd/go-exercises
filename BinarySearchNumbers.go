package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNrBetween(low int, high int) int {
	return rand.Intn(high-low) + low + 1
}

func guessNumberBetweenNumbers(number int, low int, high int) {
	var tries = 0
	rand.Seed(time.Now().UnixNano())
	for {
		tries++
		fmt.Printf("\nTry %1v - ", tries)
		fmt.Printf("Guess between %2v and %3v: ", low, high)
		var guess = generateRandomNrBetween(low, high)

		if guess == number {
			fmt.Println("Guess correct. The number was 17")
			break
		} else if guess > number {
			high = guess
			fmt.Printf("Your guess %3v is too big", guess)
		} else {
			low = guess
			fmt.Printf("Your guess %3v is too small", guess)
		}
	}
}
