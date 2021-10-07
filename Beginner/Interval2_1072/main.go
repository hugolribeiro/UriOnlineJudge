package main

import "fmt"

func main() {
	var amountNumbers int
	fmt.Scan(&amountNumbers)
	var actualNumber int
	var amountIn int
	var amountOut int

	for i := 0; i < amountNumbers; i++ {
		fmt.Scan(&actualNumber)
		if actualNumber < 10 || actualNumber > 20 {
			amountOut += 1
			continue
		}
		amountIn += 1
	}
	fmt.Printf("%d in\n", amountIn)
	fmt.Printf("%d out\n", amountOut)
}
