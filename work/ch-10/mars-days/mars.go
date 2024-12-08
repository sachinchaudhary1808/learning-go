package main

import "fmt"

func main() {
	age := 18.0
	marsDays := 687.0     // days = 1 year on mars
	earthDays := 365.2425 // days = 1 year in earth

	fmt.Println("I am", age*earthDays/marsDays, "years old on Mars,")

	// to fix this we need same types of variables but that could provide irrational result

	// result :=  I am 9.569672489082969 years old on Mars,

}

