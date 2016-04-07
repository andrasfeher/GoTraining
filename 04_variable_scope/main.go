package main

import "fmt"

var x int = 42 // package level scope

func main() {

	fmt.Println(x)
	foo()

	y := 17 // block level scope valid from this line only
	fmt.Println(y)

}

func foo() {
	fmt.Println(x)
}
