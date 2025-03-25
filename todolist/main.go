package main

import (
	"fmt"
	"slices"
)

func main() {

	tasks := make([]string, 0)

	tasks = append(tasks, "raften be bashkah", "f9jisvhvd")

	for i := 0; i == 0; {

		var option string

		fmt.Println("1 Add a task")

		fmt.Println("2 View Tasks")

		fmt.Println("3 Remove a Tasks")

		fmt.Println("4 Exit")

		fmt.Scanln(&option)

		if option == "1" {

			tasks = AddTask(tasks)

		} else if option == "2" {
			ViewTask(tasks)

		} else if option == "3" {

			fmt.Println("slice???")
			ViewTask(tasks)

			var a int
			fmt.Scanln(&a)

			tasks = slices.Delete(tasks, a-1, a)

		} else if option == "4" {
			i = 1

		} else {
			fmt.Println("1- 2 -3 -4")
		}
	}
}

func AddTask(todoList []string) []string {

	var a string

	fmt.Scanln(&a)

	todoList = append(todoList, a)

	return todoList
}
func ViewTask(todoList []string) {

	fmt.Println(todoList)

}
