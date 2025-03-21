package main

import "fmt"

func main() {
	word := ""
	fmt.Println(IsPallindrum(word))

	fmt.Println()
}
func IsPallindrum(a string) bool {

	var d bool

	for i := 0; i < len(a)/2; i++ {

		if a[i] == a[len(a)-1-i] {

			d = true

		} else {
			d = false
		}

	}

	return d
}
