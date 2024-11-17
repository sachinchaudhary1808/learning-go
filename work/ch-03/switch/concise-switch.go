package main

import "fmt"

func main() {
	fmt.Println("There is a cavern entrance here aand path to east.")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You had further up to mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")

	}

}
