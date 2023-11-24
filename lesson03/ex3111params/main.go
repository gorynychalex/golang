package main

import "fmt"

func main() {
	sum := add(4,5)
	mult := multiply(4,5)

	fmt.Println(sum)
	fmt.Println(mult)
}

// сигнатура: func(int,int) int
func add(a, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}
