// does nothing just for learning prupose
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggybank = 00

	for piggybank < 2000 {
		switch rand.Intn(3) {
		case 0:
			piggybank = piggybank + 5
		case 1:
			piggybank = piggybank + 10
		case 2:
			piggybank = piggybank + 25
		}

		dollers := piggybank / 100
		cents := piggybank % 100

		fmt.Printf("$%v.%02v\n", dollers, cents)

	}

}
