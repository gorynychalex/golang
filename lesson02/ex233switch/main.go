package main

import "fmt"

func main() {
	var c uint32
	fmt.Scan(&c)
	switch {
		case 1 <= c && c <= 9:
			fmt.Println("от 1 до 9")
		case 100 <= c && c <= 250:
			fmt.Println("от 100 до 250")
		case 1000 <= c && c <= 6000:
			fmt.Println("от 1000 до 6000")
	}
}