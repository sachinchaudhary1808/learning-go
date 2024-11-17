package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for count := 10; count > 0; count-- {
		year := rand.Intn(2024-1900+1) + 1900
		month := rand.Intn(12) + 1
		day_in_month := 31

		switch month {
		case 2:
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				day_in_month = 29
			} else {
				day_in_month = 28
			}
		case 4, 6, 9, 11:
			day_in_month = 30
		}

		day := rand.Intn(day_in_month) + 1
		fmt.Println(era, year, month, day)

	}
}
