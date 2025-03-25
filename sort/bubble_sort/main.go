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

	fmt.Println(bubbleSort(numbers))

}

func bubbleSort(n []int) []int {

	for i := 0; i < len(n); i++ {

		for j := 0; j < len(n)-1-i; j++ {

			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}

		}

	}

	return n
}
