package main

import "fmt"

func main() {

	f := square(10)
	fmt.Println(f())

}

func square(y int) func() int {
	x := 4
	f:= func() int {
		
		return x * y
	}
	x = 10
	return f
}