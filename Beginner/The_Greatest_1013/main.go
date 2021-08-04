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
	var integerNumbers []int
	for i := 0; i < len(valuesSplitted); i++ {
		actualNumber, err := strconv.ParseInt(valuesSplitted[i], 10, 64)
		if err != nil {

		}
		integerNumbers = append(integerNumbers, int(actualNumber))
	}
	highestNumber := -9999
	for i := 0; i < len(integerNumbers); i++ {
		if integerNumbers[i] > highestNumber {
			highestNumber = integerNumbers[i]
		}
	}

	fmt.Printf("%d eh o maior\n", highestNumber)
}
