package main

import "fmt"

func main() {
	greet("Jane", "Doe")
	fmt.Println(formattedGreet("John", "Doe"))

	h, g := extFormattedGreet("John", "Doe")
	fmt.Println(h, g)

	fmt.Println("average: ", average(1, 2, 4, 5, 100))

	// a ... feldarabolja a slice-t (listát) különálló elemekre, így lehet
	// átadni a variadic függvénynek
	data := []float64{1, 2, 3, 5, 6}
	fmt.Println("average: ", average(data...))

	// slice átadása függvénynek
	fmt.Println("slice average: ", sliceAverage(data))

	// func expression: assing anonymous function to variable
	// function inside of a function
	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
	fmt.Printf("%T\n", greeting) // prints "func()"

	// use function returned by makeGreeter
	greeter := makeGreeter()
	fmt.Println(greeter())
	fmt.Printf("%T\n", greeter) // prints "func() string"

	// increment value
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())

	// callback
	visit([]int{6, 12, 14, 16}, multiplyByTwo)

	// callback with anonymous function
	visit([]int{6, 12, 14, 16}, func(n int) {
		fmt.Println(n * 3)
	})

	// callback filter
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)

	fmt.Println(factorial(40))

}

// elég egyszer megadni a típust, ha azonosak
func greet(fname, lname string) {
	fmt.Println(fname, lname)
}

func formattedGreet(fname, lname string) string {
	return fmt.Sprintf("Hello %s %s", fname, lname)
}

// return twe results
func extFormattedGreet(fname, lname string) (string, string) {
	return fmt.Sprintf("Hello %s %s", fname, lname), "How are you?"
}

// változó számú paraméterlista (variatic)
func average(sf ...float64) float64 {
	fmt.Println(sf)

	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

func sliceAverage(sf []float64) float64 {
	fmt.Println(sf)

	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

// create and return anonymous function which returns string
func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}

}

func makeCounter() func() int {

	// automatically initialized to 0
	var x int
	return func() int {
		x++
		return x
	}
}

func multiplyByTwo(n int) {
	fmt.Println(n * 2)
}

// callback function
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

// callback filter
func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
