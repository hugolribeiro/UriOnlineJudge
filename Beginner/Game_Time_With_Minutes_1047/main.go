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
	fmt.Println(valuesSplitted[0])
	initialHour, _ := strconv.ParseInt(valuesSplitted[0], 10, 64)
	initialMinutes, _ := strconv.ParseInt(valuesSplitted[1], 10, 64)
	endHour, _ := strconv.ParseInt(valuesSplitted[2], 10, 64)
	endMinutes, _ := strconv.ParseInt(valuesSplitted[3], 10, 64)

	initialDuration := initialHour*60 + initialMinutes
	endDuration := endHour*60 + endMinutes

	totalMinutes := endDuration - initialDuration
	if totalMinutes <= 0 {
		totalMinutes = totalMinutes + 24*60
	}

	totalHours := totalMinutes / 60
	totalMinutes = totalMinutes % 60
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)", totalHours, totalMinutes)
}
