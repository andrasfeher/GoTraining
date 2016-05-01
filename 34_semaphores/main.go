package main

import (
    "fmt"
)

func main(){
    
    c := make(chan int)
    done := make(chan bool)
    
    go func(){
        
        for i := 0; i < 10; i++ {
            c <- i
        }
        
        done <- true
    }()
    
    go func(){
        
        for i := 0; i < 10; i++ {
            c <- i
        }

        done <- true
    }()
    
    go func(){
        
        // waits until a value comes through
        // you can throw the values away _ is not needed
        <-done
        <-done
        
        close(c)
    }()
    
    // drains the values when the channel is closed
    for n := range c {
        fmt.Println(n)
    }
    
}