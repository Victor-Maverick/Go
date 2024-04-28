package main

import (
	"fmt"
	"math"
)

func main() {
	getExponential()
}

func getExponential() {
	numberOfTerms := collectNumberOfTerms()
	constant := 1.0
	for i := 1; i <= numberOfTerms; i++ {
		result := math.Pow(float64(numberOfTerms), float64(i))
		factorialProduct := factorial(i)
		constant += result / float64(factorialProduct)
	}
	fmt.Println("The value of constant is ", constant)
}

func collectNumberOfTerms() int {
	numberOfTerms := 0
	fmt.Println("Enter number of terms you want to calculate for: ")
	_, err := fmt.Scan(&numberOfTerms)
	if err != nil {
		return 0
	}
	return numberOfTerms
}

func factorial(count int) int {
	total := 1
	for i := 1; i <= count; i++ {
		total *= i
	}
	return total
}
