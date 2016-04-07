package main

import "fmt"

func main() {

	for i := 0; i < 200; i++ {
		fmt.Printf("%03d \t %b \t %#x %q\n", i, i, i, i) // q: utf-8 character
	}

}
