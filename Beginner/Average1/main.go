package main

import (
	"fmt"
)

func main() {

	/**
	 * Escreva a sua solução aqui
	 * Code your solution here
	 * Escriba su solución aquí
	 */
	var a float32
	fmt.Scan(&a)
	var b float32
	fmt.Scan(&b)
	average := ((a * 3.5) + (b * 7.5)) / 11
	fmt.Printf("MEDIA = %.5f\n", average)
}
