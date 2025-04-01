package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	start := time.Date(2005, 3, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 3, 31, 23, 59, 59, 0, time.UTC)
	randomTimes := generateRandomTimes(10, start, end)

	fmt.Println("Generated random times:")

	for _, t := range randomTimes {
		fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	}

	// Sort these times

	// sort(randomTimes)

	fmt.Println("Sorted :")
	for _, t := range sortTime(randomTimes) {
		fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	}
}

func generateRandomTimes(n int, start time.Time, end time.Time) []time.Time {

	var times []time.Time
	duration := end.Sub(start).Seconds()

	for range n {
		randomOffset := time.Duration(rand.Float64() * duration * float64(time.Second))
		randomTime := start.Add(randomOffset)
		times = append(times, randomTime)
	}

	return times
}

// SORT TIMES
func sortTime(times []time.Time) []time.Time {

	timeMap := map[int]time.Time{}

	intTimes := []int{}

	for _, t := range times {

		intTime := convertTimeToInt(t)

		timeMap[intTime] = t

		intTimes = append(intTimes, intTime)
	}

	for i := 0; i < len(intTimes); i++ {

		for j := i + 1; j < len(intTimes); j++ {

			if intTimes[i] > intTimes[j] {
				intTimes[i], intTimes[j] = intTimes[j], intTimes[i]
			}
		}
	}

	for i := range times {

		times[i] = timeMap[intTimes[i]]

	}

	return times

}

// CONVERT TIME TO INT
func convertTimeToInt(t time.Time) int {

	year := t.Year()
	monthString := t.Month().String()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	m := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	month := m[monthString]

	var a int

	a += year
	a = a * 100
	a = a + month

	a = a * 100
	a = a + day

	a = a * 100
	a = a + hour

	a = a * 100
	a = a + minute

	a = a * 100
	a = a + second

	return a
}
