package main

import (
	"fmt"
)

func main() {
	var inputtedNumber int
	fmt.Scan(&inputtedNumber)
	repeatTimes := 6
	if inputtedNumber%2 == 0 {
		inputtedNumber += 1
	}
	for number := inputtedNumber; number < (inputtedNumber + (repeatTimes * 2)); number += 2 {
		fmt.Println(number)
	}
}
