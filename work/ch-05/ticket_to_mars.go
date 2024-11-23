// package main is for just learning
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const (
		distance     = 62100000 // km/s
		depatureDate = "13/10/2020"
		minSpeed     = 16 // km/s
		maxSpeed     = 30 // km/s
		minPrice     = 36 // million
		maxPrice     = 50 // million
	)

	fmt.Printf("%-16v %4v %-10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")
	for num := 10; num > 0; num-- {

		Spaceline := ""

		switch rand.Intn(3) {
		case 0:
			Spaceline = "Space Adventures"
		case 1:
			Spaceline = "SpaceX"
		case 2:
			Spaceline = "Virgin Galactic"

		}

		speed := rand.Intn(15) + 16
		price := 20.0 + speed

		tripType := ""
		if rand.Intn(2) == 0 {
			tripType = "One-way"
		} else {
			tripType = "Round-trip"
			price = (20 + speed) * 2

		}

		hoursInday := 24
		days := distance / (speed * 60 * 60 * hoursInday)

		fmt.Printf("%-16v %3v %-11v $%4v\n", Spaceline, days, tripType, price)
	}
}
