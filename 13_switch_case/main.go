package main

import "fmt"

func main() {

	// break is not needed
	// fallthrough is optional
	// prints:
	// case 3
	// case default

	switch "case 3" {
	case "case 1":
		fmt.Println("case 1")
	case "case 2":
		fmt.Println("case 2")
		fallthrough
	case "case 3":
		fmt.Println("case 3")
		fallthrough
	default:
		fmt.Println("case default")
	}

	// multiple cases
	// prints "case 2 or case 3"
	switch "case 3" {
	case "case 1":
		fmt.Println("case 1")
	case "case 2", "case 3":
		fmt.Println("case 2 or case 3")
	default:
		fmt.Println("case default")
	}

	// no expression: runs the first matching
	// default if no match

	friend := "Tercsi"

	switch {
	case len(friend) >= 2:
		fmt.Println("Tercsi")
	case friend == "Fercsi":
		fmt.Println("Fercsi")
	default:
		fmt.Println("nothing matched")
	}

}
