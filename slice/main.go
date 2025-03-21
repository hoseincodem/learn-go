package main

import (
	"fmt"
	"math/rand"
)

func main() {

	numbers := make([]int, 0)

	numbers = append(numbers, 20, 30, 40, 50)
	fmt.Println(numbers)

	//---------------------------------------------------------------

	var numbers2 []int

	numbers2 = append(numbers2, 20, 30, 40, 50)
	fmt.Println(numbers2)

	//---------------------------------------------------------------

	numbers3 := []int{}

	numbers3 = append(numbers3, 20, 30, 40, 50)
	fmt.Println(numbers3)

	//---------------------------------------------------------------
	// make(type of slice, len(length) of slice, cap(capacity) of slice)

	numbers4 := make([]int, 30)

	for i := 0; i < 30; i++ {
		numbers4[i] = rand.Intn(100)
	}
	fmt.Println(numbers4)

}
