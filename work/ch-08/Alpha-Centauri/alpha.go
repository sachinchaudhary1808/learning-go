// does nothing just for learning purpose
package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const SecondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / SecondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
}
