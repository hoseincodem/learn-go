package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

func main() {

	List := []string{"Stone", "Paper", "Scissors"}

	sad := exec.Command("cmd", "/c", "cls")
	sad.Stdout = os.Stdout

	for i := 0; i == 0; {

		var user string
		fmt.Println("1 : Stone 2 : paper 3 : Scissors")

		fmt.Print("Enter an Option :")

		fmt.Scan(&user)

		compuder := List[rand.Intn(len(List))]

		if user == compuder {

			for {

				fmt.Println("User :" + user)
				fmt.Println("Compuder :" + compuder)
				fmt.Print("Enter an Option :")
				fmt.Scan(&user)
				sad.Run()

				break
			}

		} else if user == "stone" && compuder == "Scissors" {
			fmt.Println("User :" + user)
			fmt.Println("Compuder :" + compuder)
			fmt.Println("User Win")
			sad.Run()

		} else if user == "scissors" && compuder == "Stone" {
			fmt.Println("User :" + user)
			fmt.Println("Compuder :" + compuder)
			fmt.Println("Compuder Win")
			sad.Run()

		} else if user == "paper" && compuder == "Scissors" {

			fmt.Println("User :" + user)
			fmt.Println("Compuder :" + compuder)
			fmt.Println("Compuser Win")
			sad.Run()

		} else if user == "scissors" && compuder == "Paper" {

			fmt.Println("User :" + user)
			fmt.Println("Compuder :" + compuder)
			fmt.Println("User Win")
			sad.Run()

		} else if user == "stone" && compuder == "Paper" {

			fmt.Println("User :" + user)
			fmt.Println("Compuder :" + compuder)
			fmt.Println("Compuder Win")
			sad.Run()

		} else if user == "paper" && compuder == "Stone" {

			fmt.Println("User :" + user)
			fmt.Println("Compuder :" + compuder)
			fmt.Println("User Win")
			sad.Run()

		}

	}

}
