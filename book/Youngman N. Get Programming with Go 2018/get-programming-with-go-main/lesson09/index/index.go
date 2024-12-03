package main

import "fmt"

func main() {
	message := "shalom"
	c := message[5]
	fmt.Printf("%c %[1]v\n", c)

	message[0] = 'd'

}
