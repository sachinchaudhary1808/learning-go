package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown, err := strconv.Atoi("255")
	if err != nil {
		println("something went wrong")
	} else {
		fmt.Println(countdown)
	}
}
