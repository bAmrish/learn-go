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
