package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c int
	fmt.Scan(&c)
	var d int
	fmt.Scan(&d)

	diferenca := (a*b - c*d)
	fmt.Printf("DIFERENCA = %d", diferenca)
}
