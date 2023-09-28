// This is  simple function for go

// All go program requires a main package.
// Execution of go programs starts from within the main package.
package main

// "fmt" is a standard library function that is used to print formatted string.
import "fmt"

// All go programs require a main function within main package to start execution.
func main() {

	// In go we use the 'var' keyword to declare age.
	// we can have multiple variables declared in a single statement.
	var name, age = "Joe", 32

	// fmt.Println function prints a line to stdout.
	// Go supports utf-8 encoding. Hence we can include emoji's
	fmt.Println("Hello, Gopher! 🙌😎🐼🐼🐔🦄")

	// fmt.Printf is similar to 'printf' function in C.
	fmt.Printf("%s is %d years old\n", name, age)

	// You can use %v to print value of a variable and %T to print type of the variable
	fmt.Printf("Type of %v is %T and type of %v %T\n", name, name, age, age)

	// In Go you can print raw string literal between backticks (``)
	fmt.Println(`
	Roses are red
	Violets are blue
	Go is fun
	and so are you!	
	`)

	// Initializing a variable in Go, you declare a type and assign a value.
	var aNumber int = 20
	var aString string = "The Gopher"

	// This is a short form for assigning variables
	// This is called short declaration operator.
	// It can only be used inside a function (As oppose to package level)
	aBoolean := true

	fmt.Printf("aNumber = %v\n", aNumber)
	fmt.Printf("aString = %v\n", aString)
	fmt.Printf("aBoolean = %v\n", aBoolean)

	// Go does not allow variables that are unused (code pollution)
	// your code will throw error if you uncomment the line below.
	// var anUnusedVariable = "Will throw error on uncomment"
}
