package main

import "fmt"

func main() {


	var x int = 10
	var y int = 20

	var p *int 

	if p != nil { fmt.Println("Value: ", p)}

	fmt.Println(&x)
	fmt.Println(&y)


}
