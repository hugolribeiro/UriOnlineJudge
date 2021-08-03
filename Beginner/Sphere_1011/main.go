package main

import (
	"fmt"
	"math"
)

const PI = 3.14159

func main() {
	var radius float64
	fmt.Scan(&radius)
	volume := (4.0 / 3) * PI * math.Pow(radius, 3)
	fmt.Printf("VOLUME = %.3f\n", volume)
}
