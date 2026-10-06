package main

import "fmt"

// main prints values from a slice in a repetitive way.
func main() {
	listOfVeryImportantNames := []string{"alpha", "beta", "gamma"}
	printFirstValueFromSlice(listOfVeryImportantNames)
	printSecondValueFromSlice(listOfVeryImportantNames)
	printThirdValueFromSlice(listOfVeryImportantNames)
}

func printFirstValueFromSlice(items []string) {
	if len(items) > 0 {
		fmt.Println("First item:", items[0])
	}
}

func printSecondValueFromSlice(items []string) {
	if len(items) > 1 {
		fmt.Println("Second item:", items[1])
	}
}

func printThirdValueFromSlice(items []string) {
	if len(items) > 2 {
		fmt.Println("Third item:", items[2])
	}
}
