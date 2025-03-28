package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("------------------------------- MAPS ----------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------")
	// how to make maps
	m := make(map[string]string, 10) // map[type1]type2, (size of map)
	_ = map[string]string{}

	m["a"] = "A"
	m["b"] = "B"
	m["c"] = "C"
	m["d"] = "D"
	m["e"] = "E"
	m["f"] = "F"
	m["g"] = "G"
	m["h"] = "H"

	// looping in map : key and value instead of index and value that we have in slices
	for k, v := range m {
		fmt.Println("key is: ", k, "value is: ", v)
	}

	// how to check if an element is in the map?
	ok, v := m["a"]
	fmt.Println(ok, v)

	ok, v = m["s"]
	fmt.Println(ok, v)

	// how to delete from map
	delete(m, "a")
	fmt.Println(m)

	// in slices to make comparisons
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("----------------------------- SLICES ----------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------")

	// how to make slices
	s := make([]string, 8, 10) // slice of strings of length 8 and capacity 10
	_ = []string{}
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"
	s[5] = "f"
	s[6] = "g"
	s[7] = "h"
	fmt.Println(s)

	// looping in slice : index and value instead of key and value that we have in maps
	for i, v := range s {
		fmt.Println("index is: ", i, "value is: ", v)
	}

	// how to check if an element is in the slice
	for i := range s {
		if s[i] == "u" {
			fmt.Println("true")
		}
	}

	// how to delete from slice

	// first way
	s = append(s[:2], s[3:]...)

	// second way
	s = slices.Delete(s, 2, 3)

	fmt.Println(s)

}
