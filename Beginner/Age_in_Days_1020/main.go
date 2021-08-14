package main

import (
	"fmt"
)

func main() {
	var personAgeInDays int
	fmt.Scan(&personAgeInDays)
	years := int(personAgeInDays / 365)
	personAgeInDays = personAgeInDays % 365
	months := int(personAgeInDays / 30)
	personAgeInDays = personAgeInDays % 30

	fmt.Print(years, " ano(s)\n")
	fmt.Print(months, " mes(es)\n")
	fmt.Print(personAgeInDays, " dia(s)\n")
}
