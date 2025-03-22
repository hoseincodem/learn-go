package main

import (
	"fmt"
	"math/rand"
)

func main() {

	numbers := make([]int, 50)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}
	fmt.Println(numbers)
	fmt.Println(min(numbers))
	fmt.Println(max(numbers))

}

func min(n []int) int {

	minNumber := n[0]

	for i := 0; i < len(n); i++ {

		if n[i] < minNumber {
			minNumber = n[i]
		}

	}

	return minNumber
}

func max(n []int) int {

	MaxNumber := n[0]

	for i := 0; i < len(n); i++ {

		if n[i] > MaxNumber {
			MaxNumber = n[i]
		}

	}
	return MaxNumber
}
