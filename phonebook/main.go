package main

import (
	"fmt"
	"slices"
)

func main() {
	name := make([]string, 0)
	phone := make([]string, 0)

	var a string

	fmt.Println("1_add - 2_view - 3_remove- 4_search")

	fmt.Scanln(&a)

	if a == "add" {
		add(name, phone)
	} else if a == "view" {

		view(name, phone)

	} else if a == "remove" {
		remove(name, phone)
	} else if a == "search" {
		search(name)
	}

}
func add(n []string, p []string) ([]string, []string) {

	fmt.Println("name :")
	fmt.Scan(&n)

	fmt.Println("nuberPhone :")
	fmt.Scan(&p)

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
func search(n []string) []string {
	var a string
	fmt.Scanln(a)
	for i := 0; i < len(n); i++ {

		if a == n[i] {
			fmt.Println(n[i])

		} else {
			fmt.Println("Nist.")
		}
	}
	return n
}
