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

	fmt.Println(numbers)

}
