package main

import "fmt"

func main() {
	name := make([]string, 0)
	phone := make([]string, 0)

	var a string

	fmt.Println("1_add - 2_view - 3_remove- 4_search")

	if a == "add" {
		add(name, phone)
	} else if a == "view" {

		view(name, phone)

	} else if a == "remove" {
		remove(name, phone)
	}

}
func add(n []string, p []string) {

	fmt.Println("name :")
	fmt.Scan(n)

	fmt.Println("nuberPhone :")
	fmt.Scan(p)

}
func view(n []string, p []string) {
	fmt.Println("Name :", n)
	fmt.Println("Phone :", p)
}
func remove(n []string, p []string) {

}
func search() {

}
