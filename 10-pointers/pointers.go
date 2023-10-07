package main

import "fmt"

func main() {
	// x is a simple x variable
	x := 42
	fmt.Println("x =", x) // x = 45

	// `&x` refers to the address of the memory where variable x is store
	fmt.Println("&x =", &x) // &x = 0xc0000120e0

	//We can assign that address to another variable and this
	// another variable has a special name called pointer.
	// so POINTER is a special VARIABLE that holds ADDRESS of
	// another variable.
	a := &x
	fmt.Println("a =", a) // "a = 0xc0000120e0"

	// Once you have pointer variable,
	// you can access the VALUE AT THE ADDRESS of pointer
	// by using the `*` operator also know as "dereferencing" operator
	fmt.Println("*a =", *a)           // *a = 42
	fmt.Println("a == &x :", a == &x) // a == &x : true

	// Now lets talk about types
	fmt.Printf("type of x = %T\n", x)   // type of x = int
	fmt.Printf("type of &x = %T\n", &x) // type of &x = *int
	fmt.Printf("type of a = %T\n", a)   // type of a = *int
	fmt.Printf("type of *a = %T\n", *a) // type of *a = int

	// When you modify the value at the address directly
	// you modify the value of the original variable
	*a = 45
	fmt.Println("x =", x) // x = 45

}
