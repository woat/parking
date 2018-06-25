package main

import "testing"

func TestTimeDifference(t *testing.T) {
	tables := []struct {
		timeIn  int
		timeOut int
		delta   int
	}{
		{1630, 1800, 130},
		{1600, 1800, 200},
		{1600, 1830, 230},
		{1645, 1915, 230},
		{900, 2100, 1200},
		{915, 1255, 340},
	}

	for _, tb := range tables {
		diff := timeDifference(tb.timeIn, tb.timeOut)
		if diff != tb.delta {
			t.Errorf("timeDifference(%d, %d) => %d, wanted: %d", tb.timeIn, tb.timeOut, diff, tb.delta)
		}
	}
}

func TestTotalCost(t *testing.T) {
	tables := []struct {
		timeIn  int
		timeOut int
		cost    int
	}{
		{1630, 1800, 4},
		{1600, 1800, 4},
		{1600, 1830, 6},
		{1645, 1915, 6},
		{900, 2100, 24},
		{915, 1255, 8},
	}

	for _, tb := range tables {
		diff := totalCost(tb.timeIn, tb.timeOut)
		if diff != tb.cost {
			t.Errorf("totalCost(%d, %d) => %d, expected: %d", tb.timeIn, tb.timeOut, diff, tb.cost)
		}
	}
}
