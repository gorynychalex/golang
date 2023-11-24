package main

import "fmt"

func main() {


	// multiply := func(x, y int) int { return x * y }

	// fmt.Println(multiply(20,30))

	result := func(x, y int) int { return x * y }(100,200)

	fmt.Println(result)

}
