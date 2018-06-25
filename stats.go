package main

var timeMap = map[string]int{
	"sixAM":    0,
	"sevenAM":  0,
	"eightAM":  0,
	"nineAM":   0,
	"tenAM":    0,
	"elevenAM": 0,
	"twelveAM": 0,
	"onePM":    0,
	"twoPM":    0,
	"threePM":  0,
	"fourPM":   0,
	"fivePM":   0,
	"sixPM":    0,
	"sevenPM":  0,
	"eightPM":  0,
}

func getStats(m map[string]int, y []int) map[string]int {
	for _, time := range y {
		switch {
		case time > 2000:
			m["eightPM"] = m["eightPM"] + 1
		case time > 1900:
			m["sevenPM"] = m["sevenPM"] + 1
		case time > 1800:
			m["sixPM"] = m["sixPM"] + 1
		case time > 1700:
			m["fivePM"] = m["fivePM"] + 1
		case time > 1600:
			m["fourPM"] = m["fourPM"] + 1
		case time > 1500:
			m["threePM"] = m["threePM"] + 1
		case time > 1400:
			m["twoPM"] = m["twoPM"] + 1
		case time > 1300:
			m["onePM"] = m["onePM"] + 1
		case time > 1200:
			m["twelveAM"] = m["twelveAM"] + 1
		case time > 1100:
			m["elevenAM"] = m["elevenAM"] + 1
		case time > 1000:
			m["tenAM"] = m["tenAM"] + 1
		case time > 900:
			m["nineAM"] = m["nineAM"] + 1
		case time > 800:
			m["eightAM"] = m["eightAM"] + 1
		case time > 700:
			m["sevenAM"] = m["sevenAM"] + 1
		case time > 600:
			m["sixAM"] = m["sixAM"] + 1
		}
	}
	return m
}
