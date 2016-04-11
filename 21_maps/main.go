package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Tercsi"] = 5

	value, ok := m["Tercsi"]
	fmt.Println(value, ok) // 5 true

	delete(m, "Tercsi")
	value2, ok2 := m["Tercsi"]
	fmt.Println(value2, ok2) // 0 false

	// declare and initialize, you can leave out make
	n := map[string]int{
		"Tercsi": 1,
		"Fercsi": 2,
	}
	// this is also ok: n := map[string]int{}
	fmt.Println("map:", n) // map: map[Fercsi:2 Tercsi:1]
	fmt.Println(len(n))    // 2

	// deleting
	delete(n, "Fercsi")
	fmt.Println("map:", n) // map: map[Tercsi:1]

	// testing existence
	if _, exists := n["Tercsi"]; exists {
		fmt.Println("exists")
	} else {
		fmt.Println("not exists")
	}

	// map inside map
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	fmt.Println(elements["Li"]["name"]) // Lithium

	// range loop
	for key, val := range elements {
		fmt.Println(key, "-", val)
	}

}
