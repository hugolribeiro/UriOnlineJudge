package main

import "fmt"

func main() {
	var highest int
	numberIndex := 1
	fmt.Scan(&highest)
	var actualNumber int
	for index := 2; index <= 100; index++ {
		fmt.Scan(&actualNumber)
		if actualNumber > highest {
			highest = actualNumber
			numberIndex = index
		}
	}
	fmt.Println(highest)
	fmt.Println(numberIndex)
}
