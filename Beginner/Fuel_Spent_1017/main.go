package main

import (
	"fmt"
)

const fuelSpent = 12

func main() {
	var spentTime float64
	fmt.Scan(&spentTime)
	var averageSpeed float64
	fmt.Scan(&averageSpeed)
	distance := averageSpeed * spentTime / fuelSpent
	fmt.Printf("%.3f\n", distance)
}
