// last thing of lesson 10
package main

import "fmt"

func main() {

	string := "8"
	var truefalse bool
	switch string {

	case "true", "yes", "1":
		truefalse = true
	case "false", "no", "0":
		truefalse = false

	default:
		fmt.Println("not valid")

	}

	println(truefalse)

	// works lol

}
