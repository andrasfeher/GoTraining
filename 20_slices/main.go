package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 6, 7}
	fmt.Printf("%T\n", mySlice) // [int]
	fmt.Println(mySlice)        // [1 2 3 4 6 7]

	myStringSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(myStringSlice)      // [a b c d e f g]
	fmt.Println(myStringSlice[2:4]) // [c d] from 2 to 4 *not including* 4
	fmt.Println(myStringSlice[2])   // c
	fmt.Println(myStringSlice[:2])  // [a b]
	fmt.Println(myStringSlice[5:])  // [f g]
	fmt.Println(myStringSlice[:])   // [a b c d e f g]
	fmt.Println("myString"[2])      // 83: "S"

	fmt.Println(len(myStringSlice)) // number of items
	fmt.Println(cap(myStringSlice)) // capacity size of the underliying array

	myMadeSlice := make([]int, 0, 5) // length: 0, capacity: 5 (= lenght of the underliying array)
	fmt.Println("--------------------------")
	fmt.Println(myMadeSlice)      // []
	fmt.Println(len(myMadeSlice)) // number of items: 0
	fmt.Println(cap(myMadeSlice)) // capacity size of the underliying array: 5
	fmt.Println("--------------------------")

	// allways duplicates capacity
	for i := 0; i < 80; i++ {
		myMadeSlice = append(myMadeSlice, i)
		fmt.Println("Len:", len(myMadeSlice), "Capacity:", cap(myMadeSlice), "Value: ", myMadeSlice[i])
	}

	//append slice to slice
	weekDays := []string{"Mo", "Tue", "We", "Thu", "Fri"}
	nonWorkDays := []string{"Sat", "Sun"}

	weekDays = append(weekDays, nonWorkDays...) // notice the syntax!
	fmt.Println(weekDays)                       // [Mo Tue We Thu Fri Sat Sun]

	// delete elements from slice by appending ...
	weekDays = append(weekDays[:2], weekDays[3:5]...)
	fmt.Println(weekDays) // [Mo Tue Thu Fri]

	// multi dimensional slices
	attendanceSheet := make([][]string, 0)

	// emp1
	emp1 := make([]string, 8)
	emp1[0] = "Tercsi"
	emp1[1] = "9:00"
	emp1[2] = "9:05"
	emp1[3] = "8:55"

	// store emp1
	attendanceSheet = append(attendanceSheet, emp1)

	// emp1
	emp2 := make([]string, 8)
	emp2[0] = "Fercsi"
	emp2[1] = "9:10"
	emp2[2] = "9:15"
	emp2[3] = "8:50"

	// store emp2
	attendanceSheet = append(attendanceSheet, emp2)

	fmt.Println(attendanceSheet) // [[Tercsi 9:00 9:05 8:55    ] [Fercsi 9:10 9:15 8:50    ]]

	// converting array to slice
	var myArray [24]byte
	myArray[0] = 2
	s := myArray[:] // [:] returns slice
	fmt.Println(s)

}
