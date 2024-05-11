package main

import (
	"fmt"
	"strings"
)

func main() {
	goToMainMenu()
}

func goToMainMenu() {
	var itemsBought []string
	var numbersBought []int64
	var pricesBought []float64
	collectCustomerName()
	item := collectItemBought()
	itemsBought = append(itemsBought, item)
	numberOfItems := collectNumberOfItems()
	numbersBought = append(numbersBought, numberOfItems)
	price := collectItemPrice()
	pricesBought = append(pricesBought, price)
	for strings.ToLower(strings.TrimSpace(collectMoreDecision())) != "no" {
		item = collectItemBought()
		itemsBought = append(itemsBought, item)
		numberOfItems = collectNumberOfItems()
		numbersBought = append(numbersBought, numberOfItems)
		price = collectItemPrice()
		pricesBought = append(pricesBought, price)
	}

	collectCashierName()
	collectDiscount()
	collectAmountPaid()
}

func collectAmountPaid() float64 {
	amountPaid := 0.0
	fmt.Println("How much did the user give to you: ")
	_, err := fmt.Scan(&amountPaid)
	if err != nil {
		return 0
	}
	for amountPaid <= 0 {
		fmt.Println("How much did the user give to you: ")
		_, err := fmt.Scan(&amountPaid)
		if err != nil {
			return 0
		}
	}
	return amountPaid
}

func collectDiscount() float32 {
	discount := 0.0
	fmt.Println("How much discount will he get:")
	_, err := fmt.Scan(&discount)
	if err != nil {
		return 0
	}
	for discount <= 0.0 {
		fmt.Println("How much discount will he get:")
		_, err := fmt.Scan(&discount)
		if err != nil {
			return 0
		}
	}
	return float32(discount)
}

func collectItemPrice() float64 {
	price := 0.0
	fmt.Println("How much per unit: ")
	_, err := fmt.Scan(&price)
	if err != nil {
		return 0
	}
	for price < 1 {
		fmt.Println("How much per unit: ")
		_, err := fmt.Scan(&price)
		if err != nil {
			return 0
		}
	}
	return price
}

func collectCashierName() string {
	cashierName := ""
	fmt.Println("What is your name: ")
	_, err := fmt.Scan(&cashierName)
	if err != nil {
		return ""
	}
	for strings.TrimSpace(cashierName) == "" {
		fmt.Println("What is your name: ")
		_, err := fmt.Scan(&cashierName)
		if err != nil {
			return ""
		}
	}
	return cashierName
}

func collectMoreDecision() string {
	decision := ""
	fmt.Println("Add more items?")
	_, err := fmt.Scan(&decision)
	if err != nil {
		return ""
	}
	for strings.TrimSpace(decision) == "" {
		fmt.Println("Add more items?")
		_, err := fmt.Scan(&decision)
		if err != nil {
			return ""
		}
	}
	return decision
}

func collectNumberOfItems() int64 {
	numberOfItems := 0
	fmt.Println("How many pieces")
	_, err := fmt.Scanf("%d", &numberOfItems)
	if err != nil {
		return 0
	}
	for numberOfItems < 1 {
		fmt.Println("How many pieces")
		_, err := fmt.Scanf("%d", &numberOfItems)
		if err != nil {
			return 0
		}
	}
	return int64(numberOfItems)
}

func collectItemBought() string {
	itemBought := ""
	fmt.Println("What did the user buy: ")
	_, err := fmt.Scan(&itemBought)
	if err != nil {
		return ""
	}
	for strings.TrimSpace(itemBought) == "" {
		fmt.Println("What did the user buy: ")
		_, err := fmt.Scan(&itemBought)
		if err != nil {
			return ""
		}
	}
	return itemBought
}

func collectCustomerName() string {
	customerName := ""
	fmt.Println("What is the customer's name: ")
	_, err := fmt.Scan(&customerName)
	if err != nil {
		return ""
	}
	for strings.TrimSpace(customerName) == "" {
		fmt.Println("What is the customer's name: ")
		_, err := fmt.Scan(&customerName)
		if err != nil {
			return ""
		}
	}
	return customerName

}
