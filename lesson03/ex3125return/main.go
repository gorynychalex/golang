package main

import "fmt"

func main() {

	fn()

	i, _ := fn()
	fmt.Println(i)

	_, err := fn()
	if err == nil {
		fmt.Println("Ошибок нет")
	}
}

func fn() (int, error) {
	// Какая-то полезная работа
	// ...
	return 0, nil
}