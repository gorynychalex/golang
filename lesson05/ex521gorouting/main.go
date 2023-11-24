package main

import (
	"fmt"
	"time"
)

func main() {

	go func(){fmt.Println("hello from other") }()

	time.Sleep(1 * time.Second)

	fmt.Println("main")
}

// func myFunc() {
// 	fmt.Println("hello from other")
// }