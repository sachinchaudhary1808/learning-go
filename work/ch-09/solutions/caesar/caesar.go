package main

import "fmt"

func main() {
	var message = "L fdph, L vdz, L frqtxhuhg"

	for i := 0; i < len(message); i++ {

		char := message[i]
		if char >= 'a' && char <= 'z' {
			char = char - 3
			if char < 'a' {
				char = char + 26
			}
		}

		if char >= 'A' && char <= 'Z' {
			char = char - 3
			if char < 'A' {
				char = char + 26
			}
		}
		fmt.Printf("%c", char)
	}

}
