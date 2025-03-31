package main

import (
	"fmt"
)

func main() {

	var password string
	fmt.Print("Enter Your Password :")
	fmt.Scanln(&password)

	if len(password) < 8 {

		fmt.Println("Your Password Is Not Big Enonugh.")

	}
	if !HasSmallLetters(password) {
		fmt.Println("Your Password Does Not Have Small Letters.")

	}
	if !HasCaptalLetters(password) {
		fmt.Println("Your Password Does Not Have Captal Letters.")

	}
	if !HasSign(password) {
		fmt.Println("Your Password Does Not Have Punctuations.")

	}
	if !HasNumber(password) {
		fmt.Println("Your Password Does Not Have Numbers.")

	}
	if len(password) >= 8 && HasSmallLetters(password) && HasCaptalLetters(password) && HasSign(password) && HasNumber(password) {
		fmt.Println("Ok!")
	}

}
func HasSmallLetters(p string) bool {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for i := 0; i < len(letters); i++ {

		for j := 0; j < len(p); j++ {
			if string(p[j]) == letters[i] {
				return true
			}
		}

	}
	return false
}
func HasCaptalLetters(p string) bool {
	LETTERS := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i := 0; i < len(LETTERS); i++ {

		for j := 0; j < len(p); j++ {
			if string(p[j]) == LETTERS[i] {
				return true
			}
		}

	}
	return false
}
func HasSign(p string) bool {
	sign := []string{"!", "@", "#", "$", "&", "*", "%"}

	for i := 0; i < len(sign); i++ {

		for j := 0; j < len(p); j++ {
			if string(p[j]) == sign[i] {
				return true
			}
		}

	}
	return false
}
func HasNumber(p string) bool {
	number := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := 0; i < len(number); i++ {

		for j := 0; j < len(p); j++ {
			if string(p[j]) == number[i] {
				return true
			}
		}

	}
	return false
}
