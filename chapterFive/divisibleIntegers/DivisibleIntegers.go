package main

import "fmt"

func main() {
	sumOfThreeDivisibles := 0
	for count := 1; count < 30; count++ {
		if count%3 == 0 {
			sumOfThreeDivisibles += count
		}
	}
	fmt.Println("The sum of integers between 1 and 30 that are divisible by 3 are", sumOfThreeDivisibles)
}
