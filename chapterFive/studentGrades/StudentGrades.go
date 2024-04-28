package main

import (
	"fmt"
	"strings"
)

func main() {
	aCount := 0
	bCount := 0
	cCount := 0
	dCount := 0
	grade := ""
	for count := 0; count < 5; count++ {
		fmt.Println("Enter the grade for student ", count+1)
		_, err := fmt.Scan(&grade)
		if err != nil {
			return
		}
		switch strings.ToLower(grade) {
		case "a":
			aCount++
		case "b":
			bCount++
		case "c":
			cCount++
		case "d":
			dCount++
		default:
			fmt.Println("Enter grade from A-D only for student ", count+1)
			_, err := fmt.Scan(&grade)
			if err != nil {
				return
			}
		}
	}
	fmt.Printf("Number of A's\t%d\nNumber of B's\t%d\nNumber of C's\t%d\nNumber of D's\t%d\n", aCount, bCount, cCount, dCount)

}
