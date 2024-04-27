package main

import "fmt"

func main() {
	processTax()
}

func processTax() {
	firstCitizenName := ""
	firstCitizenEarning := 0.0
	totalTax1 := 0.0
	fmt.Println("Enter the name of the first citizen")
	_, err := fmt.Scan(&firstCitizenName)
	if err != nil {
		return
	}

	fmt.Println("Enter the yearly earning of ", firstCitizenName)
	_, err = fmt.Scan(&firstCitizenEarning)
	if err != nil {
		return
	}
	if firstCitizenEarning > 0.0 && firstCitizenEarning <= 30000.0 {
		totalTax1 = firstCitizenEarning * 0.15
	} else if firstCitizenEarning > 30000.0 {
		totalTax1 = firstCitizenEarning * 0.2
	}

	secondCitizenName := ""
	secondCitizenEarning := 0.0
	totalTax2 := 0.0
	fmt.Println("Enter the name of the second citizen")
	fmt.Scan(&secondCitizenName)

	fmt.Println("Enter the monthly earning of ", secondCitizenName)
	fmt.Scan(&secondCitizenEarning)
	if secondCitizenEarning > 30000.0 {
		totalTax2 = secondCitizenEarning * 0.2
	} else if secondCitizenEarning > 0.0 && secondCitizenEarning <= 30000.0 {
		totalTax2 = secondCitizenEarning * 0.15
	}

	thirdCitizenName := ""
	thirdCitizenEarning := 0.0
	totalTax3 := 0.0
	fmt.Println("Enter the name of the third citizen")
	_, err = fmt.Scan(&thirdCitizenName)
	if err != nil {
		return
	}

	fmt.Println("Enter the yearly earning of ", thirdCitizenName)
	_, err = fmt.Scan(&thirdCitizenEarning)
	if err != nil {
		return
	}
	if thirdCitizenEarning > 30000.0 {
		totalTax3 = thirdCitizenEarning * 0.2
	} else if thirdCitizenEarning > 0.0 && thirdCitizenEarning <= 30000.0 {
		totalTax3 = thirdCitizenEarning * 0.15
	}

	fmt.Print("Name\t\tTotalTax\n", firstCitizenName, "\t\t", totalTax1, "\n", secondCitizenName, "\t\t", totalTax2, "\n", thirdCitizenName, "\t\t", totalTax3)

}
