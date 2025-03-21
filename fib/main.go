package main

import "fmt"

func main() {

	fmt.Println(fib(6))
	fmt.Println(fib2(6))

}

func fib(a int) int {

	if a == 1 || a == 2 {
		return 1
	}

	return fib(a-1) + fib(a-2)
}

func fib2(a int) int {

	s := 0
	f1 := 0
	f2 := 0

	for i := 1; i < a; i++ {

		if i <= 2 {
			f1 = 1
			f2 = 1
			s = 1
		}

		s = f1 + f2
		f1 = f2
		f2 = s

	}

	return s

}

/*

	1 1 2 3 5 8 13 21 34


*/
