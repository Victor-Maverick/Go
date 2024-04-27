package main

import "fmt"

func main() {
	printTable()
}

func printTable() {
	fmt.Printf("N		N2		N3		N4\n")
	for count := 1; count <= 5; count++ {
		fmt.Printf("%d\t\t%d\t\t%d\t\t%d\t\t\n", count, count*count, count*count*count, count*count*count*count)
	}
}
