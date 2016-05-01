package main

import (
    "fmt"
)

func main(){
    c := make(chan int)
    
    go func ()  {
      for i := 0; i < 10; i++ {
          c <- i
      }
      
      // from the moment I close the channel, I cannot write to it anymore
      // but the content can be read
      close(c)  
    }()
    
    // read the content of the channel until nothing is left on the channel
    // no need for sleep as for... blocks the execution of the code
    for n := range c {
        fmt.Println(n)
    }
}