package lesson01

import "fmt"

func main() {
	

	var a []int

	var b []int = []int {1,2,3}

	c:= []int {4,5,6}

	d := []int {1: 12}

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
	fmt.Println(d, len(d), cap(d))

	f := make([]int, 10, 15)

	fmt.Println(f, len(f), cap(f))


}