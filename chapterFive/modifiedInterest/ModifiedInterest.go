package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0

	fmt.Println("Year\t\tRate\t\t\tAmount")
	for year := 1; year <= 10; year++ {
		for rate := 0.05; rate <= 0.10; rate += 0.01 {

			amount := principal * math.Pow(1.0+rate, float64(year))

			fmt.Printf("%d\t\t\t%.2f\t\t\t%.2f\n", year, rate, amount)

		}
	}

}
