package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	values := bufio.NewReader(os.Stdin)
	line, _ := values.ReadString('\n')
	line = strings.Trim(line, "\n")
	valuesSplitted := strings.Split(line, " ")
	var floatNumbers []float64
	for i := 0; i < len(valuesSplitted); i++ {
		actualNumber, err := strconv.ParseFloat(valuesSplitted[i], 64)
		if err != nil {
		}
		floatNumbers = append(floatNumbers, actualNumber)
	}
	a, b, c := floatNumbers[0], floatNumbers[1], floatNumbers[2]

	delta := math.Pow(b, 2) - (4 * a * c)

	if a == 0 || delta < 0 {
		fmt.Println("Impossivel calcular")
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("R1 = %.5f\n", x1)
		fmt.Printf("R2 = %.5f\n", x2)
	}
}
