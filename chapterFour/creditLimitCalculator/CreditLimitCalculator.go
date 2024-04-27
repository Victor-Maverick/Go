package main

import "fmt"

var accountNumber string

func main() {
	processCreditLimit()
}

func collectAccountNumber() string {
	var accountNumber string
	fmt.Println("Enter your account number: ")
	_, err := fmt.Scan(&accountNumber)
	if err != nil {
		return ""
	}
	for len(accountNumber) != 10 {
		fmt.Println("Enter your account number: ")
		_, err := fmt.Scan(&accountNumber)
		if err != nil {
			return ""
		}
	}
	return accountNumber
}

func collectBeginningBalance() float64 {
	var balance float64
	fmt.Println("Enter your beginning balance: ")
	_, err := fmt.Scan(&balance)
	if err != nil {
		return 0
	}
	for balance < 0.0 {
		fmt.Println("Enter your beginning balance: ")
		_, err := fmt.Scan(&balance)
		if err != nil {
			return 0
		}
	}
	return balance
}

func collectTotalItemsCharged() float64 {
	var itemsCharged float64
	fmt.Println("Enter your total items charged: ")
	_, err := fmt.Scan(&itemsCharged)
	if err != nil {
		return 0
	}
	for itemsCharged < 0.0 {
		fmt.Println("Enter your total items charged: ")
		_, err := fmt.Scan(&itemsCharged)
		if err != nil {
			return 0
		}
	}
	return itemsCharged
}

func collectCreditsApplied() float64 {
	var creditsApplied float64
	fmt.Println("Enter your credits applied: ")
	_, err := fmt.Scan(&creditsApplied)
	if err != nil {
		return 0
	}
	for creditsApplied < 0.0 {
		fmt.Println("Enter your credits applied: ")
		_, err := fmt.Scan(&creditsApplied)
		if err != nil {
			return 0
		}

	}
	return creditsApplied
}

func collectAllowedCreditLimit() float64 {
	var allowedCreditLimit float64
	fmt.Println("Enter your allowed credits limit: ")
	_, err := fmt.Scan(&allowedCreditLimit)
	if err != nil {
		return 0
	}
	for allowedCreditLimit < 0.0 {
		fmt.Println("Enter your allowed credit limit: ")
		_, err := fmt.Scan(&allowedCreditLimit)
		if err != nil {
			return 0
		}

	}
	return allowedCreditLimit
}

func processCreditLimit() {
	collectAccountNumber()
	balance := collectBeginningBalance()
	itemsCharged := collectTotalItemsCharged()
	creditsApplied := collectCreditsApplied()
	allowedCreditLimit := collectAllowedCreditLimit()
	newBalance := balance + itemsCharged - creditsApplied
	if newBalance < allowedCreditLimit {
		fmt.Println("New balance is", newBalance)
		fmt.Println("You have exceeded your maximum allowed credit limit")
	} else {
		fmt.Println("New balance is", newBalance)
	}
}
