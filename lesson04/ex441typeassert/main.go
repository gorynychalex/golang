package main

import "fmt"

func main() {
	

	var i interface{} = 12

	if v,ok := i.(int); ok {
		fmt.Println(v + 12)
	}

}