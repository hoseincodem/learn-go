package main

import "fmt"

func main() {

	a := 10
	b := 20

	a, b = swap(a, b)

	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

func swap(a int, b int) (int, int) {

	var c int

	c = a
	a = b
	b = c

	return a, b

}
