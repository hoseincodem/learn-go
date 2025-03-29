package main

import (
	"fmt"
	"slices"
)

func main() {
	name := make([]string, 0)
	phone := make([]string, 0)

	name = append(name, "ahmad")
	phone = append(phone, "090377184612")
	name = append(name, "hosein")
	phone = append(phone, "090377184612")
	name = append(name, "ali")
	phone = append(phone, "08888888888")

	for {

		a := showMenu()

		if a == "add" || a == "1" {

			name, phone = add(name, phone)
			view(name, phone)

		} else if a == "view" || a == "2" {

			view(name, phone)

		} else if a == "remove" || a == "3" {

			name, phone = remove(name, phone)
			view(name, phone)

		} else if a == "search" || a == "4" {
			search(name, phone)
		}
	}
}

func showMenu() string {

	fmt.Println("1- add\n2- view\n3- remove\n4- search")
	fmt.Println("Enter what you want to do")

	var a string
	fmt.Scan(&a)

	return a

}
func add(n []string, p []string) ([]string, []string) {

	var a string
	var c string

	fmt.Print("ENTER The name : ")
	fmt.Scan(&a)

	n = append(n, a)

	fmt.Print("ENTER The phone :")
	fmt.Scan(&c)

	p = append(p, c)

	return n, p
}
func view(n []string, p []string) {

	fmt.Println("names and phones")
	fmt.Println("---------------------------------------------------------------------")
	for i := range n {

		fmt.Println("| name is: ", n[i], "\t\t| phone is: ", p[i])
		fmt.Println("---------------------------------------------------------------------")
	}

}
func remove(n []string, p []string) ([]string, []string) {

	var f string
	fmt.Println("Enter the name you want to remove")
	fmt.Scan(&f)

	for i := range n {

		if f == n[i] {
			n = slices.Delete(n, i, i+1)
			p = slices.Delete(p, i, i+1)
		}

	}
	return n, p
}
func search(n []string, p []string) {

	var a string
	var index int
	var found bool = false

	fmt.Println("Enter Your Name :")
	fmt.Scan(&a)

	for i := range n {
		if n[i] == a {
			found = true
			index = i
			break
		}
	}
	if found == true {
		fmt.Println("| name is: ", n[index], "\t\t| phone is: ", p[index])
		return
	}
	fmt.Println("Not Found.")
	return
}
