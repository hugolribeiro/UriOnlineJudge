package main

import "fmt"

const amountNumbers = 6

func main() {
	var actualNumber float64
	var amountPositiveNumbers int
	for count := 0; count < amountNumbers; count++ {
		fmt.Scan(&actualNumber)
		if actualNumber >= 0 {
			amountPositiveNumbers += 1
		}
	}
	fmt.Printf("%d valores positivos\n", amountPositiveNumbers)
}
