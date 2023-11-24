package main

import "fmt"

func main() {
	
	intCh := make(chan int, 5)
	intCh <- 10
	intCh <- 20
	
	close(intCh)
	
	for i := 0; i < cap(intCh); i++ { 
		if val, opened := <-intCh; opened { 
		   fmt.Println(val) 
		} else { 
		   fmt.Println("Channel closed!") 
		} 
  }

}