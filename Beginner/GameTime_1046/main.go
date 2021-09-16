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

	initialHour, _ := strconv.ParseInt(valuesSplitted[0], 64, 10)
	endHour, _ := strconv.ParseInt(valuesSplitted[1], 64, 10)

	duration := 24

	if endHour > initialHour {
		duration = int(endHour) - int(initialHour)
	} else if endHour < initialHour {
		duration = 24 + int(endHour) - int(initialHour)
	}
	fmt.Printf("O JOGO DUROU %d HORA(S)\n", duration)

}
