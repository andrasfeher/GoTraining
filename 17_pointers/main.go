package main

import "fmt"

// change value of variable that is not in scope
func main() {
	age := 44
	fmt.Println(&age) //memory address

	changeMe(&age)

	fmt.Println(&age) //memory address
	fmt.Println(age)  // 24

	// passing slice: there is not need
	// for pointers, it is already a reference
	// type
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMeSlice(m)
	fmt.Println(m)

}

// z: pointer to int
func changeMe(z *int) {
	fmt.Println(z)  // memory address of age
	fmt.Println(*z) // dereferencing, prints 44

	*z = 24

	fmt.Println(z)  // memory address of age
	fmt.Println(*z) // 24

}

func changeMeSlice(z []string) {
	z[0] = "Todd"
	fmt.Println(z) // [Todd]
}
