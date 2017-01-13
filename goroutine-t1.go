package main

import (
    "fmt"
    "sync"
)

func gofunc(i int, rs chan int){
    fmt.Println(i, "hw f")
    rs<-100 + i
}

func main(){
    var wg sync.WaitGroup
    rs := make(chan int)
    for i:=1; i<10; i++ {
        wg.Add(1)
        go func(x int){
            gofunc(x, rs)
            wg.Done()
        }(i)
        fmt.Println(<-rs)
    }

    close(rs)
    wg.Wait()
}
