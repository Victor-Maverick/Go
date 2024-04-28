package main

import (
	"fmt"
	"strconv"
)

func main() {
	encryptData()
}

func encryptData() {
	userInput := collectInput()
	swappedDigits := swapDigits(userInput)
	returnDigit := ""

	for count := 0; count < len(swappedDigits); count++ {
		number, _ := strconv.ParseInt(string(swappedDigits[count]), 10, 64)
		number = (number + 7) % 10
		returnDigit += strconv.FormatInt(number, 10)
	}
	fmt.Println(returnDigit)
}

func swapDigits(input string) string {

	var swappedString string
	swappedString += string(input[2])
	swappedString += string(input[3])
	swappedString += string(input[0])
	swappedString += string(input[1])

	return swappedString

}

func collectInput() string {
	userInput := ""
	fmt.Println("Enter a four digit number")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return ""
	}

	for len(userInput) != 4 {
		fmt.Println("Enter a four digit number")
		_, err := fmt.Scan(&userInput)
		if err != nil {
			return ""
		}
	}
	return userInput
}
