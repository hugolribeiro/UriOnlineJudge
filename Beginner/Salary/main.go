package main

import (
	"fmt"
)

func main() {
	var employeeNumber int
	fmt.Scan(&employeeNumber)
	var workedHours int
	fmt.Scan(&workedHours)
	var amountPerHour float32
	fmt.Scan(&amountPerHour)

	var salary float32 = amountPerHour * float32(workedHours)

	fmt.Printf("NUMBER = %d\n", employeeNumber)
	fmt.Printf("SALARY = U$ %.2f\n", salary)
}
