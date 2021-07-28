package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	soma := a + b
	fmt.Printf("SOMA = %d\n", soma)
}
