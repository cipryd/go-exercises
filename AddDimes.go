package main

import (
	"fmt"
	"math/rand"
)

func AddDimes() {
	dimes := [3]float64{0.05, 0.1, 0.25}
	piggyBank := 0.0
	for piggyBank < 20 {
		dime := dimes[rand.Intn(3)]
		piggyBank += dime
		fmt.Printf("dime %.2f   ", dime)
		fmt.Printf("%2.2f\n", piggyBank)
	}
	fmt.Println(piggyBank)
}
