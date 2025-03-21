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
}

func min(n []int) int {

	return 0
}

func max(n []int) int {

	return 0
}
