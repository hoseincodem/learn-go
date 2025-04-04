package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {

	numbers := make([]int, 50)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}
	fmt.Println(numbers)

	fmt.Println(sort(numbers))

}

func sort(n []int) []int {

	b := make([]int, 0)

	lenOfSlice := len(n)

	for i := 0; i < lenOfSlice; i++ {

		minNumber, minIndex := minWithIndex(n)

		b = append(b, minNumber)

		n = slices.Delete(n, minIndex, minIndex+1)

	}

	return b
}

func minWithIndex(n []int) (int, int) {

	minNumber := n[0]
	minIndex := 0

	for i := 0; i < len(n); i++ {

		if n[i] < minNumber {
			minNumber = n[i]
			minIndex = i
		}
	}

	return minNumber, minIndex
}
