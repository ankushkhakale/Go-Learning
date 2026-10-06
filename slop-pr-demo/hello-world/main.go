package main

import "fmt"

// This is the main function for this program.
// This program prints hello world text.
func main() {
	veryImportantMessageToDisplayToUser := buildHelloWorldMessageForOutput()
	fmt.Println(veryImportantMessageToDisplayToUser)
}

func buildHelloWorldMessageForOutput() string {
	return "Hello, World!"
}
