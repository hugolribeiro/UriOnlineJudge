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
	a, _ := strconv.ParseFloat(valuesSplitted[0], 64)
	b, _ := strconv.ParseFloat(valuesSplitted[1], 64)
	c, _ := strconv.ParseFloat(valuesSplitted[2], 64)

	triangleArea := a * c / 2
	circleArea := 3.14159 * math.Pow(c, 2)
	trapeziumArea := (a + b) * c / 2
	squareArea := b * b
	rectangleArea := a * b

	fmt.Printf("TRIANGULO: %.3f\n", triangleArea)
	fmt.Printf("CIRCULO: %.3f\n", circleArea)
	fmt.Printf("TRAPEZIO: %.3f\n", trapeziumArea)
	fmt.Printf("QUADRADO: %.3f\n", squareArea)
	fmt.Printf("RETANGULO: %.3f\n", rectangleArea)

}
