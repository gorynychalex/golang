package main

import (
	"fmt"
	// "time"
)

func main() {
	
	var intChannel chan int = make(chan int)
	
	intChannel <- 10

	// go func ()  {
	// 	// fmt.Println("горутина с каналом")
	// 	time.Sleep(2*time.Second)
	// 	n := 15
	// 	intChannel <- n
	// }()

	i := <-intChannel
	fmt.Println(i)
	fmt.Println("The end")
}