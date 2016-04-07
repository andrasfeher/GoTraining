package main

import "fmt"

func main() {
	x := 0
	increment := func() int { // anonymous func
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
