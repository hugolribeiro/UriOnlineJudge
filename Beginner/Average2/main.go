package main

import (
	"fmt"
)

func main() {
	var a float32
	fmt.Scan(&a)
	var b float32
	fmt.Scan(&b)
	var c float32
	fmt.Scan(&c)

	average := ((a * 2) + (b * 3) + (c * 5)) / 10
	fmt.Printf("MEDIA = %.1f\n", average)
}
