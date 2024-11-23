package main

import "fmt"

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysInYear = 365

	const distanceInLightYear = distance / lightSpeed / (secondsPerDay * daysInYear)

	fmt.Println(distanceInLightYear)

}
