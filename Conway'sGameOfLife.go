package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 80
	height = 15
)

// Universe is a type which holds a 2d field of cells.
// Each cell will be either dead(false) or alive(true)
type Universe [][]bool

// NewUniverse creates a Universe with heigth rows and width columns per row
func NewUniverse() Universe {
	universe := make(Universe, height)

	for i := range universe {
		universe[i] = make([]bool, width)
	}

	return universe
}

// Show prints the universe to the screen
func (u Universe) Show() {
	for i := range u {
		for j := range u[i] {
			if u[i][j] {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

// Seed sets randomly ~25% of the cells to alive
func (u Universe) Seed() {
	for i := range u {
		for j := range u[i] {
			if rand.Intn(100) < 25 {
				u[i][j] = true
			}
		}
	}
}

// Alive determines wheter a cell is dead or alive
func (u Universe) Alive(x, y int) bool {
	for x < 0 {
		x += height
	}
	for y < 0 {
		y += width
	}

	if x >= height {
		x = x % height
	}
	if y >= width {
		y = y % width
	}

	return u[x][y]
}

// Neighbors determines the number of alive neighbors for a given cell from 0 to 8
func (u Universe) Neighbors(x, y int) int {
	aliveNeighbors := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(i == x && j == y) {
				if u.Alive(i, j) {
					aliveNeighbors++
				}
			}
		}
	}
	return aliveNeighbors
}

// Next determines whether a cell should be dead or alive in the next generation
func (u Universe) Next(x, y int) bool {
	aliveCell := u.Alive(x, y)
	liveNeighbors := u.Neighbors(x, y)

	if aliveCell && liveNeighbors < 2 {
		return false
	}
	if aliveCell && (liveNeighbors == 2 || liveNeighbors == 3) {
		return true
	}
	if aliveCell && liveNeighbors > 3 {
		return false
	}

	if !aliveCell && liveNeighbors == 3 {
		return true
	}

	return false
}

// Step sets the next generation of Universe a to Universe b
func Step(a, b Universe) {
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.Next(i, j)
		}
	}
}

// ClearScreen clear the screen on Windows
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// RunSimulation runs Conway's Game Life for n generations
func RunSimulation(n int) {
	a := NewUniverse()
	b := NewUniverse()
	a.Seed()

	for i := 0; i < n; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(100 * time.Millisecond)
		ClearScreen()
		a, b = b, a
	}
}
