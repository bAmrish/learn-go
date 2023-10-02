package main

import "fmt"

// ------------------------------------------------------------
// CREATING STRUTS
// ------------------------------------------------------------

// You can create a composite `type` with strut literal

type Country struct {
	name       string
	capital    string
	population float64
}

func main() {

	// ------------------------------------------------------------
	// CREATING STRUTS
	// ------------------------------------------------------------

	// assigning value to stuct type

	india := Country{
		name:       "India",
		capital:    "New Delhi",
		population: 1_428_627_663,
	}
	fmt.Printf("%v is of type %T\n", india, india)

	// ------------------------------------------------------------
	// ANONYMOUS STRUTS
	// ------------------------------------------------------------
	// You can declare anonymous structs by defining the structure
	// of fields (and methods, which we will see later) inline.
	usa := struct {
		name       string
		capital    string
		population float64
	}{
		name:       "United States of America",
		capital:    "Washington D.C.",
		population: 339_996_563,
	}
	fmt.Printf("%v is of type %T\n", usa, usa) //type =  type struct { name string; capital string; population float64 }

	// ------------------------------------------------------------
	// IMPORTANT
	// ------------------------------------------------------------
	// In Go ANONYMUS STRUTS ARE UNTYPED UNTIL THEY ARE ASSINGED to
	// variable of specific type. So you can do the following:

	var c1 Country = usa
	// c1 = usa
	fmt.Printf("%v is of type %T\n", c1, c1) //type = main.Country

}
