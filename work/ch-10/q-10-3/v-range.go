package main

import "math"
import "fmt"

func main() {
	// v := 12321323
	//
	// if v < 0 || v > math.MaxUint8 {
	// 	println("v is out uint8 range")
	// } else {
	// 	println("v is in the range of unit8")
	// }

	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("converted:", v8)
	}

	// both are fine right ? lol  i think so

	// in second approch they converted the the v to uint8 which is nice
}
