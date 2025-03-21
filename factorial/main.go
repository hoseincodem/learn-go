package main

import "fmt"

func main() {

	factorial(4)

	fmt.Println(factorial2(4))

}

func factorial(a int) int {

	if a <= 1 {
		return 1
	}
	s := 1

	for i := a; i >= 1; i-- {
		s = s * i
	}

	fmt.Println(s)

	return s
}

func factorial2(a int) int {

	if a == 1 {
		return 1
	}

	return a * factorial2(a-1)

}

/*

	1,1,2,3,5,8,13,21,...

	fib(1) = 1
	fib(2) = 1

	for n >= 3{

		fib(n) = fib(n-1) + fib(n-2)
	}


*/
