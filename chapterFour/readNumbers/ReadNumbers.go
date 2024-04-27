package main

import "fmt"

func main() {
	number := 0
	numbers := 0
	total := 0
	fmt.Println("Enter a number")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}
	for total < number {
		fmt.Println("Enter another number: ")
		_, err := fmt.Scan(&numbers)
		if err != nil {
			return
		}
		total += numbers
	}
}
