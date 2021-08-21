package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tests(a, b, c, d int64) bool {
	if !(b > c && d > a) {
		return false
	}
	if !(c+d > a+b) {
		return false
	}
	if !(c > 0 && d > 0 && a%2 == 0) {
		return false
	}
	return true
}

func main() {
	values := bufio.NewReader(os.Stdin)
	line, _ := values.ReadString('\n')
	line = strings.Trim(line, "\n")
	valuesSplitted := strings.Split(line, " ")
	a, _ := strconv.ParseInt(valuesSplitted[0], 10, 64)
	b, _ := strconv.ParseInt(valuesSplitted[1], 10, 64)
	c, _ := strconv.ParseInt(valuesSplitted[2], 10, 64)
	d, _ := strconv.ParseInt(valuesSplitted[3], 10, 64)

	var message string
	if tests(a, b, c, d) {
		message = "Valores aceitos"
	} else {
		message = "Valores nao aceitos"
	}

	fmt.Println(message)
}
