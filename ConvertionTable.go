package main

import (
	"fmt"
)

func drawTable(samples int, temp func(int) (t1 float64, t2 float64)) {
	rowLength := 25
	firstColumn := "°C"
	secondColumn := "°F"
	drawRow(rowLength)
	drawHeader(firstColumn, secondColumn, rowLength)
	drawRow(rowLength)
	for i := 0; i < samples; i++ {
		value1, value2 := temp(i)
		drawValues(value1, value2)
		drawRow(rowLength)
	}
}

func getValues(nr int) (t1 float64, t2 float64) {
	temp1 := celsius(-40.0 + 5.0*nr)
	temp2 := temp1.fahrenheit()
	return float64(temp1), float64(temp2)
}

func drawRow(nr int) {
	for i := 0; i < nr; i++ {
		fmt.Printf("=")
	}
	fmt.Println()
}

func drawHeader(s1 string, s2 string, nr int) {
	fmt.Printf("| %5v     ", s1)
	fmt.Printf("| %5v     ", s2)
	fmt.Printf("|\n")
}

func drawValues(s1 float64, s2 float64) {
	fmt.Printf("| %7.2f   ", s1)
	fmt.Printf("| %7.1f   |\n", s2)
}
