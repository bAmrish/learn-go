package main

import "fmt"

func main() {

	type Continent struct {
		name string
	}

	type Country struct {
		name       string
		capital    string
		population float64
		// Go does not allow inheritance like some OOP languages
		// but you can compose a complex data structure by
		// embedding on struct into another
		Continent
	}

	india := Country{
		name:       "Republic of India",
		capital:    "New Delhi",
		population: 1_428_627_663,
		Continent:  Continent{name: "Asia"},
	}

	fmt.Printf("%v country is of type %T\n", india, india)

	// You can use the `.` operator to access properties of struts
	fmt.Printf("%v is on the continent of %v\n", india.name, india.Continent.name)

	// ------------------------------------------------------------
	// INHERITANCE AND OVERLOADING
	// ------------------------------------------------------------

	// Property overloading is default in Go.
	// Lets consider the following example

	type person struct {
		first string
		last  string
		age   int
	}

	type teacher struct {
		person
		subject string
	}

	// All the properties of person are available to teacher implicty.

	albert := person{
		first: "Albert",
		last:  "Einstine",
		age:   42,
	}
	physicsTeacher := teacher{
		person:  albert,
		subject: "Physcis",
	}

	fmt.Println(physicsTeacher) //{{Albert Einstine 42} Physcis}

	// Even though `first` and `last` fields on `physicsTeacher` is on an embeded struct `person`
	// you can access them directly using `physcisTeacher.first` and `physicsTeacher.last` as shown below
	// This mimics the inheritance in most Object Oriented languages
	fmt.Printf("%v %v is teacher of %v subject\n", physicsTeacher.first, physicsTeacher.last, physicsTeacher.subject)

	// Lest consider another two structs shown below
	// manager type has a field of `role` and an embedded struct of `employee`.
	// As we have seen before this mimics the inheritance patten of some object oriented languages.
	// when you access manager.role, the value of role in the manager type will be returned.
	// If you want to access the role of the manager as an employee then you can use `manager.employee.role`
	type employee struct {
		name       string
		department string
		role       string
	}

	type manager struct {
		employee
		role string
	}

	ironMan := employee{
		name:       "Tony Stark",
		department: "Avengers",
		role:       "Offence",
	}

	avengerLead := manager{
		employee: ironMan,
		role:     "Captian",
	}

	fmt.Println(avengerLead) //{{Tony Stark Avengers Offence} captian}

	fmt.Printf("%v has role of %v in %v\n",
		avengerLead.name, avengerLead.role, avengerLead.department) //Tony Stark has role of captian in Avengers%

	// You can access pervious role before he became a manger.
	fmt.Printf("Before %v became %v he was %v\n",
		avengerLead.name, avengerLead.role, avengerLead.employee.role) //Before Tony Stark became captian he was Offence
}
