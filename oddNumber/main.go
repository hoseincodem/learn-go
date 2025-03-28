package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := make([]int, 100)

	for i := 0; i < len(number); i++ {
		number[i] = rand.Intn(100)
		odd := []int{}

		if number[i]%2 != 0 {
			odd = append(odd, i)
			fmt.Print(odd)
		}
	}
	fmt.Println(number)

}
