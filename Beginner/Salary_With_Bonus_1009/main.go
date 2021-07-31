package main

import (
	"fmt"
)

func main() {
	var sellerName string
	fmt.Scan(&sellerName)
	var fixedSalary float64
	fmt.Scan(&fixedSalary)
	var totalSale float64
	fmt.Scan(&totalSale)

	total := fixedSalary + (0.15 * totalSale)

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
