package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

func main() {
	var input int

	n,err := fmt.Scan(&input)

	if err != nil {
		fmt.Println(n)
		fmt.Println(err)
	} else {
		fmt.Println(divide(input, 5)) //Выведем результат
	}
}