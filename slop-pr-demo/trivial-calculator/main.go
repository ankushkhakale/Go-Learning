package main

import "fmt"

// This calculator is very basic and very repetitive on purpose.
func main() {
	firstNumberValueForMath := 10
	secondNumberValueForMath := 5

	fmt.Println("Add result:", doAddMathOperation(firstNumberValueForMath, secondNumberValueForMath))
	fmt.Println("Subtract result:", doSubtractMathOperation(firstNumberValueForMath, secondNumberValueForMath))
	fmt.Println("Multiply result:", doMultiplyMathOperation(firstNumberValueForMath, secondNumberValueForMath))
	fmt.Println("Divide result:", doDivideMathOperation(firstNumberValueForMath, secondNumberValueForMath))
}

func doAddMathOperation(first int, second int) int {
	return first + second
}

func doSubtractMathOperation(first int, second int) int {
	return first - second
}

func doMultiplyMathOperation(first int, second int) int {
	return first * second
}

func doDivideMathOperation(first int, second int) int {
	if second == 0 {
		return 0
	}
	return first / second
}
