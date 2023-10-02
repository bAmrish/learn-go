package main

import (
	"fmt"
	"math"
)

//------------------------------------------------------------
// INTERFACES
//------------------------------------------------------------

// In Go interfaces are collection of functions

// Shape interface defining 2 functions `area()` and `perimeter`
type Shape interface {
	area() float64
	perimeter() float64
}

// Rect type
type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

//------------------------------------------------------------
// IMPORTANT
//------------------------------------------------------------
// A type implements an interface by implementing its methods.
// There is NO EXPLICIT DECLARATION OF INTENT, no "implements" keyword.
// Implicit interfaces decouple the definition of an interface from its implementation,
// which could then appear in any package without prearrangement.

// Rect type implementing interfaces
func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// ------------------------------------------------------------
// Circle type implementing interfaces
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// ------------------------------------------------------------
// Now you can write functions that utilizes the interfaces
func measure(s Shape) {
	fmt.Println(s)
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

//------------------------------------------------------------
// Stringer INTERFACE
//------------------------------------------------------------
// Stringer is an interface similar to `toString()` in java
// It has a single String() method that
// returns string representation of that type

func (r Rect) String() string {
	return fmt.Sprintf(`Rectangle (width: %v, height: %v)`, r.width, r.height)
}

func (c Circle) String() string {
	return fmt.Sprintf(`Circle (radius: %v)`, c.radius)
}

// ------------------------------------------------------------
// EMPTY INTERFACE
// ------------------------------------------------------------
// You can create an empty interface as shown below to represent
// `any` type
type any interface{}

// A function that can take in any type of parameter
// can use this interface as argument type

func myPrintln(a any) {
	fmt.Println("My custom print")
	fmt.Println(a)
}

// ------------------------------------------------------------
// TYPE ASSERTIONS
// ------------------------------------------------------------

// type assertions provide access to interface's underlying value type
func getType(t any) {

	if _, ok := t.(string); ok {
		fmt.Println("The type is string")
	} else if _, ok := t.([]int); ok {
		fmt.Println("The type is []int")
		return
	} else if _, ok := t.(Rect); ok {
		fmt.Println("The type is Rect")
		return
	} else if _, ok := t.(Circle); ok {
		fmt.Println("The type is Circle")
		return
	} else {
		fmt.Println("Unable to determine type")
		return
	}
}

func main() {
	r := Rect{width: 2, height: 3}
	c := Circle{radius: 10}

	measure(r)
	fmt.Println("-----------------------------------")
	measure(c)

	myPrintln("Some string")
	myPrintln([]int{1, 2, 3, 4})
	myPrintln(r)
	myPrintln(c)

	getType("Some string")
	getType([]int{1, 2, 3, 4})
	getType(r)
	getType(c)
}
