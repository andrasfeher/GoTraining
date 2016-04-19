package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

type person struct {
	Name string
	Age  int
}

// control the default print format of the custom type
func (p person) String() string {
	return fmt.Sprintf("DATA %s: %d", p.Name, p.Age)
}

// ByAge implements the sort.Interface for []person based on the Age field
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	studyGroup := people{"Zeno", "John", "All", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	// you cannot attach method directly to the built-in type
	// use the StringSlice (= []string) type from the sort package
	// it already has the necessary methods
	s := []string{"Zeno", "John", "All", "Jenny"}

	// convert []string to StringSlice in sort package
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	fmt.Println(sort.StringSlice(s).Search("John")) // 2

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	personList := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 18},
		{"Jenny", 26},
	}

	fmt.Println(personList[0])
	fmt.Println(personList) // [DATA Bob: 31 DATA John: 42 DATA Michael: 18 DATA Jenny: 26]
	sort.Sort(ByAge(personList))
	fmt.Println(personList) // [DATA Michael: 18 DATA Jenny: 26 DATA Bob: 31 DATA John: 42]
}
