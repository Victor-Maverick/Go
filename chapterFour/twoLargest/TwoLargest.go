package main

import "fmt"

func main() {
	findTwoLargest()
}

func findTwoLargest() {
	largest := 0
	number := 0
	secondLargest := 0
	for count := 1; count <= 10; count++ {
		fmt.Println("Enter a number")
		fmt.Scan(&number)
		if number > largest {
			secondLargest = largest
			largest = number
		}
		
	}
	fmt.Printf("The largest number is %d and the second largest number is %d.\n", largest, secondLargest)
}
