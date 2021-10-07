package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Scan(&number)
	for i := 2; i <= number; i += 2 {
		fmt.Printf("%d^2 = %.0f\n", i, math.Pow(float64(i), 2))
	}
}
