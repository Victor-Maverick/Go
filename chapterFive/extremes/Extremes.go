package main

import "fmt"

func main() {
	processInput()
}

func processInput() {
	numberOfInputs := collectNumberOfInputs()
	maximum := 0
	var minimum int64
	for i := 0; i < numberOfInputs; i++ {

	}

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
