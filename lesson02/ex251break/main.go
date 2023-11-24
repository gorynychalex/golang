package main

import "fmt"

func main() {
	
	sum := 0

	for i := 1; i < 10; i++ {
		
		if i %2 == 0 {continue}

		if i > 7 { break }

		sum += i

		fmt.Println(i)
	}

	fmt.Println(sum)
}