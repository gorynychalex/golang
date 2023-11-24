package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i<5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibbonaci(c, quit)
}

func fibbonaci(c, quit chan int) {
	x, y := 0,1

	for {
		select {
			case c <- x:
				x, y = y, x+y
			case <- quit:
				fmt.Println("quit")
				return
		}
	}
}