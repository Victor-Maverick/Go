package main

import "fmt"

func main() {
	print()
}

func print() {

	for length := 1; length <= 10; length++ {
		for input := 1; input <= length; input++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	for length := 1; length <= 10; length++ {
		for input := length; input <= 10; input++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for length := 1; length <= 10; length++ {
		for input := length; input <= 10; input++ {
			fmt.Print(" ")
		}
		for length1 := 1; length1 <= length; length1++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
	fmt.Println()
	for length := 1; length <= 10; length++ {
		for input := 1; input <= length; input++ {
			fmt.Print(" ")
		}
		for length1 := length; length1 <= 10; length1++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
