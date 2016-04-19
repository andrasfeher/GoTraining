package main

import (
    "fmt"
    "sort"
)

type people []string

func (p people) Len() int {return len(p)}
func (p people) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p people) Less(i, j int) bool {return p[i] < p[j]}


func main()  {
    
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
    
}