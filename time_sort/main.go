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

	sort(randomTimes)

	fmt.Println("Sorted :")
	for _, t := range randomTimes {
		fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	}
}
func sort(times []time.Time) {

	for j := 0; j < len(times); j++ {

		for i := 0; i < len(times)-1-j; i++ {

			if times[i].Year() > (times[i+1]).Year() && times[i].Minute() > (int(times[i+1].Month())) && times[i].Day() > (int(times[i+1].Day())) && times[i].Minute() > (int(times[i+1].Minute())) {

				times[i], times[i+1] = times[i+1], times[i]

			}
		}
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
