package main

import "fmt"

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	universe := make(Universe, height)

	for i := range universe {
		universe[i] = make([]bool, width)
	}

	fmt.Printf("%v %v", len(universe), cap(universe))
	return universe
}

func main() {
	NewUniverse()
}
