package main

import "fmt"

func main() {
	var pi = 960
	var alpha = 940
	var omega = 969
	var bang byte = 33

	grade := '*'

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)
	fmt.Println(grade)
	fmt.Printf("%c ", grade)
}
