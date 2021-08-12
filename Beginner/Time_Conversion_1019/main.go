package main

import (
	"fmt"
)

func main() {
	var durationSeconds int
	fmt.Scan(&durationSeconds)

	hours := durationSeconds / 3600
	durationSeconds = durationSeconds % 3600
	minutes := durationSeconds / 60
	durationSeconds = durationSeconds % 60

	fmt.Printf("%d:%d:%d\n", hours, minutes, durationSeconds)
}
