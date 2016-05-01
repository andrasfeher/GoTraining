package main

import (
    "fmt"
)

func main(){
    
    n := 10
    c := make(chan int)
    done := make(chan bool)

    // launch n go routines, each puts 10 values on the channel    
    for i := 0; i < n; i++ {
        go func(){
            
            for i := 0; i < 10; i++ {
                c <- i
            }
            
            done <- true
        }()
        
    }
    
    
    go func(){
        
        // waits until all go routines finished
        for i := 0; i < n; i++ {
            <-done
        }

        close(c)
    }()
    
    
    // drains the values when the channel is closed
    for n := range c {
        fmt.Println(n)
    }
    
}