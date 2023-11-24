package main

import "fmt"

func main() {
	rs := []rune("Это срез рун")

	// Итерируясь мы будем заменять символ 'р' на '*'
	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '$'
		}
	}
	fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs))
}