/**
 * A simple go package to demonstrate execution flow and some concepts
**/
// Every go program has a "main" package. Execution of the program starts within the main package.
package main

import "fmt"

func init() {
	fmt.Println("Program initialized.")
}

// Every go program has a function. Execution of the program starts with main function.
// Execution of the program ends when the main function ends.
func main() {
	fmt.Println("Hello Go!")
}
