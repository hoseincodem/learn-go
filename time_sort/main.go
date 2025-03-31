package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func main() {

	start := time.Date(2005, 3, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 3, 31, 23, 59, 59, 0, time.UTC)
	randomTimes := generateRandomTimes(10, start, end)

	fmt.Println("Generated random times:")
	for _, t := range randomTimes {
		fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	}

	// Sort these times
}
