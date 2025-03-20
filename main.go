package main

import "fmt"

func main() {

	word := "AAA"

	reversedWord := reverse(word)
	fmt.Println(word)
	fmt.Println(reversedWord)

	if word == reversedWord {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func reverse(s string) string {

	var d string

	for i := 1; i <= len(s); i++ {

		d += string(s[len(s)-i])

	}

	return d
}

/*

	AsssA -> True
	AssaA -> False

	len(s) -> 8
	A B C D E F G H
	0 1 2 3 4 5 6 7

	7 -> 8 - 1
	6 -> 8 - 2
	5 -> 8 - 3
	4 -> 8 - 4
	3 -> 8 - 5
	2 -> 8 - 6
	1 -> 8 - 7

*/
