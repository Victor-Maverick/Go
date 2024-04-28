package main

import "fmt"

func main() {
	processInput()
}

func processInput() {
	numberOfInputs := collectNumberOfInputs()
	number := 0
	maxNumber := 0
	var minimum int
	for i := 0; i < numberOfInputs; i++ {
		fmt.Println("Enter a number")
		_, err := fmt.Scan(&number)
		if err != nil {
			return
		}
		if number > maxNumber {
			maxNumber = number
		}
		minimum = number
		if number < minimum {
			minimum = number
		}
	}
	fmt.Println("The maximum number is", maxNumber, "and the minimum number is", minimum)

}

func collectNumberOfInputs() int {
	numberOfInputs := 0
	fmt.Println("Enter number of inputs you want to enter")
	_, err := fmt.Scan(&numberOfInputs)
	if err != nil {
		return 0
	}
	for numberOfInputs < 1 {
		fmt.Println("Enter number of inputs you want to enter")
		_, err := fmt.Scan(&numberOfInputs)
		if err != nil {
			return 0
		}
	}
	return numberOfInputs
}
