package main

import (
	"fmt"
)

func main() {
	count := 1000000

	for count > 0 {
		fmt.Println(count)
		count--
	}
	fmt.Println("Liftoff!")
}
