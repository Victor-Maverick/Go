package main

import "fmt"

func main() {
	var bars [5]int
	number := 0
	for count := 0; count < len(bars); count++ {
		fmt.Println("Enter a positive number")
		_, err := fmt.Scan(&number)
		if err != nil {
			return
		}

		for number < 1 {
			fmt.Println("Enter a positive number")
			_, err := fmt.Scan(&number)
			if err != nil {
				return
			}
		}
		bars[count] = number
	}
	for count := 0; count < len(bars); count++ {
		for index := 1; index <= bars[count]; index++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
