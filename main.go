package main

import "fmt"

func main() {

	var firstNumber int
	var secondNumber int

	fmt.Println("Enter a number")
	fmt.Scanln(&firstNumber)

	fmt.Println("Enter another number")
	fmt.Scanln(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Println("a is bigger than b")
	} else if secondNumber > firstNumber {
		fmt.Println("b is bigger than a")
	} else if secondNumber == firstNumber {
		fmt.Println("a and b are equal")
	}

}
