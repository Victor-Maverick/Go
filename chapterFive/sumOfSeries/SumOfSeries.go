package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("n\t\t\tSum")
	for count := 1; count <= 100; count++ {
		sum += count
		fmt.Println(count, "\t\t\t", sum)
	}
}
