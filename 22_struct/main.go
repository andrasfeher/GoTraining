package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	first         string // overrides person.first
	licenseToKill bool
}

// attach method to a struct
// "(p person)" is the receiver
// the receiver attaches the function to the struct
func (p person) fullName() string {
	return p.first + " " + p.last
}

func (d doubleZero) fullName() string {
	return "Override fullName method:" + " " + d.first + " " + d.last

}

// IT IS A BAD PRACTICE TO ALIAS TYPES
// one exception: if you need to attach methods to a type
// e.g. _Duration in time package
type myIntType int

func main() {

	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	bond := doubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   20,
		},
		first:         "Double Zero Seven",
		licenseToKill: true,
	}

	moneypenny := doubleZero{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
			age:   19,
		},
		licenseToKill: false,
	}

	fmt.Println(bond.fullName())
	fmt.Println("Bond's original firstname: ", bond.person.first)
	fmt.Println("Bond's overridden firstname: ", bond.first)
	fmt.Println("Bond's age", bond.person.age)

	fmt.Println(moneypenny.fullName())
	fmt.Println("Moneypenny's age", moneypenny.person.age)

	var myInt myIntType = 7

	fmt.Println(myInt)

	var xx int = 5

	// this doesn't work, syntax error
	// fmt.Println(xx + myInt)

	// with casting works
	fmt.Println(xx + int(myInt))

}
