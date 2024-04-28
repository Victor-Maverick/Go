package main

import "fmt"

func main() {
	printShapes()
}

func printShapes() {
	for count := 0; count < 10; count++ {
		for index := 0; index <= count; index++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
	for count := 10; count > 0; count-- {
		for index := 0; index < count; index++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()

	for count := 10; count > 0; count-- {

		for index := count; index > count; index-- {
			fmt.Print(" ")
		}

		for index := 0; index < count; index++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}

}
