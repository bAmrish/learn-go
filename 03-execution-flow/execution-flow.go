// A simple go package to demonstrate execution flow and some concepts

// Every go program has a "main" package. Execution of the program starts within the main package.
package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Program initialized.")
}

// Every go program has a function. Execution of the program starts with main function.
// Execution of the program ends when the main function ends.
func main() {
	fmt.Println("Hello Go!")

	//------------------------------------------------------------
	// IF STATEMENTS
	//------------------------------------------------------------
	x := 20

	// A simple if else statement
	if x < 50 {
		fmt.Printf("%v is less than 50\n", x)
	} else if x == 50 {
		fmt.Printf("%v is equal to 50", x)
	} else {
		fmt.Printf("%v is more than 50", x)
	}

	// Go also allow us to declare a statement before the condition.
	// its called if statement-statement form(?)
	if r := rand.Intn(100); r < 50 {
		fmt.Printf("%v is less than 50\n", r)
	} else if r == 50 {
		fmt.Printf("%v is equal to 50\n", r)
	} else {
		fmt.Printf("%v is more than 50\n", r)
	}

	// There is a comma or format fr if statement that we are going to look into later.

	//------------------------------------------------------------
	// SWITCH STATEMENTS
	//------------------------------------------------------------
	s := 30

	// In GO switch statements does not have default fallthrough
	// like some other `C` based languages. So no `break` statements
	// are required after each case. In fact if you need a fallthrough
	// you would have to explicitly specify it
	switch {
	case s < 50:
		fmt.Printf("%v is less than 50\n", s)
		fallthrough // explicit fallthrough required.
	case s == 30:
		fmt.Printf("%v is equal to 30\n", s)
	case s > 50 && s < 100:
		fmt.Printf("%v is between 50 and 100\n", s)
	default:
		fmt.Printf("%v is more than 100\n", s)
	}

	//------------------------------------------------------------
	// SELECT STATEMENTS
	//------------------------------------------------------------
	// We will look at select statements when we learn about `channels` in go.
	// select statements are similar to `switch` statements but they help select
	// values from channels.

	//------------------------------------------------------------
	// FOR LOOPS
	//------------------------------------------------------------
	// Go has for loops similar to `C` language for loop
	// it follows the following pattern
	// for initialization; increment; post {
	//     do stuff
	// }

	println("Simple C like for loop")
	for counter := 0; counter < 5; counter++ {
		fmt.Printf("For loop counter: %v\n", counter)
	}

	// Just specify a condition in the for simulates the `while` loop in `C` like languages.
	y := 10
	for y >= 5 {
		fmt.Printf("While counter: %v\n", y)
		y--
	}

	// OR you can have a for(;;) simulate by skipping the the parts altogether
	// for also supports `break` and `continue` keyword
	z := 0
	for {
		z++
		if z == 20 {
			break
		}

		if z%2 == 0 {
			continue
		}

		fmt.Printf("%v is an odd number\n", z)
	}

	// The range operator also helps loop over arrays and maps (which we will learn later)

	//define a slice of integers
	numbers := []int{20, 21, 22, 23, 24}

	// loop over the slice of integers
	for index, value := range numbers {
		fmt.Printf("%v. %v\n", index+1, value)
	}

	// define a map of countries and capitals
	countries := map[string]string{
		"Russia":   "Moscow",
		"Mongolia": "Ulaanbaatar",
		"China":    "Beijing",
		"India":    "New Delhi",
		"Japan":    "Tokyo",
	}

	// In case of map range operator returns key, value pair from the map
	for country, capital := range countries {
		fmt.Printf("Capital of %v is %v\n", country, capital)
	}
}
