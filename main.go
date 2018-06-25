package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	totalTickets := totalTicketsGenerator()
	fmt.Println(len(totalTickets))
	start(totalTickets)
}
