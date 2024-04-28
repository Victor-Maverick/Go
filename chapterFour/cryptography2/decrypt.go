package main

import "fmt"

func main() {
	decryptData()
}

func decryptData() {
	collectEncryptedData()
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
