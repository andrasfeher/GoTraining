package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

type PersonWithTags struct {
	First string
	Last  string `json:"-"`            // do not include into json
	Age   int    `json:"wisdom score"` // replace field name in json
}

type PersonForUnmarshal struct {
	First string
	Last  string
	Age   int `json:"wisdom score"` // tag for unmarshalling
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)         // character codes
	fmt.Printf("%T \n", bs) // []uint8
	fmt.Println(string(bs)) // {"First":"James","Last":"Bond","Age":20} "notExported" lost

	p2 := PersonWithTags{"James", "Bond", 20}
	bs2, _ := json.Marshal(p2)
	fmt.Println(string(bs2)) // {"First":"James","wisdom score":20} "Age" replaced, "Last" is lost

	var p3 PersonForUnmarshal

	// prints empty strings and 0
	fmt.Println(p3.First)
	fmt.Println(p3.Last)
	fmt.Println(p3.Age)

	bs3 := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`) // backtick raw string literal because of double quotes
	json.Unmarshal(bs3, &p3)

	fmt.Println("---------------")
	fmt.Println(p3.First)   // James
	fmt.Println(p3.Last)    // Bond
	fmt.Println(p3.Age)     // 20
	fmt.Printf("%T \n", p3) // main.PersonForUnmarshal

	p4 := Person{"James", "Bond", 20, 007}
	fmt.Println("Encoded:")
	json.NewEncoder(os.Stdout).Encode(p4)

	var p5 Person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p5)

	fmt.Println(p5.First)
	fmt.Println(p5.Last)
	fmt.Println(p5.Age)

}
