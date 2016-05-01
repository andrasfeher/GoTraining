package main

func main()  {
    _ = "breakpoint"
    sum := 0
    
    for i := 0; i < 10; i++ {
        _ = "breakpoint"
        sum += i
    }
}