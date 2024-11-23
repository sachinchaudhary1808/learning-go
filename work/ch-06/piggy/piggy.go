package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggybank = 0.0
	for piggybank < 20.00 {
		switch rand.Intn(3) {
		case 0:
			piggybank += 0.05
		case 1:
			piggybank += 0.10
		case 2:
			piggybank += 0.25

		}
		fmt.Printf("$%5.2f\n", piggybank)

	}
}
