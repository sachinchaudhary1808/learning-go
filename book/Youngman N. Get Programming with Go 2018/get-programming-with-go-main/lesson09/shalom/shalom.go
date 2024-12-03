package main

import (
	"fmt"
)

func main() {
	message := "shalom"
	for i := 0; i < len(message); i++ {
		fmt.Printf("%c  %[1]v\n", message[i])

	}
}
