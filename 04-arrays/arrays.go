package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//------------------------------------------------------------
	// ARRAYS
	//------------------------------------------------------------

	// Arrays in Go are similar to arrays in C where they are of fixed type
	// and element type

	// Creating an arary with array literal
	flavors := [3]string{"Chocolate", "Mango", "Strawberry"}

	fmt.Println("flavors:", flavors)

	// When creating an Array with literal syntax you can leave out the size.
	weekdays := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println("\nWeekdays:")
	fmt.Println()
	for _, day := range weekdays {
		fmt.Println(day)
	}

	// Or you can create arrays by separating out declaration and initalization
	var num [5]int
	// Since you have not initalized the arrary, it will print the zero-value of the type
	// which in this case would be '0'
	fmt.Println("number\t", num)
	// but you can initialize the arrays with number base indices

	num[0] = rand.Intn(100)
	num[1] = rand.Intn(100)
	num[2] = rand.Intn(100)
	num[3] = rand.Intn(100)
	num[4] = rand.Intn(100)

	fmt.Println("number\t", num)

	// You can get the length of an array with the inbuilt `len` function
	fmt.Printf("Size of %v of type %T is %v\n", num, num, len(num))
}
