// // How long does it take to get to Mars?
// package main
//
// import "fmt"
//
//	func main() {
//		const lightSpeed = 299792 // Km/s
//		var distance = 56000000   // Km
//
//		fmt.Println(distance/lightSpeed, "seconds")
//
//		distance = 401000000
//		fmt.Println(distance/lightSpeed, "seconds")
//
// }
package main

import "fmt"

func main() {
	const (
		hours_in_day = 24     // hour
		speed        = 100800 // km/h
	)
	var distance = 96300000 // km
	fmt.Printf("it will take %v day to reach mars with this speed", distance/speed/hours_in_day)

}
