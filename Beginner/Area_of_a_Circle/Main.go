package main

import (
    "fmt"
)

const PI float64 = 3.14159

func main() {

    /**
     * Escreva a sua solução aqui
     * Code your solution here
     * Escriba su solución aquí
     */
 	var radius float64
	fmt.Scan(&radius)
	area := radius * radius * PI
	fmt.Printf("A=%.4f\n", area)

}
