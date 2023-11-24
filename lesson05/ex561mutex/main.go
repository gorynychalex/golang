package main

import (
	"fmt"
	"sync"
)

var counter int = 0             //  общий ресурс

func main() {
    
	ch := make(chan bool)       // канал
    
	var mutex sync.Mutex

	for i := 1; i < 5; i++{
        go work(i, ch, &mutex)
    }     
    for i := 1; i < 5; i++{
        <-ch
    }
}
func work (number int, ch chan bool, mutex *sync.Mutex){
	mutex.Lock()
    counter = 0
    for k := 1; k <= 5; k++{
        counter++
        fmt.Println("Goroutine", number, "-", counter)
    }
	mutex.Unlock()
    ch <- true
}