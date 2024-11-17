package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Set the target number
	num := 47

	// Start a loop to keep guessing
	for {
		// Generate a random number between 1 and 100
		guess := rand.Intn(100) + 1

		// Print the guess
		fmt.Println("Computer guesses:", guess)

		// Check if the guess is correct
		if guess == num {
			// If the guess is correct, print "You win!" and break out of the loop
			fmt.Println("Computer guessed correctly! The number was:", num)
			break
		} else if guess > num {
			// If the guess is too high
			fmt.Println("Too high!")
		} else {
			// If the guess is too low
			fmt.Println("Too low!")
		}
	}
}
