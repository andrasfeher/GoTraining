package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

var wg sync.WaitGroup
var counter int64
var mutex sync.Mutex

func main()  {
    wg.Add(2)

    go incrementor("Foo:")
    go incrementor("Bar:")

    wg.Wait()
    
    fmt.Println("Final counter: ", counter)
}

func incrementor(s string)  {
    for i := 0; i < 20; i++ {
        time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
        atomic.AddInt64(&counter, 1)
        fmt.Println(s, i, "Counter:", counter)
    } 
    wg.Done()
}

// go run -race main.go
// vs
// go run main.go