package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown, err := strconv.Atoi("10")
	if err != nil {
		println("something went wrong")
	}
	fmt.Println(countdown)
}
