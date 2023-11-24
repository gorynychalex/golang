package main

import "fmt"

func main() {

	d:=5
	fmt.Println(d)
	changeValue(&d)
	fmt.Println(d)
}

func changeValue(i *int) {
	
	*i=10
}


