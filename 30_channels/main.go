package main

import (
    "fmt"
    "time"
)

func main()  {
    
    // channel must be created with make
    // there are also buffered channels, make(chan int, 3587)
    // c is a channel for single int values   
    c := make(chan int)
    
    go func ()  {
        for i := 0; i < 10; i++ {
            // putting a value into the channel
            c <- i
        }
    }()
    
    go func ()  {
        for {
            // getting value from the channel
            fmt.Println(<-c)
        }
    }()
    
    // gives time to run to the routines
    time.Sleep(time.Second)
    
}