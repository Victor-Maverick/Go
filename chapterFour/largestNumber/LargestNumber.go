package main

import "fmt"

func main() {
	findLargest()
}

func findLargest() {
	number := 0
	largestNumber := 0
	for count := 1; count <= 10; count++ {
		fmt.Println("Enter a number")
		_, err := fmt.Scan(&number)
		if err != nil {
			return
		}
		if number > largestNumber {
			largestNumber = number
		}
	}
	fmt.Println("The largest number is ", largestNumber)
}
