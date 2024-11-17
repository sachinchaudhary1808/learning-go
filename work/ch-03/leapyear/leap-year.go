package main

import "fmt"

func main() {
	var year = 2024
	fmt.Printf("The year is %v, should you leap?\n", year)

	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("it's a leap year.")

	} else {
		fmt.Println("it's not a leap year.")
	}
}
