package main

import "fmt"

func main() {

	x, y := swap(5,4)

	fmt.Println(x,y)
}

func swap (a, b int) (int, int) {
	return b, a
}