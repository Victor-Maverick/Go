package main

import "fmt"

func main() {
	displayWorldPopulation()
}

func displayWorldPopulation() {
	worldPopulation := 7975105156.00
	increaseRate := 0.89
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("Year\t\tExpected Population\t\t\tNumeric Increase")
	fmt.Println("-------------------------------------------------------------")
	for count := 1; count <= 75; count++ {
		increase := increaseRate * worldPopulation
		worldPopulation += increase
		fmt.Println(count, "\t\t", worldPopulation, "\t\t", increase)
	}
	fmt.Println("-------------------------------------------------------------")
}
