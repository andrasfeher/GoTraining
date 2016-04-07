package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	j := 1

	for j <= 10 {
		fmt.Println(j)
		j++
	}

	// this would run forever
	/*
		k := 0
		for {
			fmt.Println(k)
			k++
		}
	*/

	// prints only odd numbers until 51
	l := 0

	for {
		l++
		if l%2 == 0 {
			continue
		}

		fmt.Println(l)

		if l >= 50 {
			break
		}
	}

}
