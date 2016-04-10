package main

import "fmt"

func main() {

	var x [58]string

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i, v := range x {
		fmt.Printf("%v - %T - %s\n", v, v, v)

		// only the first 50 items
		if i > 50 {
			break
		}
	}
}
