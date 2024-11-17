package main

import "fmt"

func main() {
	var haveTorch = true
	var liTorch = false
	if haveTorch || !liTorch {
		fmt.Println("Nothing to see here.")
	}
}
