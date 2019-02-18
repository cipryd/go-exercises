package main

import (
	"math/rand"
)

// GenerateRandomDate returns a random year, month and a day till 2018
func GenerateRandomDate() (int, int, int) {
	year := rand.Intn(2018) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		if year%400 == 0 {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	return year, month, day
}
