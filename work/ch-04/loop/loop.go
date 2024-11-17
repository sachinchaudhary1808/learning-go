package main

import (
	"fmt"
)

func main() {
	count := 5
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println("count is:", count)
}
