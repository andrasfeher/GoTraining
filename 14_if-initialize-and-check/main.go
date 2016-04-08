package main

import "fmt"

func main() {

	// hibaellenőrzésre jó pl.
	// if err := file.Chmod(0664); err != nil {....}
	// "to keep scope tight"
	if friend := "Tercsi"; len(friend) > 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if false {
		fmt.Println("first")
	} else if true {
		fmt.Println("second")
	} else {
		fmt.Println("third")
	}

}
