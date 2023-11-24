package main

import "fmt"

func main() {
	fmt.Println(<-someChanel(5))
}

func someChanel(n int) chan int {

	ch := make(chan int)
	
	go func(){ ch <- n }()

	return ch
}