package main

import "fmt"

func main() {

	//------------------------------------------------------------
	// CREATING MAPS
	//------------------------------------------------------------

	// You can define an create maps in go with a map literal

	capitals := map[string]string{
		"India":       "New Delhi",
		"Pakistan":    "Islamabad",
		"Afghanistan": "Nepal",
		"Sri Lanka":   "Sri Jayawardenepura Kotte",
	}
	fmt.Printf("%#v\n", capitals)

	// You can also you the make function to create map
	// and assign values the "traditional way"
	countries := make(map[string]string)

	//------------------------------------------------------------
	// ACCESSING AND LOOPING OVER MAPS
	//------------------------------------------------------------
	// You can loop over the maps with the `range` keyword in for loop
	// The order of maps is not guaranteed

	fmt.Println("--------------------------------------------------------------------------------")
	countries["India"] = "Asia"
	countries["United States of America"] = "North America"
	countries["United Kingdom"] = "Europe"
	countries["Republic of Congo"] = "Africa"

	for k, v := range countries {
		fmt.Printf("%v is located in the %v continent\n", k, v)
	}

	numbers := map[string]int{"one": 1, "two": 2}
	fmt.Printf("%#v\n", numbers)

	// If a map key does not exists then its zero value is returned.
	fmt.Println("three: ", numbers["three"]) // three: 0

	// So in case of checking if a value exists (since sometime, 0 can be a valid value)
	// you use the "comma-exists" idiom
	value, exists := numbers["three"]
	if exists {
		fmt.Printf("value of three is %v\n", value)
	} else {
		fmt.Println("value of three does not exits")
	}

	// You can combine the above two checks in a single statement
	if v, ok := numbers["three"]; ok {
		fmt.Printf("value of three is %v\n", v)
	} else {
		fmt.Println("value of three does not exits")
	}

	//------------------------------------------------------------
	// DELETING MAP KEY
	//------------------------------------------------------------

	// You can use the inbuilt `delete` function to delete a map key.
	// If the key does not exists it will be a `noop`

	fmt.Println("--------------------------------------------------------------------------------")

	delete(countries, "India")

	for k, v := range countries {
		fmt.Printf("%v is located in the %v continent\n", k, v)
	}
}
