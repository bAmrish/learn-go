package main

import "fmt"

// In GO we define function with the `func` keyword
// variable types are defined in the list of arguments accepted by the function
func add(x, y int) int {
	return x + y
}

//Go can have multiple return type
func swap(x, y string) (string, string) {
	return y, x
}

// In Go you can define variable names for the return type.
// If you do that those variables can be defined at the top of the function for you
// and you can return a "naked" return type.
func divide(dividend, divisor int) (quotient int, reminder int) {
	reminder = dividend % divisor
	quotient = (dividend - reminder) / divisor
	//This is called as naked return
	return
}

func main() {
	sum := add(20, 30)
	fmt.Println("Sum = ", sum)

	a, b := swap("hello", "world!")
	fmt.Println(a, b)

	q, r := divide(133, 13)
	fmt.Printf("133 divided by 13 gives us %v as quotient and %v as reminder", q, r)
}
