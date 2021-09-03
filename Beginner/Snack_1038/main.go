package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var snack = map[string]float64{
		"1": 4.0,
		"2": 4.5,
		"3": 5.0,
		"4": 2.0,
		"5": 1.5,
	}

	values := bufio.NewReader(os.Stdin)
	line, _ := values.ReadString('\n')
	line = strings.Trim(line, "\n")
	valuesSplitted := strings.Split(line, " ")
	code := valuesSplitted[0]
	amount, _ := strconv.ParseFloat(valuesSplitted[1], 64)

	total := snack[code] * amount
	fmt.Printf("Total: R$ %.2f\n", total)
}
