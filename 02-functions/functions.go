package main

import (
	"fmt"
	"io"
	"os"
)

// In GO we define function with the `func` keyword
// variable types are defined in the list of arguments accepted by the function
func add(x, y int) int {
	return x + y
}

// Go can have multiple return type
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

//------------------------------------------------------------
// defer
//------------------------------------------------------------

// You can defer the execution of the function call using the defer keyword
// When you call a method with defer keyword, its executed at the end of the parent function before it returns.
// Its added to the stack. If there are multiple deferred methods they are all added to the stack
// and are executed in LIFO.

// The arguments to the deferred function are evaluated when the defer executes,
// not when the call executes.
// Besides avoiding worries about variables changing values as the function executes,
// this means that a single deferred call site can defer multiple function executions. Here's a silly example.
func reversePrint() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

// Defer is useful when you want to executed code that releases a resource
func readFile() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error opening file sample.txt", err)
		return
	}

	// We defer the close of the file once the function ends.
	// We don't have to worry about if the function runs successfully
	// or if there is an error reading the file.
	// This call guarantees the file will be closed
	defer f.Close()

	var output []byte
	buffer := make([]byte, 100)

	for {
		n, err := f.Read(buffer[0:])
		output = append(output, buffer[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error read file", err)
			return
		}
	}

	fmt.Println(string(output))
}

// We can use these properties of defer to trace call execution
// in an interesting manner

func trace(s string) string {
	fmt.Println("Entering", s)
	return s
}

func un(s string) {
	fmt.Println("Leaving", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("In a")
	b()
}

func b() {
	defer un(trace("b"))
	fmt.Println("In b")
	c()
}

func c() {
	defer un(trace("c"))
	fmt.Println("In c")
}

func main() {
	sum := add(20, 30)
	fmt.Println("Sum = ", sum)

	x, y := swap("hello", "world!")
	fmt.Println(x, y)

	q, r := divide(133, 13)
	fmt.Printf("133 divided by 13 gives us %v as quotient and %v as reminder", q, r)
	reversePrint()
	readFile()
	a()

	//------------------------------------------------------------
	// FUNCTION EXPRESSION
	//------------------------------------------------------------
	// Functions are fist class citizens in go.
	// So they can be assigned to a variable, passed in as function arguments
	// and can we return type of other functions
	add := func(x, y int) int {
		return x + y
	}

	fmt.Printf("2 + 3 = %v\n", add(2, 3))

}
