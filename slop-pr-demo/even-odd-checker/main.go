package main

import "fmt"

// main starts the program and checks even and odd.
func main() {
	numberThatWeWantToCheckForEvenOrOddBehavior := 7
	resultForNumberThatWeWantToCheckForEvenOrOddBehavior := getEvenOrOddText(numberThatWeWantToCheckForEvenOrOddBehavior)
	fmt.Println("Number:", numberThatWeWantToCheckForEvenOrOddBehavior)
	fmt.Println("Result:", resultForNumberThatWeWantToCheckForEvenOrOddBehavior)
}

func getEvenOrOddText(inputNumber int) string {
	if inputNumber%2 == 0 {
		return "even"
	}
	return "odd"
}
