package main

import (
	"fmt"
	"strings"
)

func main() {
	passCheck := 0

	fmt.Println("Global Warming brain tease\nThe Evidence for Rapid Climate Change is Compelling, kindly provide answers from option A-D")

	fmt.Println("1. What is the effect of Global warming on Temperature?\nA.Global Temperature is increasing\nB.Global Temperature is decreasing\nC.Global Temperature remains unchanged\nD.Global Temperature is doubling")
	firstAnswer := ""
	_, err := fmt.Scan(&firstAnswer)
	if err != nil {
		return
	}

	if strings.ToLower(firstAnswer) == "a" {
		passCheck++
	}

	fmt.Println("\n2. Due to Global warming, the ocean is getting?\nA.Colder\nB.Warmer\nC.Deeper\nD.Shallower")
	secondAnswer := ""
	_, err = fmt.Scan(&secondAnswer)
	if err != nil {
		return
	}

	if strings.ToLower(secondAnswer) == "b" {
		passCheck++
	}

	fmt.Println("\n3. What is an effect of Global warming on Greenland and Antarctic ice sheets?\nA.Ice sheets are shrinking\nB.ice sheets are expanding\nC.Ice sheets remain unchanged\nD.Ice sheets are spreading")
	thirdAnswer := ""
	_, err = fmt.Scan(&thirdAnswer)
	if err != nil {
		return
	}

	if strings.ToLower(thirdAnswer) == "a" {
		passCheck++
	}

	fmt.Println("\n4. With the current rate of global warming, snow cover is gradually __\nA.Increasing\nB.Unchanged\nC.Decreasing\nD.Covering trees")
	fourthAnswer := ""
	_, err = fmt.Scan(&fourthAnswer)
	if err != nil {
		return
	}

	if strings.ToLower(fourthAnswer) == "c" {
		passCheck++
	}

	fmt.Println("\n5.Global warming has caused a/an ___ in Sea level\nA.Fall\nB.Even Level\nC.Rise\nD.No change")
	fifthAnswer := ""
	_, err = fmt.Scan(&fifthAnswer)
	if err != nil {
		return
	}

	if strings.ToLower(fifthAnswer) == "c" {
		passCheck++
	}

	if passCheck == 5 {
		fmt.Println("\nExcellent")
	} else if passCheck == 4 {
		fmt.Println("\nVery good")
	} else if passCheck <= 3 {
		fmt.Println("\nTime to brush up on your knowledge of global warming")
	}

	fmt.Println("\nWebsite: www.climate.nasa.gov")
}
