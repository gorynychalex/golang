package main

import "fmt"

func main() {

	p := new(int)

	fmt.Println(p)
	fmt.Println(*p)

	*p = 25

	fmt.Println(*p)

}
