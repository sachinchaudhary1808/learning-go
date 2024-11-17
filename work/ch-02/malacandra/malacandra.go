package main

import "fmt"

func main() {
	const hours_Per_day = 24
	var distance = 56000000
	var days = 28
	fmt.Println(distance/days/hours_Per_day, "km/h")
}
