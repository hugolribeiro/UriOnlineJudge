package main

import (
	"fmt"
)

func main() {
	var distanceTraveled int
	fmt.Scan(&distanceTraveled)
	var spentFuel float64
	fmt.Scan(&spentFuel)

	averageconsumption := float64(distanceTraveled) / (spentFuel)

	fmt.Printf("%.3f km/l\n", averageconsumption)
}
