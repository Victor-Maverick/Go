package main

import "fmt"

func main() {
	var numberOfTerms int
	fmt.Println("Enter number of terms you want")
	_, err := fmt.Scan(&numberOfTerms)
	if err != nil {
		return
	}
	for numberOfTerms < 0 {
		fmt.Println("Enter number of terms you want")
		_, err = fmt.Scan(&numberOfTerms)
		if err != nil {
			return
		}
	}
	constant := 1.0
	for count := 1; count < numberOfTerms; count++ {
		result := factorial(count)
		constant += 1 / float64(result)
	}
	fmt.Println(constant)
}

func factorial(count int) int {
	total := 1
	for i := 1; i <= count; i++ {
		total *= i
	}
	return total
}
