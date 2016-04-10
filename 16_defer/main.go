package main

import "fmt"

func main() {
	// main lefutása után fogja végrehajtani
	// hasznos pl file lezárásnál, ahol szeretném
	// a lezárást a nyitás közelében látni, de csak
	// a program végrehajtása után lefuttatni
	defer printWorld()
	printHello()
}

func printHello() {
	fmt.Print("Hello ")
}

func printWorld() {
	fmt.Println("world")
}
