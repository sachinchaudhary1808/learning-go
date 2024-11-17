package main

import "fmt"

func main() {
	var command = "mountains"
	if command == "cave" {
		fmt.Println("u choose cave, welcome")
	} else if command == "entrance" {
		fmt.Println("you choose the entrance let's go")
	} else if command == "mountains" {
		fmt.Println("u choose the mountains let's go")
	} else {
		fmt.Println("don't know what that means")
	}
}
