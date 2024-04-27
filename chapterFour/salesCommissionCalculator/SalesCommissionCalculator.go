package main

import "fmt"

func main() {
	processCommission()
}

func collectNumberOfItems() int64 {
	numberOfItemCategories := 0
	fmt.Println("Enter the number of item categories you sold: ")
	_, err := fmt.Scan(&numberOfItemCategories)
	if err != nil {
		return 0
	}
	return int64(numberOfItemCategories)
}

func collectValueOfItem() float64 {
	var valueOfItem float64
	fmt.Println("Enter the value of item you sold: ")
	_, err := fmt.Scan(&valueOfItem)
	if err != nil {
		return 0
	}
	return valueOfItem
}

func processCommission() {
	totalSales := 0.0
	numberOfItems := collectNumberOfItems()
	for i := 0; i < int(numberOfItems); i++ {
		fmt.Println("Item ", i+1)
		priceOfItem := collectValueOfItem()
		totalSales += priceOfItem
	}
	earnings := 0.09*totalSales + 200
	fmt.Println("Your accumulated earning for the week is ", earnings)
}
