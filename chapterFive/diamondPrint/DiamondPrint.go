package main

import "fmt"

func main() {
	printShape(5)
}

func printShape(numberOfCols int) {
	for i := 0; i < numberOfCols*2-1; i++ {
		noOfSpaces := 0
		if i > numberOfCols {
			noOfSpaces = i - numberOfCols
		} else {
			noOfSpaces = numberOfCols
		}
		for j := noOfSpaces; j > 0; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
