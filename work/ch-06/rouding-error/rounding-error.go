package main

import (
	"fmt"
	"math"
)

func main() {
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "º F\n")
	fmt.Print((9.0/5.0*celsius)+32, "º F\n")

	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Println(fahrenheit, "º F")
	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank == 0.3)

	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)
}
