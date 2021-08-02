package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total float64
	for i := 0; i < 2; i++ {
		values := bufio.NewReader(os.Stdin)
		line, err := values.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err == nil {
			valuesSplitted := strings.Split(line, " ")
			amountUnits, err := strconv.ParseInt(valuesSplitted[1], 10, 64)
			priceUnit, erro := strconv.ParseFloat(valuesSplitted[2], 64)
			if err == nil && erro == nil {
				averageTotal := priceUnit * float64(amountUnits)
				total += averageTotal
			}
		}
	}
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}