package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // memory address

	// int típusú pointer "a" változó címére
	var b *int = &a

	fmt.Println(b)  // memory address same as &a
	fmt.Println(*b) // 43 (dereference the memory address)

	*b = 42 // change the value at the address to 42
	fmt.Println(a)

	x := 5
	zero(&x)       // pass by value the memory address
	fmt.Println(x) // x is 0

}
