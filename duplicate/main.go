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
	a := make([]int, 0)
	c := make([]int, 0)

	a, c = GetUnique(numbers)

	fmt.Println("unique: ", a, c)

}

func GetUnique(a []int) ([]int, []int) {
	b := make([]int, 0, 10)
	c := make([]int, 0, 10)

	for i := range a {

		if Search(a[i], b) == false {
			b = append(b, a[i])
		} else {
			c = append(c, a[i])
		}

	}

	return b, c

}

func Search(n int, a []int) bool {

	for i := range a {
		if a[i] == n {
			return true
		}
	}

	return false
}
