package main

import "fmt"

// main is the main function for number to word program.
func main() {
	numberValueToConvert := 3
	convertedWordValue := convertNumberToWord(numberValueToConvert)
	fmt.Println("Input number:", numberValueToConvert)
	fmt.Println("Output word:", convertedWordValue)
}

func convertNumberToWord(input int) string {
	if input == 1 {
		return "one"
	}
	if input == 2 {
		return "two"
	}
	if input == 3 {
		return "three"
	}
	if input == 4 {
		return "four"
	}
	if input == 5 {
		return "five"
	}
	return "unknown"
}
