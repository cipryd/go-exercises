package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Spaceline        Days Trip type  Price
// ======================================
// Virgin Galactic    23 Round-trip $  96
// Virgin Galactic    39 One-way    $  37
// SpaceX             31 One-way    $  41
// Space Adventures   22 Round-trip $ 100
// Space Adventures   22 One-way    $  50
// Virgin Galactic    30 Round-trip $  84
// Virgin Galactic    24 Round-trip $  94
// Space Adventures   27 One-way    $  44
// Space Adventures   28 Round-trip $  86
// SpaceX             41 Round-trip $  72

// Use October 13, 2020 as the departure date for all tickets.
// Mars will be 62,100,000 km away from Earth at the time.

// Randomly choose the speed the ship will travel, from 16 to 30 km/s.
// This will determine the duration for the trip to Mars and also the ticket price.
// Make faster ships more expensive, ranging in price from $36 million to $50 million.
// Double the price for round trips.

//Init pseudo-random generator
func Init() {
	rand.Seed(time.Now().Unix())
}

//GenerateTicketsToMars generates a ticket for a trip to Mars
//it containts the spaceline, nr of days, trip type and price
func GenerateTicketsToMars() {
	Init()
	const distanceToMars = 62100000 //km
	var tableHeader = "Spaceline        Days Speed Trip type    Price"
	rowDelimiter := ""
	for i := 0; i < len(tableHeader); i++ {
		rowDelimiter += "="
	}
	fmt.Printf("%v\n%v\n", tableHeader, rowDelimiter)

	var companies = []string{"Virgin Galactic", "SpaceX", "Space Adventures"}
	var tripType = []string{"Round-trip", "One-way"}

	for i := 0; i < 10; i++ {
		companyIndex := rand.Intn(len(companies))
		trip := rand.Intn(len(tripType))
		speed := rand.Intn(14) + 16
		days := distanceToMars / speed / 3600 / 24
		price := speed + 20

		switch tripType[trip] {
		case "Round-trip":
			price *= 2
		}

		fmt.Printf("%-16v  %2v   %v    %-10v  $ %3v\n", companies[companyIndex], days, speed, tripType[trip], price)
	}
}
