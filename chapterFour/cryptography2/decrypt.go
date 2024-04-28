package main

import (
	"fmt"
	"strconv"
)

func main() {
	decryptData()
}

func decryptData() {
	userInput := collectEncryptedData()
	swappedDigits := swapDigits(userInput)
	returnDigit := ""

	for count := 0; count < len(swappedDigits); count++ {
		number, _ := strconv.ParseInt(string(swappedDigits[count]), 10, 64)
		number = (number + 3) % 10
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

func collectEncryptedData() string {
	userInput := ""
	fmt.Println("Enter the four-digit encrypted number")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return ""
	}

	for len(userInput) != 4 {
		fmt.Println("Enter the four-digit encrypted number")
		_, err := fmt.Scan(&userInput)
		if err != nil {
			return ""
		}

	}
	return userInput
}
