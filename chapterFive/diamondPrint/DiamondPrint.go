package main

import "fmt"

func main() {
	diamondPrint(5)
}

func diamondPrint(number int) {
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
