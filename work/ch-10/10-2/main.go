package main

import "fmt"

func main() {
	var bh float64 = 32768 // maximum length of int16
	var h = int16(bh)
	fmt.Println(h)
	// will work fine but if we just add 1 more to float64 the h will be recalculated to max + 1 so it will put the int16 to it's starting point which is -32768
}
