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

	fmt.Println(sort(slices.Clone(numbers)))
	fmt.Println(sort2(slices.Clone(numbers)))

}

func sort(n []int) []int {

	b := make([]int, 0)
	c := make([]int, len(n))

	lenOfSlice := len(n)

	for i := 0; i < lenOfSlice; i++ {

		maxNumber, maxIndex := maxWithIndex(n)

		b = append(b, maxNumber)

		n = slices.Delete(n, maxIndex, maxIndex+1)

	}

	// Reverse b

	for i := 0; i < len(b); i++ {

		c[i] = b[len(b)-1-i]

	}

	return c
}

func sort2(n []int) []int {

	b := make([]int, len(n))

	lenOfSlice := len(n)

	for i := 0; i < lenOfSlice; i++ {

		maxNumber, maxIndex := maxWithIndex(n)

		b[len(b)-1-i] = maxNumber

		n = slices.Delete(n, maxIndex, maxIndex+1)

	}

	return b
}

func maxWithIndex(n []int) (int, int) {

	maxNumber := n[0]
	maxIndex := 0

	for i := 0; i < len(n); i++ {

		if n[i] > maxNumber {
			maxNumber = n[i]
			maxIndex = i
		}
	}

	return maxNumber, maxIndex
}
