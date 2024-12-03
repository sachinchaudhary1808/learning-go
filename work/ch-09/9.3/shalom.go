package main

import "fmt"

func main() {
	message := "shalom"

	for i := 0; i < 6; i++ {
		char := message[i]
		fmt.Printf("%c\n", char)
	}
}
