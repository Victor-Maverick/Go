package main

import "fmt"

func main() {
	printDecimalEquivalent()
}

func printDecimalEquivalent() {
	binaryNumber := collectBinaryInput()
	counter := 1
	decimalNumber := 0
	for count := len(binaryNumber) - 1; count >= 0; count-- {
		fmt.Println(string(binaryNumber[count]))
		digit := int(binaryNumber[count] - '0')
		decimalNumber += digit * counter
		counter *= 2
	}
	fmt.Println("The decimal equivalent of ", binaryNumber, "is ", decimalNumber)
}

func collectBinaryInput() string {
	binaryNumber := ""
	fmt.Println("Enter a Binary Number:")
	_, err := fmt.Scan(&binaryNumber)
	if err != nil {
		return ""
	}
	for !isValidBinary(binaryNumber) {
		fmt.Println("Enter a Binary Number:")
		_, err := fmt.Scan(&binaryNumber)
		if err != nil {
			return ""
		}
	}
	return binaryNumber
}

func isValidBinary(binaryNumber string) bool {
	for _, char := range binaryNumber {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}
