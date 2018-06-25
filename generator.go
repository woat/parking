package main

import (
	"math/rand"
	"sort"
)

const (
	minTime = 630
	maxTime = 2100
)

type Ticket struct {
	timeIn  int
	timeOut int
	cost    int
}

// Signature for a func to determine whether a time is within a constraint.
type weighter func(int) bool

// Generate time weight probabilities by utilizing a slice and a random idx.
// [1, 1, 1, 2] -> 3/4 chance of choosing 1.
func timeGenerator(w weighter) []int {
	var weightedSlice []int
	i := minTime
	for ; i <= maxTime; i++ {
		// No such time exists between 60 -> 99.
		if i%100 > 59 {
			i += 40
		}
		// The default weight is ~2/n.
		weight := rand.Intn(2)
		// More traffic exists during certain time constraints.
		// If the time exists between such a constraint
		// give it's weight a probability of (~2+~5)/n
		if w(i) {
			weight += rand.Intn(5)
		}
		for ; weight >= 0; weight-- {
			weightedSlice = append(weightedSlice, i)
		}
	}
	return weightedSlice
}

func timeDifference(timeIn, timeOut int) int {
	delta := timeOut - timeIn
	// The raw delta may produce an unreadable number.
	// In the case of 1800 - 1630, the result is 170.
	if delta%100 > 60 {
		// 170 -> 70 > 60 -> true
		delta -= 40
		// 170 -> 130
	}
	return delta
}

func totalCost(timeIn, timeOut int) int {
	timeD := timeDifference(timeIn, timeOut)
	// I had to think about this over a lunch at waffle house
	// to figure this out.
	// You can view my full explanation here: https://pastebin.com/8VT98V1H
	rounder := (100 - (timeD % 100)) % 100
	timeD += rounder
	return int((float64(timeD) * 0.01) * 2)
}

func totalTicketsGenerator() []Ticket {
	var totalTickets []Ticket
	timesInWeight := timeGenerator(func(i int) bool { return i >= 630 && i <= 900 })
	timesOutWeight := timeGenerator(func(i int) bool { return i >= 445 && i <= 630 })
	// ~450 parking spots exist in the parking garage.
	// todaysTickets := rand.Intn(200) + 200
	todaysTickets := rand.Intn(200) + 200
	// Not only do I assume that the garage never reaches max capacity
	// I also assume that the combined total of tickets produced does not
	// exceed max capacity as well.
	for i := 0; i <= todaysTickets; i++ {
		timeIn := timesInWeight[rand.Intn(len(timesInWeight))]
		timeOut := timesOutWeight[rand.Intn(len(timesOutWeight))]
		// Make sure time out is less than time in or
		// the total time spent parking is at least 1 hour.
		for timeOut < timeIn || timeDifference(timeIn, timeOut) < 100 {
			timeOut = timesOutWeight[rand.Intn(len(timesOutWeight))]
			// Anyone who comes after 8PM will always leave at 9PM.
			if timeIn >= 2000 {
				timeOut = 2100
				break
			}
		}
		totalTickets = append(totalTickets, Ticket{timeIn, timeOut, totalCost(timeIn, timeOut)})
	}

	sort.Slice(totalTickets, func(i, j int) bool {
		return totalTickets[i].timeIn < totalTickets[j].timeIn
	})

	return totalTickets
}
