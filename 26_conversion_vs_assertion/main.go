package main

import "fmt"

func main()  {
    // this is an assertion
    // works with interfaces
    // casting works with primitive types
    var val interface{} = 7
    
    // this would not work:
    // fmt.Println(val + 5)
    
    // this works: 
    fmt.Println(val.(int) + 5) // 12
    
    // remainder: this is casting: int(val) this is assertion: val.(int)
}