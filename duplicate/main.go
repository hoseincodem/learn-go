package main

import (
	"fmt"
	"math/rand"
)

func main() {

	numbers := make([]int, 10)

	for i := range numbers {
		numbers[i] = rand.Intn(6)
	}

	fmt.Println("numbers: ", numbers)
	fmt.Println("unique: ", GetUnique(numbers))

}

func GetUnique(a []int) []int {

	b := make([]int, 0, 10)

	for i := range a {

		if Search(a[i], b) == false {
			b = append(b, a[i])
		}
	}

	return b

}

func Search(n int, a []int) bool {

	for i := range a {
		if a[i] == n {
			return true
		}
	}

	return false
}
