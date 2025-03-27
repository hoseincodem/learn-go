package main

import (
	"fmt"
	"math/rand"
)

func main() {

	list := []string{"rock", "paper", "scissors"}
	user := ""
	fmt.Println("Rock - Paper - Scissors")

	fmt.Scanln(&user)

	computer := list[rand.Intn(len(list))]

	if user == computer {

		fmt.Println("Equal")

	} else if user == "rock" && computer == "paper" {

		fmt.Println("You lost.")

	} else if user == "rock" && computer == "scissors" {
		fmt.Println("You Win.")

	} else if user == "paper" && computer == "rock" {
		fmt.Println("You Win.")

	} else if user == "paper" && computer == "scissors" {
		fmt.Println("You lost.")

	} else if user == "scissors" && computer == "rock" {
		fmt.Println("You lost.")

	} else if user == "scissors" && computer == "paper" {
		fmt.Println("You Win.")
	}

}
