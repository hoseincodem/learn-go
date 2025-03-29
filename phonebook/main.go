package main

import (
	"fmt"
	"slices"
)

func main() {
	name := make([]string, 0)
	phone := make([]string, 0)

	name = append(name, "ahmad")
	phone = append(phone, "0903771846")

	var a string

	for {
		fmt.Println("1_add - 2_view - 3_remove- 4_search")

		fmt.Scanln(&a)

		if a == "add" {
			name, phone = add(name, phone)

			fmt.Println(name, phone)

		} else if a == "view" {

			name, phone = view(name, phone)

		} else if a == "remove" {
			name, phone = remove(name, phone)
		} else if a == "search" {
			search(name, phone)
		}
	}
}
func add(n []string, p []string) ([]string, []string) {

	var a string
	var c string

	fmt.Println("name :")

	fmt.Scan(&a)

	n = append(n, a)

	fmt.Println("nuberPhone :")
	fmt.Scan(&c)
	p = append(p, c)

	return n, p
}
func view(n []string, p []string) ([]string, []string) {
	fmt.Println("Name :", n)
	fmt.Println("Phone :", p)
	return n, p
}
func remove(n []string, p []string) ([]string, []string) {
	var f string
	fmt.Scanln(&f)

	for i := 0; i < len(n); i++ {

		if f == n[i] {
			n = slices.Delete(n, i, i)
			p = slices.Delete(p, i, i)
		}

	}
	return n, p
}
func search(n []string, p []string) ([]string, []string) {
	var a string
	fmt.Scanln(&a)
	for i := 0; i < len(n); i++ {

		if a == n[i] {
			fmt.Println(p[i])

		} else {
			fmt.Println("Nist.")
		}
	}
	return n, p
}
