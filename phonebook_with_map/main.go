package main

import "fmt"

func main() {

	phonebook := make(map[string]string)

	phonebook["ahmad"] = "090377184612"
	phonebook["hosein"] = "090377184612"
	phonebook["ali"] = "08888888888"

	for {

		a := showMenu()

		if a == "add" || a == "1" {
			phonebook = add(phonebook)
			view(phonebook)
		} else if a == "view" || a == "2" {
			view(phonebook)
		} else if a == "remove" || a == "3" {
			phonebook = remove(phonebook)
			view(phonebook)
		} else if a == "search" || a == "4" {
			search(phonebook)
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

func add(phonebook map[string]string) map[string]string {

	var a string
	var b string

	fmt.Print("Enter The name : ")
	fmt.Scan(&a)

	fmt.Print("Enter The phone :")
	fmt.Scan(&b)

	phonebook[a] = b

	return phonebook
}

func view(phoneBook map[string]string) {

	fmt.Println("names and phones")
	fmt.Println("---------------------------------------------------------------------")
	for key, value := range phoneBook {

		fmt.Println("| name is: ", key, "\t\t| phone is: ", value)
		fmt.Println("---------------------------------------------------------------------")
	}

}

func remove(phoneBook map[string]string) map[string]string {

	var a string

	fmt.Println("Enter the name you want to remove")
	fmt.Scan(&a)

	delete(phoneBook, a)

	return phoneBook
}

func search(phoneBook map[string]string) {

	var a string

	fmt.Println("Enter Your Name :")
	fmt.Scan(&a)

	value, ok := phoneBook[a]
	if ok {
		fmt.Println("| name is: ", a, "\t\t| phone is: ", value)
		return
	}

	fmt.Println("Not Found.")
	return
}
