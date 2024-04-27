package main

import "fmt"

func main() {
	firstNumber := 0
	secondNumber := 0

	fmt.Println("Enter First Number: ")
	_, err := fmt.Scan(&firstNumber)
	if err != nil {
		return
	}

	fmt.Println("Enter Second Number: ")
	_, err = fmt.Scan(&secondNumber)
	if err != nil {
		return
	}

	if firstNumber == secondNumber {
		fmt.Println(0)
	} else if firstNumber > secondNumber {
		fmt.Println(1)
	} else if firstNumber < secondNumber {
		fmt.Println(-1)
	}
}
