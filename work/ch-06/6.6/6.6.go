// does nothing just for learning purpose
package main

import (
	"fmt"
	"math"
)

func main() {
	var pigbank float64
	for num := 0; num < 11; num++ {
		pigbank += 0.10

	}
	fmt.Println(pigbank)

	fmt.Println(math.Abs(pigbank-1.10) < 0.0001)
}
