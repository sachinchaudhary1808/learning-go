package main

import "fmt"

func main() {
	age := 41
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("my age on mars is ", marsAge)

	// shit i'm still weak at math

	// result := 21.79

	// will write about this in book  | got it

}
