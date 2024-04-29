package main

import "fmt"

func main() {
	printShape()
}

func printShape() {
	number := collectInput()
	number = number/2 + 1
	for row := 0; row < 2*number; row++ {
		col := 0
		if row > number {
			col = 2*number - row
		} else {
			col = row
		}
		numberOfSpaces := number - col
		for space := 0; space < numberOfSpaces; space++ {
			fmt.Print(" ")
		}

		for column := 0; column < col; column++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}

func collectInput() int {
	number := 0
	fmt.Println("Enter an odd number from 1 to 19")
	_, err := fmt.Scan(&number)
	if err != nil {
		return 0
	}

	for number > 19 || number < 1 && number%2 != 1 {
		fmt.Println("Enter an odd number from 1 to 19")
		_, err := fmt.Scan(&number)
		if err != nil {
			return 0
		}
	}
	return number
}
