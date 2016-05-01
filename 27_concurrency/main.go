package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main()  {
    // I am going to add two members to the wait group
    wg.Add(2)
    go foo()
    go bar()
    
    // wait until all members of the wait group tells that 
    // they are done (counter reduces to 0)
    wg.Wait()
}

func foo() {
    for i := 0; i < 45; i++ {
        fmt.Println("Foo:", i)
        time.Sleep(time.Duration(2 * time.Millisecond))
    }
    
    // I am done
    wg.Done()
}

func bar(){
    for i := 0; i < 45; i++ {
        fmt.Println("Bar:", i)
        time.Sleep(time.Duration(2 * time.Millisecond))
    }
    
    // I am done
    wg.Done()
}