package main

import "fmt"

func main() {
	processMileage()
}

func processMileage() {
	var decision int
	var totalMilesDrivenForTheTrips float64 = 0
	var totalGallonsForTheTrips int64 = 0

	for decision != -1 {
		milesDriven := collectMilesDriven()
		totalMilesDrivenForTheTrips += milesDriven
		numberOfGallons := collectNumberOfGallonsForTrip()
		totalGallonsForTheTrips += numberOfGallons
		milesPerGallon := calculateMilesPerTrip(milesDriven, numberOfGallons)
		fmt.Println("The miles per gallon for this trip are", milesPerGallon)
		fmt.Print("Enter -1 to quit or enter any number to continue: ")
		_, err := fmt.Scan(&decision)
		if err != nil {
			return
		}
	}
	totalMilesDrivenPerTotalGallons := totalMilesDrivenForTheTrips / float64(totalGallonsForTheTrips)
	fmt.Println("The total miles driven for all the gallons is", totalMilesDrivenPerTotalGallons)
}

func collectNumberOfGallonsForTrip() int64 {
	fmt.Println("Enter the gallons used for this trip")
	var numberOfGallons int64
	fmt.Scanf("%d", &numberOfGallons)
	return numberOfGallons
}

func collectMilesDriven() float64 {
	fmt.Print("Enter the number of miles driven for this trip: ")
	var milesDriven float64
	fmt.Scanf("%f", &milesDriven)
	return milesDriven
}

func calculateMilesPerTrip(input float64, numberOfGallons int64) float64 {
	milesPerGallon := input / float64(numberOfGallons)
	return milesPerGallon
}
