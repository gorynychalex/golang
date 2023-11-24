package main

import "fmt"

func factorial(n int, ch chan <- int){
	result := 1
	for  i :=1 ; i <= n; i++{
		result *=i
	}
	ch <-result
}

func main(){

	var inCh chan int = make(chan int)

	factorial(5, inCh)

	fmt.Println(<-inCh)
}
