package main

import (
    "fmt"
)

func gofunc(ch chan int, qt chan bool){
	x := 100
    for {
        select {
            case ch<-x:
				x += 1
            case <-qt:
                fmt.Println("finish.")
                return
        }
    }
}

func main(){
    ch := make(chan int)
    qt := make(chan bool, 1)
	go func(){
		for i:=1;i<10;i++ {
			fmt.Println(i, <-ch)
		}
		qt<-true
	}()
    gofunc(ch, qt)
    close(ch)
    close(qt)
}
