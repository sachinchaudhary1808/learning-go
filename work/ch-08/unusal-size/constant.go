package main

import "fmt"

func main() {
	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	fmt.Println("Andromeda Galaxy is", 24000000000000000000/299792/86400, "light days away.")
	fmt.Println("Andromeda Galaxy is", distance/lightSpeed/secondsPerDay, "light days away.")

}
