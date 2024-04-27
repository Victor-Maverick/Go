package main

import "fmt"

func main() {
	checkForPalindrome()
}

func checkForPalindrome() {
	reverseNumber := ""
	number := collectNumber()
	for count := len(number) - 1; count >= 0; count-- {
		reverseNumber += string(number[count])
	}
	if number == reverseNumber {
		fmt.Println(number, "is a Palindrome")
	} else {
		fmt.Println(number, "is not a Palindrome")
	}
}

func collectNumber() string {
	number := ""
	fmt.Print("Enter a five-digit number: ")
	fmt.Scan(&number)
	for len(number) != 5 {
		fmt.Print("Enter a five-digit number: ")
		fmt.Scan(&number)
	}
	return number
}
