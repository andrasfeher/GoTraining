package main

import "fmt"

const p string = "death & taxes"

const (
	Pi       = 3.1415926
	Language = "Go"
)

const (
	A = iota // 0
	B        // 1
	C        // 2
)

// újra kezdi a sorszámozást
const (
	D = iota // 0
	E
	F
)

const (
	_ = iota      // 0
	H = iota * 10 // 1 * 10
	I = iota * 10 // 2 * 10
)

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("Pi - ", Pi)
	fmt.Println("Language - ", Language)

	fmt.Println("A - ", A)
	fmt.Println("B - ", B)
	fmt.Println("C - ", C)

	fmt.Println("D - ", D)
	fmt.Println("E - ", E)
	fmt.Println("F - ", F)

	fmt.Println("H - ", H)
	fmt.Println("I - ", I)

}
