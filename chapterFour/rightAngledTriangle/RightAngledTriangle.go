package main

import "fmt"

func main() {
	printTriangle()
}

func printTriangle() {
	base := collectTriangleBaseSize()
	for count := 1; count <= base; count++ {
		for index := 1; index <= count; index++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func collectTriangleBaseSize() int {
	base := 0
	fmt.Println("Enter the a size for the base of the triangle: ")
	_, err := fmt.Scan(&base)
	if err != nil {
		return 0
	}
	return base
}
