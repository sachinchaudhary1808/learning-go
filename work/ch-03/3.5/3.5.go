package main

import "fmt"

func main() {
	room := "lake"
	switch room {
	case "cave":
		fmt.Println("You find yourself in a dimly lit carvern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}
