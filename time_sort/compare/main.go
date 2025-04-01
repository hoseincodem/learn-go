package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	// Generate random times
	start := time.Date(2005, 3, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 3, 31, 23, 59, 59, 0, time.UTC)

	var randomTimes []time.Time
	duration := end.Sub(start).Seconds()

	for range 10 {
		randomOffset := time.Duration(rand.Float64() * duration * float64(time.Second))
		randomTime := start.Add(randomOffset)
		randomTimes = append(randomTimes, randomTime)
	}

	fmt.Println("Generated random times:")

	for _, t := range randomTimes {
		fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	}

	//-----------------------------------------------

}
