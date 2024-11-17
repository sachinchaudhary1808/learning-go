package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const (
		distance      = 62100000 // km/s
		depature_date = "13/10/2020"
		min_speed     = 16 // km/s
		max_speed     = 30 // km/s
		min_price     = 36 // million
		max_price     = 50 // million
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

		trip_type := ""
		if rand.Intn(2) == 0 {
			trip_type = "One-way"
		} else {
			trip_type = "Round-trip"
			price = (20 + speed) * 2

		}

		hours_inday := 24
		days := distance / (speed * 60 * 60 * hours_inday)

		fmt.Printf("%-16v %3v %-11v $%4v\n", Spaceline, days, trip_type, price)
	}
}
