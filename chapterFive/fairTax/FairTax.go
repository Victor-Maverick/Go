package main

import "fmt"

func main() {
	calculateTax()
}

func calculateTax() {
	fmt.Println("Enter the expenses in the categories: \nHousing = ")
	houseExpense := 0
	_, err := fmt.Scan(&houseExpense)
	if err != nil {
		return
	}
	for houseExpense < 0 {
		fmt.Println("Enter the expense for Housing")
		_, err := fmt.Scan(&houseExpense)
		if err != nil {
			return
		}
	}

	fmt.Println("Enter the expenses on Food")
	foodExpense := 0
	_, err = fmt.Scan(&foodExpense)
	if err != nil {
		return
	}

	for foodExpense < 0 {
		fmt.Println("Enter the expenses on Food")
		_, err := fmt.Scan(&foodExpense)
		if err != nil {
			return
		}
	}
	fmt.Println("Enter the expenses on Clothing")
	clothingExpense := 0
	_, err = fmt.Scan(&clothingExpense)
	if err != nil {
		return
	}
	for clothingExpense < 0 {
		fmt.Println("Enter the expenses on Clothing")
		_, err := fmt.Scan(&clothingExpense)
		if err != nil {
			return
		}
	}

	fmt.Println("Enter the expense on Transportation ")
	transportExpense := 0
	_, err = fmt.Scan(&transportExpense)
	if err != nil {
		return
	}
	for transportExpense < 0 {
		fmt.Println("Enter the expenses on Transportation")
		_, err := fmt.Scan(&transportExpense)
		if err != nil {
			return
		}
	}

	fmt.Println("Enter the expense on Education = ")
	educationExpense := 0
	_, err = fmt.Scan(&educationExpense)
	if err != nil {
		return
	}
	for educationExpense < 0 {
		fmt.Println("Enter the expenses on Education")
		_, err := fmt.Scan(&educationExpense)
		if err != nil {
			return
		}
	}

	fmt.Println("Enter the expense on Health care = ")
	healthExpense := 0
	_, err = fmt.Scan(&healthExpense)
	if err != nil {
		return
	}
	for healthExpense < 0 {
		fmt.Println("Enter the expenses on Health")
		_, err := fmt.Scan(&healthExpense)
		if err != nil {
			return
		}
	}

	fmt.Println("Enter the expense on Vacations = ")
	vacationExpense := 0

	total := 0.0
	tax := 23.0
	fairTax := 0.0
	total = float64(houseExpense + foodExpense + clothingExpense + transportExpense + educationExpense + healthExpense + vacationExpense)
	fairTax = total * (tax * 0.01)

	fmt.Printf("The estimated FairTax you are to pay is %.2f", fairTax)
}
