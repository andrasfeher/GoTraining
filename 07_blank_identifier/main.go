package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// _ : "I know that a variable should come here, but I am not going
	// to do anything with it"
	res, _ := http.Get("http://www.mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
