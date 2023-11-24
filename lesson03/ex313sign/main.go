package main

import "fmt"
 
func add(x, y int) int{
    return x + y
}
func multiply(x, y int) int{
    return x * y
}
 
func main() {
     
    var f func(int, int) int
	
	f = add

    fmt.Println(f(3, 4))

	f = multiply
     
	fmt.Println(f(3, 4))
}