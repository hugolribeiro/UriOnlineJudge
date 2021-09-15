package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	values := bufio.NewReader(os.Stdin)
	line, _ := values.ReadString('\n')
	line = strings.Trim(line, "\n")
	valuesSplitted := strings.Split(line, " ")
	a, _ := strconv.ParseInt(valuesSplitted[0], 10, 64)
	b, _ := strconv.ParseInt(valuesSplitted[1], 10, 64)
	if a%b == 0 || b%a == 0 {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}
