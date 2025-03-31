package main

import "fmt"

func main() {

	s := make(map[string]int)

	s["a"] = s["a"] + 1
	s["a"] = s["a"] + 1
	s["a"] = s["a"] + 1
	s["a"] = s["a"] + 1
	s["a"] = s["a"] + 1
	s["a"] = s["a"] + 1

	fmt.Println(s["a"])
}
