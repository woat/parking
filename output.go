package main

import (
	"bytes"
	"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
	"time"
)

const (
	height         = 13
	width          = 24
	divider        = 6
	amountOfFloors = 5
)

type Car struct {
	ticket Ticket
	spot   [3]int
}

type Floor struct {
	floorBox  *tm.Box
	floorPlan [][]interface{}
}

func isValidSpot(r, c, level int) bool {
	switch {
	case r == 0 || r == 12:
		return c >= 1 && c <= 22
	case r == 5:
		return c == 0 || (c >= 4 && c <= 19) || c == 23
	case r == 6:
		return c == 0 || c == 23
	case r == 7:
		if level == 0 {
			return c >= 1 && c <= 19 || c == 23
		}
		if level < 4 {
			return c == 0 || (c >= 4 && c <= 19) || c == 23
		}
		return c == 0 || (c >= 4 && c <= 22) || c == 22
	default:
		if level == 0 && r >= 5 {
			return c == 23
		}
		return c == 0 || c == 23
	}
}

func notTaken(r, c, level int, f []Floor) bool {
	return f[level].floorPlan[r][c] == 0
}

func takeSpot(r, c, level int, f []Floor, cars []Car, idx int) {
	cars[idx].spot = [3]int{level, r, c}
	f[level].floorPlan[r][c] = 1
}

func leaveSpot(r, c, level int, f []Floor, cars []Car, idx int) {
	f[level].floorPlan[r][c] = 0
	cars[idx].spot = [3]int{0, 0, 0}
}

func floorGrid(level int) [][]interface{} {
	var spots [][]interface{}
	for i := 0; i < height; i++ {
		var row []interface{}
		for j := 0; j < width; j++ {
			if level == 0 && i == divider && j >= 0 && j <= 19 {
				row = append(row, "-")
				continue
			}
			if level != 0 && i == divider && j >= 4 && j <= 22 {
				row = append(row, "-")
				continue
			}
			if isValidSpot(i, j, level) {
				row = append(row, 0)
				continue
			}
			row = append(row, ".")
		}
		spots = append(spots, row)
	}
	return spots
}

func readFloor(grid [][]interface{}) string {
	var buffer bytes.Buffer

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			s := fmt.Sprintf("%v", grid[i][j])
			buffer.WriteString(s)
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func assignCar(t Ticket) Car {
	return Car{t, [3]int{0, 0, 0}}
}

func assignSpots(c []Car, f []Floor, i int) {
	level := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 3, 4, 5}
	for idx, car := range c {
		if car.ticket.timeIn == i {
			for {
				level := level[rand.Intn(cap(level)-1)]
				row := rand.Intn(13)
				col := rand.Intn(24)
				if isValidSpot(row, col, level) && notTaken(row, col, level, f) {
					takeSpot(row, col, level, f, c, idx)
					break
				}
			}
		}
	}
}

func removeSpots(c []Car, f []Floor, i int) {
	for idx, car := range c {
		if i == car.ticket.timeOut {
			leaveSpot(car.spot[1], car.spot[2], car.spot[0], f, c, idx)
			updateRevenue(totalCost(car.ticket.timeIn, car.ticket.timeOut))
		}
	}
}

func updateRevenue(sale int) {
	totalMoneyGenerated += sale
}

func createFloors(f *[]Floor) {
	for j := 0; j < amountOfFloors; j++ {
		*f = append(*f, Floor{tm.NewBox(28, 15, 0), floorGrid(j)})
	}
}

var totalMoneyGenerated int

func start(tickets []Ticket) {
	tm.Clear()

	var floors []Floor
	var cars []Car
	createFloors(&floors)
	for _, ticket := range tickets {
		car := assignCar(ticket)
		cars = append(cars, car)
	}

	i := 630
	for ; i <= 2101; i++ {
		tm.Clear()
		if i%100 >= 60 {
			i += 39
			continue
		}
		tm.MoveCursor(1, 1)
		tm.Println("Monday:", i-1)
		tm.Println("Total money generated: $", totalMoneyGenerated)
		assignSpots(cars, floors, i)
		for idx, floor := range floors {
			floor.floorBox = tm.NewBox(28, 15, 0)
			x := 33 * idx
			fmt.Fprint(floor.floorBox, readFloor(floor.floorPlan))
			tm.Print(tm.MoveTo(tm.Color(fmt.Sprintf("Floor %d", idx+1), tm.CYAN), x, 4))
			tm.Print(tm.MoveTo(floor.floorBox.String(), x, 5))
			time.Sleep(time.Millisecond)
		}
		tm.Flush()
		removeSpots(cars, floors, i)
		time.Sleep(time.Millisecond)
	}
}
