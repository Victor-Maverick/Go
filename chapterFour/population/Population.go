package main

func main() {
	displayWorldPopulation()
}

func displayWorldPopulation() {
	worldPopulation := 7975105156.00
	increaseRate := 0.89

	for count := 1; count <= 75; count++ {
		increase := increaseRate * worldPopulation
		worldPopulation += increase

	}
}
