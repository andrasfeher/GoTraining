package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// the underlying type of Shape is interface
// anything that implements the area() method
// signature, implements the Shape interface
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}


type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

// accepts any variable that implements the empty interface
// so accepts any type because all types implement the empty
// interface
func specs(a interface{})  {
	fmt.Println(a)
}



func main() {
	s := Square{10}
	c := Circle{5}
	fmt.Println(s.area())
	info(s)
	info(c)
	
	
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	specs(fido)
	specs(fifi)
	
	// list of any type
	critters := []interface{}{fido, fifi}
	fmt.Println(critters)
}
