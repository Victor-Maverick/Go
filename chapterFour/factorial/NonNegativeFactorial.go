package main

import "fmt"

func main() {
	number := 0

	fmt.Print("Enter a positive number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}

	for number < 1 {
		fmt.Print("Enter a positive number: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			return
		}
	}
	factorial := 1
	for count := 1; count <= number; count++ {
		factorial *= count
	}
	fmt.Printf("Factorial of %d is %d", number, factorial)
}
